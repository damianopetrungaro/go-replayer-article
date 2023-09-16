package user

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	BaseClient *http.Client
	BaseURL    string
}

func (c *Client) Get(ctx context.Context, id string) (*User, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/users/%s", c.BaseURL, id), http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("%w: could not create request: %s", ErrGet, err)
	}

	res, err := c.BaseClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%w: could not do request: %s", ErrGet, err)
	}
	defer res.Body.Close()

	switch {
	case res.StatusCode == http.StatusNotFound:
		return nil, fmt.Errorf("%w: %s", ErrNotFound, err)
	case res.StatusCode != http.StatusOK:
		return nil, fmt.Errorf("%w: could not handle status code: %s", ErrGet, err)
	}

	var u User
	if err := json.NewDecoder(res.Body).Decode(&u); err != nil {
		return nil, fmt.Errorf("%w: could not decode data: %s", ErrGet, err)
	}

	return &u, nil
}

func (c *Client) Save(ctx context.Context, u *User) error {
	body := &bytes.Buffer{}
	if err := json.NewEncoder(body).Encode(u); err != nil {
		return fmt.Errorf("%w: could not encode data: %s", ErrSave, err)

	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("%s/users", c.BaseURL), body)
	if err != nil {
		return fmt.Errorf("%w: could not create request: %s", ErrSave, err)
	}

	res, err := c.BaseClient.Do(req)
	if err != nil {
		return fmt.Errorf("%w: could not do request: %s", ErrSave, err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("%w: could not handle status code: %d", ErrSave, res.StatusCode)
	}

	return nil
}
