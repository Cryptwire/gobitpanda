package gobitpanda

import (
	"errors"
	"fmt"
	"time"
)

// GetPriceTicksByCode gets price ticks for specific market for interval of maximum of 4 hours. Sorted by latest first
func (c *Client) GetPriceTicksByCode(instrumentCode string, from time.Time, to time.Time) (*[]PriceTick, error) {
	priceTicks := &[]PriceTick{}

	if instrumentCode == "" {
		return nil, errors.New("instrumentCode can not be empty")
	}

	var params string

	if !from.IsZero() {
		if params == "" {
			params += "?from=" + from.UTC().Format(time.RFC3339)
		} else {
			params += "&from=" + from.UTC().Format(time.RFC3339)
		}
	}

	if !to.IsZero() {
		if params == "" {
			params += "?to=" + to.UTC().Format(time.RFC3339)
		} else {
			params += "&to=" + to.UTC().Format(time.RFC3339)
		}
	}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s%s%s", c.APIBase, "/v1/price-ticks/", instrumentCode, params), nil)
	if err != nil {
		return priceTicks, err
	}

	if err = c.Send(req, priceTicks); err != nil {
		return priceTicks, err
	}

	return priceTicks, nil
}
