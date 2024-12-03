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

type Location struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
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

func (c *Client) GetLocation(name string) (Location, error) {
	url := baseURL + "/location-area/" + name

	// create request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, fmt.Errorf("error creating request: %w", err)
	}

	requestURL := req.URL.String()

	// check the cache for this url
	if val, ok := c.Cache.Get(requestURL); ok {
		locationsResp := Location{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return Location{}, err
		}
		fmt.Println("[PAGE FOUND IN CACHE]")
		return locationsResp, nil
	}

	// make request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return Location{}, fmt.Errorf("error sending request: %w", err)
	}
	defer res.Body.Close()

	// get body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Location{}, fmt.Errorf("error reading response body: %w", err)
	}

	// parse response into struct
	var response Location
	err = json.Unmarshal(body, &response)
	if err != nil {
		return Location{}, err
	}

	c.Cache.Add(requestURL, body) // cache response for this url
	return response, nil
}
