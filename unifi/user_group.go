package unifi

import "context"

func (c *client) ListUserGroup(ctx context.Context, site string) ([]UserGroup, error) {
	return c.listUserGroup(ctx, site)
}

func (c *client) GetUserGroup(ctx context.Context, site, id string) (*UserGroup, error) {
	return c.getUserGroup(ctx, site, id)
}

func (c *client) GetUserGroupByName(ctx context.Context, site, name string) (*UserGroup, error) {
	userGroups, err := c.listUserGroup(ctx, site)
	if err != nil {
		return nil, err
	}

	for _, u := range userGroups {
		if u.Name == name {
			return &u, nil
		}
	}

	return nil, ErrNotFound
}

func (c *client) DeleteUserGroup(ctx context.Context, site, id string) error {
	return c.deleteUserGroup(ctx, site, id)
}

func (c *client) CreateUserGroup(ctx context.Context, site string, d *UserGroup) (*UserGroup, error) {
	return c.createUserGroup(ctx, site, d)
}

func (c *client) UpdateUserGroup(ctx context.Context, site string, d *UserGroup) (*UserGroup, error) {
	return c.updateUserGroup(ctx, site, d)
}
