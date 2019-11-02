package gobitpanda

import (
	"errors"
	"fmt"
)

// GetMarketTicker returns statistics on all available instruments
func (c *Client) GetMarketTicker() (*[]MarketTick, error) {
	marketTicks := &[]MarketTick{}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.APIBase, "/v1/market-ticker"), nil)
	if err != nil {
		return marketTicks, err
	}

	if err = c.Send(req, marketTicks); err != nil {
		return marketTicks, err
	}

	return marketTicks, nil
}

// GetMarketTickerByCode gets statistics on a single market
func (c *Client) GetMarketTickerByCode(instrumentCode string) (*MarketTick, error) {
	if instrumentCode == "" {
		return nil, errors.New("instrumentCode can not be empty")
	}

	marketTick := &MarketTick{}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s%s", c.APIBase, "/v1/market-ticker/", instrumentCode), nil)
	if err != nil {
		return marketTick, err
	}

	if err = c.Send(req, marketTick); err != nil {
		return marketTick, err
	}

	return marketTick, nil
}
