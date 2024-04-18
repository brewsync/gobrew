package gobrew_test

import (
	"context"
	"testing"

	"github.com/brewsync/gobrew"
)

func TestNonexistentPackageError(t *testing.T) {
	// Create a new Homebrew client.
	b := gobrew.New()

	// Get information about a package that does not exist.
	_, err := b.Info(context.Background(), "this-package-does-not-exist")
	if err == nil {
		t.Fatalf("Expected an error when package does not exist, got nil")
	}
}

func TestUnknownBrewCommandError(t *testing.T) {
	// Create a new Homebrew client.
	b := gobrew.New()

	// Run brew with bad arguments.
	_, err := b.Exec(context.Background(), "this-command-does-not-exist")
	if err == nil {
		t.Fatalf("Expected an error when running unknown brew command, got nil")
	}
}
