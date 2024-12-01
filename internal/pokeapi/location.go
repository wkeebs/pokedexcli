package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationResponse struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []LocationArea `json:"results"`
}

func (c *Client) ListLocations(limit int, index int) (LocationResponse, error) {
	url := baseURL + "/location-area/"

	// create request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationResponse{}, fmt.Errorf("error creating request: %w", err)
	}

	// add query params
	q := req.URL.Query()
	q.Add("limit", strconv.Itoa(limit))
	q.Add("offset", strconv.Itoa(index*limit))
	req.URL.RawQuery = q.Encode()

	requestURL := req.URL.String()

	// check the cache for this url
	if val, ok := c.Cache.Get(requestURL); ok {
		locationsResp := LocationResponse{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return LocationResponse{}, err
		}
		fmt.Println("[PAGE FOUND IN CACHE]")
		return locationsResp, nil
	}

	// make request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return LocationResponse{}, fmt.Errorf("error sending request: %w", err)
	}
	defer res.Body.Close()

	// get body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationResponse{}, fmt.Errorf("error reading response body: %w", err)
	}

	// parse response into struct
	var response LocationResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return LocationResponse{}, err
	}

	c.Cache.Add(requestURL, body) // cache response for this url
	return response, nil
}
