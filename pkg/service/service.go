package service

type Service struct {
	countryMap map[string]string
}

func NewService(countryMap map[string]string) *Service {
	return &Service{
		countryMap: countryMap,
	}
}
