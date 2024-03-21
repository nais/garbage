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

func ListRepositories(ctx context.Context, garClient *artifactregistry.Client, id ProjectLocationID) ([]*artifactregistrypb.Repository, error) {
	results := make([]*artifactregistrypb.Repository, 0)
	iter := garClient.ListRepositories(ctx, &artifactregistrypb.ListRepositoriesRequest{Parent: id.String()})

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

func ListDockerImages(ctx context.Context, garClient *artifactregistry.Client, repositoryID string) ([]*artifactregistrypb.DockerImage, error) {
	results := make([]*artifactregistrypb.DockerImage, 0)
	iter := garClient.ListDockerImages(ctx, &artifactregistrypb.ListDockerImagesRequest{Parent: repositoryID})

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
