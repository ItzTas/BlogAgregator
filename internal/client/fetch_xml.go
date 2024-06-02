package client

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) FetchXML(url string) (*RSS, error) {
	if dat, ok := c.cache.Get(url); ok {
		rss := RSS{}
		err := xml.Unmarshal(dat, &rss)
		if err != nil {
			return nil, err
		}
		return &rss, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return nil, fmt.Errorf("unexpected status code: %v", res.StatusCode)
	}

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	rss := RSS{}
	err = xml.Unmarshal(dat, &rss)
	if err != nil {
		return nil, err
	}

	c.cache.Add(url, dat)
	return &rss, nil
}
