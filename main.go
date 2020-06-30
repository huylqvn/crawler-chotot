package main

import (
	"crawl-data/config"
	crawl "crawl-data/crawler"
	"crawl-data/db"
	"crawl-data/service"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

func main() {
	logger := log.New()

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	var cfg config.Configuration

	if err := viper.ReadInConfig(); err != nil {
		logger.Error("Error reading config file", err)
	}
	err := viper.Unmarshal(&cfg)
	if err != nil {
		logger.Error("unable to decode into struct", err)
	}

	logger.Info("database URL ", cfg.Database.Host, cfg.Database.Username)
	logger.Info("Port service ", cfg.Server.Port)

	service := service.Service{
		Logger: logger,
		DB: db.SetupDB(cfg.Database.Host,
			cfg.Database.Username,
			cfg.Database.Password,
			cfg.Database.Database,
			cfg.Database.Port),
	}

	err = crawl.StartCrawlerChotot(&service)
	if err != nil {
		logger.Error("Crawl fail", err)
	}
}
