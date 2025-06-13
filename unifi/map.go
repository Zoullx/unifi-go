package unifi

import (
	"context"
)

func (c *client) ListMap(ctx context.Context, site string) ([]Map, error) {
	return c.listMap(ctx, site)
}

func (c *client) GetMap(ctx context.Context, site, id string) (*Map, error) {
	return c.getMap(ctx, site, id)
}

func (c *client) GetMapByName(ctx context.Context, site, name string) (*Map, error) {
	maps, err := c.listMap(ctx, site)
	if err != nil {
		return nil, err
	}

	for _, m := range maps {
		if m.Name == name {
			return &m, nil
		}
	}

	return nil, ErrNotFound
}

func (c *client) DeleteMap(ctx context.Context, site, id string) error {
	return c.deleteMap(ctx, site, id)
}

func (c *client) CreateMap(ctx context.Context, site string, d *Map) (*Map, error) {
	return c.createMap(ctx, site, d)
}

func (c *client) UpdateMap(ctx context.Context, site string, d *Map) (*Map, error) {
	return c.updateMap(ctx, site, d)
}
