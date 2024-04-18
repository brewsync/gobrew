package gobrew

import (
	"bytes"
	"context"
	"encoding/json"
	"os/exec"
)

const (
	brewCommand     = "brew"
	brewInfoCommand = "info"
)

type Brew interface {
	Info(ctx context.Context, name string) (*Package, error)
	List(ctx context.Context) ([]Package, error)
}
type brew struct{}

// New creates a new Homebrew client.
func New() Brew {
	b := &brew{}

	return b
}

// Info returns information about a package.
func (b *brew) Info(ctx context.Context, name string) (*Package, error) {
	var packages []Package
	output, err := b.Exec(ctx, brewInfoCommand, "--json=v1", name)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(output), &packages)
	if err != nil || len(packages) == 0 {
		return nil, err
	}

	return &packages[0], nil
}

// List returns a list of installed packages.
func (b *brew) List(ctx context.Context) ([]Package, error) {
	var packages []Package
	output, err := b.Exec(ctx, brewInfoCommand, "--json=v1", "--installed")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(output), &packages)
	if err != nil {
		return nil, err
	}

	return packages, nil
}

// Exec executes a raw Homebrew command. Can be used to run any Homebrew command, though it would be advisable to use the conveniece methods provided by the Brew interface.
func (b *brew) Exec(ctx context.Context, args ...string) (string, error) {
	cmd := exec.CommandContext(ctx, brewCommand, args...)

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	if err := cmd.Run(); err != nil {
		return "", err
	}

	return stdout.String(), nil
}
