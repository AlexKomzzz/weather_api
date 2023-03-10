package weatherapi

type City struct {
	Coord      `json:"coord"`
	Name       string  `json:"name"`
	Lat        float64 `json:"lat"`
	Lon        float64 `json:"lon"`
	Country    string  `json:"country"`
	Population int     `json:"population"`
}

type Coord struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type Weather struct {
	Id          int    `json:"id"`          // Идентификатор погодных условий
	MainWeath   string `json:"main"`        // Группа погодных параметров
	Description string `json:"description"` // Погодные условия в группе
}

type Main struct {
	Temp      float64 `json:"temp"`       // темп в цельсий
	TempLike  float64 `json:"feels_like"` // темп по ощущениям в С
	Pressure  int     `json:"pressure"`   // атм. давл, гПА
	Humidity  int     `json:"humidity"`   // влажн, %
	Temp_min  float64 `json:"temp_min"`
	Temp_max  float64 `json:"temp_max"`
	SeaLevel  int     `json:"sea_level"`  // Атмосферное давление на уровне моря, гПа
	GrndLevel int     `json:"grnd_level"` // Атмосферное давление на уровне земли, гПа
}

type Visibilitys struct {
	Visibility int `json:"visibility"` // видимость, метр
}

type Wind struct {
	Speed float64 `json:"speed"` // скорость, м/сек
	Deg   int     `json:"deg"`   // направление , град
	Gust  float64 `json:"gust"`  // порывы, м/сек
}

type Clouds struct {
	All int `json:"all"` // облачность, %
}

type Rain struct {
	OneHour   float64 `json:"1h"` // Объем дождя за последний 1 час, мм
	ThreeHour float64 `json:"3h"` // Объем дождя за последние 3 часа, мм
}

type Snow struct {
	OneHour   float64 `json:"1h"` // Объем снега за последний 1 час, мм
	ThreeHour float64 `json:"3h"` // Объем снега за последние 3 часа, мм
}

type Time struct {
	Dt int `json:"dt"` // Время расчета данных, unix, UTC
}

// структура текущей погоды
type CurrentWeatherData struct {
	Time
	Visibilitys
	Coord        `json:"coord"`
	Main         `json:"main"`
	Wind         `json:"wind"`
	Clouds       `json:"clouds"`
	Rain         `json:"rain"`
	Snow         `json:"snow"`
	NameCity     string    `json:"name"`
	WethersParam []Weather `json:"weather"`
}

type ListFiveDaysWeatherData struct {
	Time
	Visibilitys
	Main         `json:"main"`
	Clouds       `json:"clouds"`
	Wind         `json:"wind"`
	Rain         `json:"rain"`
	Snow         `json:"snow"`
	Pop          float64   `json:"pop"`    // Вероятность осадков, %
	Pod          string    `json:"pod"`    // Часть суток (n - ночь, d - день)
	DtTXT        string    `json:"dt_txt"` // Прогнозируемое время данных, ISO, UTC
	WethersParam []Weather `json:"weather"`
}

// структура погоды на пять дней
type FiveDaysWeatherData struct {
	City `json:"city"`
	List []ListFiveDaysWeatherData `json:"list"`
}

func InitCountryMap() map[string]string {
	CountyMap := map[string]string{"Россия": "RU", "Russia": "RU", "США": "US", "USA": "US", "Франция": "FR"}
	return CountyMap
}
