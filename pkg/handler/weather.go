package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

const (
	urlLocalNoID      = "http://api.openweathermap.org/geo/1.0/direct?q=%s&limit=5&appid=%s"
	urlWeatherByCity  = "https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&units=metric&appid=%s"
	urlWeatherByCity2 = "https://api.openweathermap.org/data/2.5/weather?q=%s,%s&APPID=%s" // https://api.openweathermap.org/data/2.5/weather?q=London,uk&APPID=2fc5eae2a7d9233fa9e282951d71139d
)

func (h *Handler) GetWeather(c *gin.Context) {
	// достать из URL название города
	nameCity := c.Param("city")
	log.Println("name city: ", nameCity)
	// отправить запрос в АПИ для получения координат города

	idAPI, ok := os.LookupEnv("idapi")
	if !ok {
		log.Println("not found IdAPI")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "not found ID API",
		})
		return
	}

	URLlocal := fmt.Sprintf(urlLocalNoID, nameCity, idAPI)

	resp, err := http.Get(URLlocal)
	if err != nil {
		log.Println("error request by local: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	defer resp.Body.Close()

	cityArr, err := h.services.DecodingCityBodyJSON(resp.Body)
	if err != nil {
		log.Println("error decoding data city: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// достать из URL название страны
	nameCountry := c.Param("country")
	log.Println("name country: ", nameCountry)

	// log.Println("city:/n", city)
	city := h.services.GetCity(cityArr, nameCountry)

	// после получения координат города делаем запрос на прогноз погоды
	URLweatherToCity := fmt.Sprintf(urlWeatherByCity, city.Lat, city.Lon, idAPI)

	resp2, err := http.Get(URLweatherToCity)
	if err != nil {
		log.Println("error request2 by weather: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer resp2.Body.Close()

	weather, err := h.services.DecodingWeatherBodyJSON(resp2.Body)
	if err != nil {
		log.Println("error decoding data weaather: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"city":    city,
		"weather": weather,
	})
}
