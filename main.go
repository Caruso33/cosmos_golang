package main

import (
	"context"
	"lulu/parts"
)

func main() {
	// parts.GetBank()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	parts.GetProposals(ctx)
}
