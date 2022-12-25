package handler

import (
	"fmt"
	"log"
	"net/http"

	weatherapi "github.com/AlexKomzzz/weather_api"
	"github.com/gin-gonic/gin"
)

const (
	urlCityData              = "http://api.openweathermap.org/geo/1.0/direct?q=%s&limit=5&appid=%s"
	urlCurrentWeatherByCity  = "https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&units=metric&appid=%s"
	urlCurrentWeatherByCity2 = "https://api.openweathermap.org/data/2.5/weather?q=%s,%s&APPID=%s" // https://api.openweathermap.org/data/2.5/weather?q=London,uk&APPID=2fc5eae2a7d9233fa9e282951d71139d
	urlFiveDayWeatherByCity  = "https://api.openweathermap.org/data/2.5/forecast?lat=%f&lon=%f&units=metric&appid=%s"
)

func (h *Handler) GetCurrentWeather(c *gin.Context) {

	// спарсить данные из тела запроса в структуру
	var cityInput weatherapi.City
	if err := c.BindJSON(&cityInput); err != nil {
		log.Println("error request? not found data by city: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// после получения координат города делаем запрос на прогноз погоды
	URLweatherToCity := fmt.Sprintf(urlCurrentWeatherByCity, cityInput.Lat, cityInput.Lon, h.idAPI)

	resp2, err := http.Get(URLweatherToCity)
	if err != nil {
		log.Println("error request2 by weather: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer resp2.Body.Close()

	weather, err := h.services.DecodingCurrentWeatherBodyJSON(resp2.Body)
	if err != nil {
		log.Println("error decoding data weaather: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"current_weather": weather,
	})
}

func (h *Handler) GetFiveDayWeather(c *gin.Context) {

	// спарсить данные из тела запроса в структуру
	var input weatherapi.City
	if err := c.BindJSON(&input); err != nil {
		log.Println("error request? not found data by city: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	URLlocal := fmt.Sprintf(urlFiveDayWeatherByCity, input.Lat, input.Lon, h.idAPI)

	resp, err := http.Get(URLlocal)
	if err != nil {
		log.Println("error request by local: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	defer resp.Body.Close()

	weather, err := h.services.DecodingFiveDayWeatherBodyJSON(resp.Body)
	if err != nil {
		log.Println("error decoding data 5 days weather: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"5_days_weather": weather,
	})
}
