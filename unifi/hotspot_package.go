package unifi

import (
	"context"
)

func (c *client) ListHotspotPackage(ctx context.Context, site string) ([]HotspotPackage, error) {
	return c.listHotspotPackage(ctx, site)
}

func (c *client) GetHotspotPackage(ctx context.Context, site, id string) (*HotspotPackage, error) {
	return c.getHotspotPackage(ctx, site, id)
}

func (c *client) GetHotspotPackageByName(ctx context.Context, site, name string) (*HotspotPackage, error) {
	hotspotPackage, err := c.listHotspotPackage(ctx, site)
	if err != nil {
		return nil, err
	}

	for _, h := range hotspotPackage {
		if h.Name == name {
			return &h, nil
		}
	}

	return nil, ErrNotFound
}

func (c *client) DeleteHotspotPackage(ctx context.Context, site, id string) error {
	return c.deleteHotspotPackage(ctx, site, id)
}

func (c *client) CreateHotspotPackage(ctx context.Context, site string, d *HotspotPackage) (*HotspotPackage, error) {
	return c.createHotspotPackage(ctx, site, d)
}

func (c *client) UpdateHotspotPackage(ctx context.Context, site string, d *HotspotPackage) (*HotspotPackage, error) {
	return c.updateHotspotPackage(ctx, site, d)
}
