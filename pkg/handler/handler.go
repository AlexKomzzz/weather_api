package handler

import (
	"github.com/AlexKomzzz/weather_api/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
	idAPI    string
}

func NewHandler(services *service.Service, idAPI string) *Handler {
	return &Handler{
		services: services,
		idAPI:    idAPI,
	}
}

func (h *Handler) InitServ() *gin.Engine {
	mux := gin.New()

	city := mux.Group("/city")
	{
		city.GET("/:city", h.GetCityCoord)
		city.GET("/:city/:country", h.GetCityCoordByCountry)
	}

	weather := mux.Group("/weather")
	{
		/* тело запроса:
		{
			"name": "Moscow",
			 "lat": 55.7504461,
			 "lon": 37.6174943
		}*/
		weather.POST("/current", h.GetCurrentWeather)
		weather.POST("/five-days", h.GetFiveDayWeather)
	}

	return mux
}
