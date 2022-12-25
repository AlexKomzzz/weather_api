package service

import (
	"log"

	weatherapi "github.com/AlexKomzzz/weather_api"
)

func (s *Service) GetCity(cityArr []weatherapi.City, country string) weatherapi.City {
	// если страну передали, то из выданных городов выбираем город, этой страны
	// если страну не передали, то выбираем первый город

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

	log.Println("Нет такой страны в мапе: ", cityArr[0])
	return cityArr[0]

}
