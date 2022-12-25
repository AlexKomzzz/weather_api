package service

import (
	"encoding/json"
	"io"
	"log"

	weatherapi "github.com/AlexKomzzz/weather_api"
)

func (s *Service) GetCity(cityArr []weatherapi.City, country string) weatherapi.City {
	// если страну передали, то из выданных городов выбираем город, этой страны
	// если страну не передали, то выбираем первый город
	if len(cityArr) == 0 {
		return weatherapi.City{}
	}

	if country != "" {
		// проверка, что у нас в мапе есть такая страна
		countryUser, ok := s.countryMap[country]
		if ok {
			for _, city := range cityArr {
				if city.Country == countryUser {
					log.Println("Есть город в такой стране: ", city)
					return city
				}
			}
		}
	}

	log.Println("Страна не задана: ", cityArr[0])

	return cityArr[0]
}

func (s *Service) DecodingCityBodyJSON(body io.ReadCloser) ([]weatherapi.City, error) {
	var cityArr []weatherapi.City

	data, err := io.ReadAll(body)
	if err != nil {
		log.Println("not found responce by api: ", err)
		return nil, err
	}

	err = json.Unmarshal(data, &cityArr)
	if err != nil {
		log.Println("error unmarshal: ", err)
		return nil, err
	}

	return cityArr, nil
}
