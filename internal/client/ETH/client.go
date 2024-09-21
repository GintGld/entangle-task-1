package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"task1/internal/models"
)

type Client struct {
	url string
	c   *http.Client
}

func New(url string) *Client {
	return &Client{
		url: url,
		c:   http.DefaultClient,
	}
}

// Amount returns NGLAmountResp.
func (c *Client) Amount(ctx context.Context) (models.NGLAmountResp, error) {
	const op = "Client.Amount"

	req, err := http.NewRequest(http.MethodGet, c.url, nil)
	if err != nil {
		return models.NGLAmountResp{}, fmt.Errorf("%s: %w", op, err)
	}

	resp, err := c.c.Do(req)
	if err != nil {
		return models.NGLAmountResp{}, fmt.Errorf("%s: %w", op, err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.NGLAmountResp{}, fmt.Errorf("%s: %w", op, err)
	}

	var amountResp models.NGLAmountResp
	if err := json.Unmarshal(body, &amountResp); err != nil {
		return models.NGLAmountResp{}, fmt.Errorf("%s: %w", op, err)
	}

	return amountResp, nil
}
