package http

import (
	"fmt"
	"log"
	"net/http"
)

func StartServer(handler *Handler, port int) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/weather", handler.GetWeatherByCity)

	addr := fmt.Sprintf(":%d", port)
	log.Printf("Servidor HTTP iniciado em http://localhost%s", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("Erro ao iniciar servidor HTTP: %v", err)
		return err
	}

	return nil
}
