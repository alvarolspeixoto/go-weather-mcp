package openweathermap

type OWCoord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type OWWeather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type OWMain struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
}

type OWWind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
	Gust  float64 `json:"gust"`
}

type OWRain struct {
	OneHour float64 `json:"1h"`
}

type OWClouds struct {
	All int `json:"all"`
}

type OWSys struct {
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"`
}

type OWWeatherResponse struct {
	Coord      OWCoord     `json:"coord"`
	Weather    []OWWeather `json:"weather"`
	Base       string      `json:"base"`
	Main       OWMain      `json:"main"`
	Visibility int         `json:"visibility"`
	Wind       OWWind      `json:"wind"`
	Rain       OWRain      `json:"rain,omitempty"`
	Clouds     OWClouds    `json:"clouds"`
	Dt         int64       `json:"dt"`
	Sys        OWSys       `json:"sys"`
	Timezone   int         `json:"timezone"`
	ID         int         `json:"id"`
	Name       string      `json:"name"`
	Cod        int         `json:"cod"`
}

type OWGeocode struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
	Country   string  `json:"country"`
	State     string  `json:"state,omitempty"`
}

type OWGeocodeResponse []OWGeocode
