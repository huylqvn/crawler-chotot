package service

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

type Service struct {
	DB     *gorm.DB
	Logger *log.Logger
}
