package domain

type CurrentWeather struct {
	Temperature float64 `json:"temp"`
	Humidity    int     `json:"humidity"`
	WindSpeed   float64 `json:"wind_speed"`
	Weather     []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

type Alert struct {
	SenderName  string `json:"sender_name"`
	Event       string `json:"event"`
	Start       int64  `json:"start"`
	End         int64  `json:"end"`
	Description string `json:"description"`
}

type WeatherResponse struct {
	Latitude       float64        `json:"lat"`
	Longitude      float64        `json:"lon"`
	Timezone       string         `json:"timezone"`
	CurrentWeather CurrentWeather `json:"current"`
	Alerts         []Alert        `json:"alerts,omitempty"`
}
