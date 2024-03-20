package main

import (
	"context"
	"flag"
	"log"

	"github.com/nais/garbage/internal/gcp"
)

func main() {
	log.Fatal("run", run())
}

func run() error {
	ctx := context.Background()

	var serviceAccountEmail string

	flag.StringVar(&serviceAccountEmail, "service-account", "", "Service account email")
	flag.Parse()

	_, err := gcp.ArtifactRegistryClient(ctx, serviceAccountEmail)
	if err != nil {
		return err
	}

	return nil
}