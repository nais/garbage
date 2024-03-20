package gcp

import (
	"context"
	"fmt"

	artifactregistry "cloud.google.com/go/artifactregistry/apiv1"
	"google.golang.org/api/option"
)

func ArtifactRegistryClient(ctx context.Context, serviceAccountEmail string) (*artifactregistry.Client, error) {
	ts, err := ImpersonatedTokenSource(ctx, serviceAccountEmail)
	if err != nil {
		return nil, fmt.Errorf("create token source: %w", err)
	}

	return artifactregistry.NewClient(ctx, option.WithTokenSource(ts))
}
