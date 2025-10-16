package http

import (
	"encoding/json"
	"net/http"

	"github.com/alvarolspeixoto/go-weather-mcp/internal/app/weather"
)

type Handler struct {
	service *weather.Service
}

func NewHandler(service *weather.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetWeatherByCity(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	if city == "" {
		http.Error(w, "city query parameter is required", http.StatusBadRequest)
		return
	}

	weather, err := h.service.GetWeatherByCityName(r.Context(), city)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weather)
}
