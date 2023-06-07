package main

import (
	"context"
	"log"
	"os"

	"github.com/aaronland/beuys"
)

func main() {

	ctx := context.Background()

	err := beuys.Feed(ctx, os.Stdout)

	if err != nil {
		log.Fatalf("Failed to wrap list of blog posts, %v", err)
	}
}
