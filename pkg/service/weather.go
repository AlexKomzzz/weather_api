package service

import (
	"encoding/json"
	"io"
	"log"

	weatherapi "github.com/AlexKomzzz/weather_api"
)

func (s *Service) DecodingCurrentWeatherBodyJSON(body io.ReadCloser) (weatherapi.CurrentWeatherData, error) {
	var weatherData weatherapi.CurrentWeatherData

	data, err := io.ReadAll(body)
	if err != nil {
		log.Println("not found responce by api: ", err)
		return weatherData, err
	}

	log.Println("Data RESP: ", string(data))

	err = json.Unmarshal(data, &weatherData)
	if err != nil {
		log.Println("error unmarshal: ", err)
		return weatherData, err
	}

	return weatherData, nil
}

func (s *Service) DecodingFiveDayWeatherBodyJSON(body io.ReadCloser) (weatherapi.FiveDaysWeatherData, error) {
	var weatherData weatherapi.FiveDaysWeatherData

	data, err := io.ReadAll(body)
	if err != nil {
		log.Println("not found responce by api: ", err)
		return weatherData, err
	}

	log.Println("Data RESP: ", string(data))

	err = json.Unmarshal(data, &weatherData)
	if err != nil {
		log.Println("error unmarshal: ", err)
		return weatherData, err
	}

	return weatherData, nil
}
