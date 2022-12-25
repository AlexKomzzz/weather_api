package handler

import (
	"github.com/AlexKomzzz/weather_api/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitServ() *gin.Engine {
	mux := gin.New()

	mux.GET("/weather/:city/:country", h.GetWeather)

	return mux
}
