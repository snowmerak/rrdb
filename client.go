package rrdb

import (
	"context"

	"github.com/rueian/rueidis"
)

type Client struct {
	ctx    context.Context
	client rueidis.Client
}

func NewClient(ctx context.Context, username, passwd string, addr ...string) (*Client, error) {
	c, err := rueidis.NewClient(rueidis.ClientOption{
		Username:    username,
		Password:    passwd,
		InitAddress: addr,
	})
	if err != nil {
		return nil, err
	}
	return &Client{ctx, c}, nil
}
