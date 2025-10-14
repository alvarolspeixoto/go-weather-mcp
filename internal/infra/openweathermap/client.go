package openweathermap

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	APIKey     string
	BaseURL    string
	HTTPClient *http.Client
}

func NewClient(apiKey string, url string) *Client {
	return &Client{
		APIKey:     apiKey,
		BaseURL:    url,
		HTTPClient: &http.Client{},
	}
}

func (c *Client) GetGeocode(city string) ([]OWGeocodeResponse, error) {
	endpoint := fmt.Sprintf("%s?q=%s&limit=1&appid=%s", c.BaseURL, url.QueryEscape(city), c.APIKey)
	resp, err := c.HTTPClient.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code from geocode API: %d", resp.StatusCode)
	}

	var results []OWGeocodeResponse
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return nil, err
	}
	return results, nil
}

func (c *Client) GetWeather(lat, lon float64) (*OWWeatherResponse, error) {
	endpoint := fmt.Sprintf("%s?lat=%f&lon=%f&appid=%s", c.BaseURL, lat, lon, c.APIKey)
	resp, err := c.HTTPClient.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code from weather API: %d", resp.StatusCode)
	}

	var result OWWeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
