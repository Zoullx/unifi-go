package unifi

import (
	"context"
	"fmt"
	"reflect"
)

type Site struct {
	ID string `json:"_id,omitempty"`

	// Hidden   bool   `json:"attr_hidden,omitempty"`
	// HiddenID string `json:"attr_hidden_id,omitempty"`
	// NoDelete bool   `json:"attr_no_delete,omitempty"`
	// NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Name        string `json:"name"`
	Description string `json:"desc"`

	// Role string `json:"role"`
}

func (c *client) ListSites(ctx context.Context) ([]Site, error) {
	var respBody struct {
		Meta Meta   `json:"Meta"`
		Data []Site `json:"data"`
	}

	err := c.Get(ctx, "self/sites", nil, &respBody)
	if err != nil {
		return nil, err
	}

	return respBody.Data, nil
}

func (c *client) GetSite(ctx context.Context, id string) (*Site, error) {
	sites, err := c.ListSites(ctx)
	if err != nil {
		return nil, err
	}

	for _, s := range sites {
		if s.ID == id {
			return &s, nil
		}
	}

	return nil, ErrNotFound
}

func (c *client) CreateSite(ctx context.Context, description string) ([]Site, error) {
	reqBody := struct {
		Cmd  string `json:"cmd"`
		Desc string `json:"desc"`
	}{
		Cmd:  "add-site",
		Desc: description,
	}

	var respBody struct {
		Meta Meta   `json:"Meta"`
		Data []Site `json:"data"`
	}

	err := c.Post(ctx, "s/default/cmd/sitemgr", reqBody, &respBody)
	if err != nil {
		return nil, err
	}

	return respBody.Data, nil
}

func (c *client) CreateSiteByModel(ctx context.Context, d *Site) (*Site, error) {
	reqBody := map[string]interface{}{}

	// Use reflection to transfer key value pairs from Site to reqBody
	// so we can add the add-site cmd
	v := reflect.ValueOf(d)
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		reqBody[t.Field(i).Name] = v.Field(i)
	}

	reqBody["cmd"] = "add-site"

	var respBody struct {
		Meta Meta   `json:"meta"`
		Data []Site `json:"data"`
	}

	err := c.Post(ctx, "s/default/cmd/sitemgr", reqBody, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, ErrNotFound
	}

	new := respBody.Data[0]

	return &new, nil
}

func (c *client) DeleteSite(ctx context.Context, id string) ([]Site, error) {
	reqBody := struct {
		Cmd  string `json:"cmd"`
		Site string `json:"site"`
	}{
		Cmd:  "delete-site",
		Site: id,
	}

	var respBody struct {
		Meta Meta   `json:"Meta"`
		Data []Site `json:"data"`
	}

	err := c.Post(ctx, "s/default/cmd/sitemgr", reqBody, &respBody)
	if err != nil {
		return nil, err
	}

	return respBody.Data, nil
}

func (c *client) UpdateSite(ctx context.Context, name, description string) ([]Site, error) {
	reqBody := struct {
		Cmd  string `json:"cmd"`
		Desc string `json:"desc"`
	}{
		Cmd:  "update-site",
		Desc: description,
	}

	var respBody struct {
		Meta Meta   `json:"Meta"`
		Data []Site `json:"data"`
	}

	err := c.Post(ctx, fmt.Sprintf("s/%s/cmd/sitemgr", name), reqBody, &respBody)
	if err != nil {
		return nil, err
	}

	return respBody.Data, nil
}

func (c *client) UpdateSiteByModel(ctx context.Context, d *Site) (*Site, error) {
	reqBody := map[string]interface{}{}

	// Use reflection to transfer key value pairs from Site to reqBody
	// so we can add the add-site cmd
	v := reflect.ValueOf(d)
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		reqBody[t.Field(i).Name] = v.Field(i)
	}

	reqBody["cmd"] = "add-site"

	var respBody struct {
		Meta Meta   `json:"meta"`
		Data []Site `json:"data"`
	}

	err := c.Post(ctx, "s/default/cmd/sitemgr", reqBody, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, ErrNotFound
	}

	new := respBody.Data[0]

	return &new, nil
}
