package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocation(location string) (LocationAreaResult, error) {
	url := baseURL + "/location-area/" + location

	if entry, ok := c.cache.Get(url); ok {
		result := LocationAreaResult{}
		err := json.Unmarshal(entry, &result)
		if err != nil {
			return result, err
		}
		return result, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaResult{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResult{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResult{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)
	}

	result := LocationAreaResult{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return LocationAreaResult{}, err
	}

	c.cache.Add(url, body)
	return result, nil
}
