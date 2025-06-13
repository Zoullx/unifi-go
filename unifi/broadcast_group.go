package unifi

import (
	"context"
)

func (c *client) ListBroadcastGroup(ctx context.Context, site string) ([]BroadcastGroup, error) {
	return c.listBroadcastGroup(ctx, site)
}

func (c *client) CreateBroadcastGroup(ctx context.Context, site string, d *BroadcastGroup) (*BroadcastGroup, error) {
	return c.createBroadcastGroup(ctx, site, d)
}

func (c *client) GetBroadcastGroup(ctx context.Context, site, id string) (*BroadcastGroup, error) {
	return c.getBroadcastGroup(ctx, site, id)
}

func (c *client) GetBroadcastGroupByName(ctx context.Context, site, name string) (*BroadcastGroup, error) {
	broadcastGroups, err := c.listBroadcastGroup(ctx, site)
	if err != nil {
		return nil, err
	}

	for _, b := range broadcastGroups {
		if b.Name == name {
			return &b, nil
		}
	}

	return nil, ErrNotFound
}

func (c *client) DeleteBroadcastGroup(ctx context.Context, site, id string) error {
	return c.deleteBroadcastGroup(ctx, site, id)
}

func (c *client) UpdateBroadcastGroup(ctx context.Context, site string, d *BroadcastGroup) (*BroadcastGroup, error) {
	return c.updateBroadcastGroup(ctx, site, d)
}
