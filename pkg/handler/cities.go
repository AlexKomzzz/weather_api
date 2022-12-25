package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetCityCoordByCountry(c *gin.Context) {

	// достать из URL название города
	nameCity := c.Param("city")
	log.Println("name city: ", nameCity)

	// достать из URL название страны
	nameCountry := c.Param("country")
	log.Println("name country: ", nameCountry)

	// отправить запрос в АПИ для получения координат города

	// вытащить API ID из среды окр
	idAPI, ok := os.LookupEnv("idapi")
	if !ok {
		log.Println("not found IdAPI")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "not found ID API",
		})
		return
	}

	// сформировать УРЛ.  Не самый удачный способ!!!
	URLlocal := fmt.Sprintf(urlCityData, nameCity, idAPI)

	// отправить запрос
	resp, err := http.Get(URLlocal)
	if err != nil {
		log.Println("error request by local: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer resp.Body.Close()

	// спарсить данные из тела ответа в структуру
	cityArr, err := h.services.DecodingCityBodyJSON(resp.Body)
	if err != nil {
		log.Println("error decoding data city: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// log.Println("city:/n", city)
	city := h.services.GetCity(cityArr, nameCountry)

	c.JSON(http.StatusOK, gin.H{
		"city": city,
	})
}

func (h *Handler) GetCityCoord(c *gin.Context) {

	// достать из URL название города
	nameCity := c.Param("city")
	log.Println("name city: ", nameCity)

	// отправить запрос в АПИ для получения координат города

	// вытащить API ID из среды окр
	idAPI, ok := os.LookupEnv("idapi")
	if !ok {
		log.Println("not found IdAPI")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "not found ID API",
		})
		return
	}

	// сформировать УРЛ.  Не самый удачный способ!!!
	URLlocal := fmt.Sprintf(urlCityData, nameCity, idAPI)

	// отправить запрос
	resp, err := http.Get(URLlocal)
	if err != nil {
		log.Println("error request by local: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer resp.Body.Close()

	// спарсить данные из тела ответа в структуру
	cityArr, err := h.services.DecodingCityBodyJSON(resp.Body)
	if err != nil {
		log.Println("error decoding data city: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// log.Println("city:/n", city)
	city := h.services.GetCity(cityArr, "")

	c.JSON(http.StatusOK, gin.H{
		"city": city,
	})
}
