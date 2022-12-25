package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	weatherapi "github.com/AlexKomzzz/weather_api"
	"github.com/gin-gonic/gin"
)

const (
	urlLocalNoID    = "http://api.openweathermap.org/geo/1.0/direct?q=Moscow&limit=5&appid=%s"
	urlWetherByCity = ""
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

	URLlocal := fmt.Sprintf(urlLocalNoID, idAPI)

	resp, err := http.Get(URLlocal)
	if err != nil {
		log.Println("error request by local: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	defer resp.Body.Close()

	var cityArr []weatherapi.City

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("not found responce by api: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = json.Unmarshal(data, &cityArr)
	if err != nil {
		log.Println("error unmarshal: ", err)
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

	c.JSON(http.StatusOK, gin.H{
		"city": city,
	})
}
