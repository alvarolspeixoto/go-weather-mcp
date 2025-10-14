package weather

import (
	"github.com/alvarolspeixoto/go-weather-mcp/internal/domain/weather"
	"github.com/alvarolspeixoto/go-weather-mcp/internal/infra/openweathermap"
)

func ToDomain(resp *openweathermap.OWWeatherResponse) *weather.Weather {
	mainDesc := ""
	if len(resp.Weather) > 0 {
		mainDesc = resp.Weather[0].Description
	}

	return &weather.Weather{
		CityName:    resp.Name,
		Country:     resp.Sys.Country,
		Temperature: resp.Main.Temp,
		FeelsLike:   resp.Main.FeelsLike,
		TempMin:     resp.Main.TempMin,
		TempMax:     resp.Main.TempMax,
		Humidity:    resp.Main.Humidity,
		Pressure:    resp.Main.Pressure,
		Description: mainDesc,
		WindSpeed:   resp.Wind.Speed,
		Cloudiness:  resp.Clouds.All,
		Timestamp:   resp.Dt,
	}
}
