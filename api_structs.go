package weatherapi

type City struct {
	Name    string  `json:"name"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
	Country string  `json:"country"`
}

func InitCountryMap() map[string]string {
	CountyMap := map[string]string{"Россия": "RU", "Russia": "RU", "США": "US", "USA": "US", "Франция": "FR"}
	return CountyMap
}
