package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	weatherapi "github.com/AlexKomzzz/weather_api"
	"github.com/AlexKomzzz/weather_api/pkg/handler"
	"github.com/AlexKomzzz/weather_api/pkg/repository"
	"github.com/AlexKomzzz/weather_api/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

const port = ":9090"

func main() {
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
	})
	if err != nil {
		log.Fatalln("failed to initialize db: ", err)
		return
	}
	defer db.Close()

	idAPI, ok := os.LookupEnv("idapi")
	if !ok {
		log.Fatal("not found IdAPI")
	}

	countryMap := weatherapi.InitCountryMap()

	service := service.NewService(countryMap)
	h := handler.NewHandler(service, idAPI)
	serv := h.InitServ()

	go func() {
		if err := serv.Run(port); err != nil {
			log.Fatal("erver crash: ", err)
		}
	}()

	log.Println("Server Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	log.Println("Server Stopted")
}

func init() {
	// Инициализируем конфигурации
	if err := initConfig(); err != nil {
		log.Fatalln("error initializing configs: ", err)
		return
	}

	// загрузка переменных окружения из файла .env
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

// Инициализация конфигураций
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
