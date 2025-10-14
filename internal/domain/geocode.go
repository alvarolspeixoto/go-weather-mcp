package domain

type GeocodeResponse struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
	Timezone  string  `json:"timezone"`
	Country   string  `json:"country"`
	State     string  `json:"state,omitempty"`
}
