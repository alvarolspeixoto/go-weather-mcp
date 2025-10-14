package openweathermap

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	APIKey         string
	WeatherBaseURL string
	GeocodeBaseURL string
	HTTPClient     *http.Client
}

func NewClient(apiKey string, weatherBaseURL string, geocodeBaseURL string) *Client {
	return &Client{
		APIKey:         apiKey,
		WeatherBaseURL: weatherBaseURL,
		GeocodeBaseURL: geocodeBaseURL,
		HTTPClient:     &http.Client{},
	}
}

func (c *Client) GetGeocode(city string) (*OWGeocode, error) {
	endpoint := fmt.Sprintf("%s?q=%s&limit=1&appid=%s", c.GeocodeBaseURL, url.QueryEscape(city), c.APIKey)
	resp, err := c.HTTPClient.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code from geocode API: %d", resp.StatusCode)
	}

	var results OWGeocodeResponse
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("no results found")
	}

	return &results[0], nil
}

func (c *Client) GetWeather(lat, lon float64) (*OWWeatherResponse, error) {
	endpoint := fmt.Sprintf("%s?lat=%f&lon=%f&appid=%s", c.WeatherBaseURL, lat, lon, c.APIKey)
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
