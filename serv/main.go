package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	weatherapi "github.com/AlexKomzzz/weather_api"
	"github.com/AlexKomzzz/weather_api/pkg/handler"
	"github.com/AlexKomzzz/weather_api/pkg/service"
	"github.com/joho/godotenv"
)

const port = ":9090"

func main() {
	countryMap := weatherapi.InitCountryMap()

	service := service.NewService(countryMap)
	h := handler.NewHandler(service)
	serv := h.InitServ()

	go func() {
		if err := serv.Run(port); err != nil {
			log.Fatal("erver crash: ", err)
		}
	}()

	log.Println("Server Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	log.Println("Server Stopted")
}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}
