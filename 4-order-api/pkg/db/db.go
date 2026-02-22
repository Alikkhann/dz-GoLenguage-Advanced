package db

import (
	"gorm.io/gorm"
	"4-project/configs"
	"gorm.io/driver/postgres"
)

type Db struct {
	*gorm.DB
}

func NewDb(config *configs.Config) *Db{
	db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &Db{db}
}