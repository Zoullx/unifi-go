package unifi

// client for dashboard.generated.go

import (
	"context"
)

func (c *client) ListDashboard(ctx context.Context, site string) ([]Dashboard, error) {
	return c.listDashboard(ctx, site)
}

func (c *client) GetDashboard(ctx context.Context, site, id string) (*Dashboard, error) {
	return c.getDashboard(ctx, site, id)
}

func (c *client) GetDashboardByName(ctx context.Context, site, name string) (*Dashboard, error) {
	dashboards, err := c.listDashboard(ctx, site)
	if err != nil {
		return nil, err
	}

	for _, d := range dashboards {
		if d.Name == name {
			return &d, nil
		}
	}

	return nil, ErrNotFound
}

func (c *client) DeleteDashboard(ctx context.Context, site, id string) error {
	return c.deleteDashboard(ctx, site, id)
}

func (c *client) CreateDashboard(ctx context.Context, site string, d *Dashboard) (*Dashboard, error) {
	return c.createDashboard(ctx, site, d)
}

func (c *client) UpdateDashboard(ctx context.Context, site string, d *Dashboard) (*Dashboard, error) {
	return c.updateDashboard(ctx, site, d)
}
