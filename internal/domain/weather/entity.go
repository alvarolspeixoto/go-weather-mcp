package weather

type Weather struct {
	CityName    string
	Country     string
	Temperature float64
	FeelsLike   float64
	TempMin     float64
	TempMax     float64
	Humidity    int
	Pressure    int
	Description string
	WindSpeed   float64
	Cloudiness  int
	Timestamp   int64
}
