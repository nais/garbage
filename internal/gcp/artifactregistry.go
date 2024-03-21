package gcp

import (
	"context"
	"errors"
	"fmt"

	artifactregistry "cloud.google.com/go/artifactregistry/apiv1"
	"cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func ArtifactRegistryClient(ctx context.Context, serviceAccountEmail string) (*artifactregistry.Client, error) {
	ts, err := ImpersonatedTokenSource(ctx, serviceAccountEmail)
	if err != nil {
		return nil, fmt.Errorf("create token source: %w", err)
	}

	return artifactregistry.NewClient(ctx, option.WithTokenSource(ts))
}

func ListDockerImages(ctx context.Context, garClient *artifactregistry.Client, parent string) ([]*artifactregistrypb.DockerImage, error) {
	results := make([]*artifactregistrypb.DockerImage, 0)
	iter := garClient.ListDockerImages(ctx, &artifactregistrypb.ListDockerImagesRequest{Parent: parent})

	for {
		x, err := iter.Next()
		if errors.Is(err, iterator.Done) {
			break
		} else if err != nil {
			return nil, err
		}
		results = append(results, x)
	}

	return results, nil
}
