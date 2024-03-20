package gcp

import (
	"context"

	"golang.org/x/oauth2"
	"google.golang.org/api/cloudresourcemanager/v3"
	"google.golang.org/api/impersonate"
)

func ImpersonatedTokenSource(ctx context.Context, serviceAccountEmail string) (oauth2.TokenSource, error) {
	return impersonate.CredentialsTokenSource(ctx, impersonate.CredentialsConfig{
		Scopes: []string{
			cloudresourcemanager.CloudPlatformScope,
		},
		TargetPrincipal: serviceAccountEmail,
	})
}
