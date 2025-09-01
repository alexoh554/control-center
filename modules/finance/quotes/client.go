package quotes

import (
	"context"
)

func NewClient() *Client {
	return &Client{}
}

type Client struct {
}

// Returns a list of Quote structs given a list of symbols
func (c *Client) GetBySymbols(ctx context.Context, symbols []string) []Quote {
	return nil
}
