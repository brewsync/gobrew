package main

import (
	"context"
	"fmt"
	"time"

	"github.com/brewsync/gobrew"
)

func main() {
	brew := gobrew.New()
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Minute,
	)
	defer cancel()

	packages, err := brew.List(ctx)
	if err != nil {
		panic(err)
	}

	for _, p := range packages {
		fmt.Printf("%s (%s)\n", p.FullName, p.Versions.Stable)
	}

	fmt.Printf("Total packages: %d\n", len(packages))
}
