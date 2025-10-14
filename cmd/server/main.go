package main

import (
	"fmt"

	"github.com/alvarolspeixoto/go-weather-mcp/internal/config"
)

func main() {
	config.Load()
	fmt.Println("Servidor iniciado")
}
