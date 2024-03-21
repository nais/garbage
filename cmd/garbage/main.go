package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nais/garbage/internal/config"
	"github.com/nais/garbage/internal/gcp"
)

func main() {
	log.Fatal("run", run())
}

func run() error {
	out, err := os.Create("/tmp/images.json")
	if err != nil {
		return err
	}
	defer out.Close()

	ctx := context.Background()
	cfg, err := config.New()
	if err != nil {
		return fmt.Errorf("configuration error: %w", err)
	}

	garClient, err := gcp.ArtifactRegistryClient(ctx, cfg.ServiceAccountEmail)
	if err != nil {
		return err
	}

	projectLocation := gcp.ProjectLocationID{
		Project:  cfg.Project,
		Location: cfg.Location,
	}

	// One repository per team.
	repositories, err := gcp.ListRepositories(ctx, garClient, projectLocation)
	if err != nil {
		return err
	}

	enc := json.NewEncoder(out)
	enc.SetIndent("", "  ")
	out.WriteString("[")

	for _, repository := range repositories {

		tm := time.Now()
		// what about SLSA attestations?
		images, err := gcp.ListDockerImages(ctx, garClient, repository.Name)
		elapsed := time.Now().Sub(tm)
		fmt.Printf("%s %s\n", elapsed.String(), repository.Name)

		_ = enc.Encode(images)
		out.WriteString(",")

		if err != nil {
			fmt.Printf("ERROR: %s", err)
		}
	}

	out.WriteString("]")
	return nil
}
