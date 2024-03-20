package naisapi

import (
	"context"
	"fmt"

	"github.com/nais/garbage/internal/naisapi/protoapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	client protoapi.ClusterClient
}

func NewClient(target string, insecureConnection bool) (*Client, error) {
	opts := []grpc.DialOption{}
	if insecureConnection {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	gclient, err := grpc.Dial(target, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to nais-api: %w", err)
	}

	return &Client{
		client: protoapi.NewClusterClient(gclient),
	}, nil
}

func (c *Client) ListContainerImagesInUse(ctx context.Context) ([]*protoapi.ContainerImage, error) {
	resp, err := c.client.ListContainerImagesInUse(ctx, &protoapi.ListContainerImagesInUseRequest{})
	if err != nil {
		return nil, err
	}

	return resp.Images, nil
}
