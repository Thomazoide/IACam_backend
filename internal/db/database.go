package db

import (
	"fmt"
	"log"

	"github.com/Thomazoide/IACam_backend/internal/config"
	"github.com/Thomazoide/IACam_backend/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(cfg config.Config) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)
	var err error = nil
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error conectando a BBDD: \n", err.Error())
		return err
	}
	DB.AutoMigrate(&models.Camera{})
	log.Println("Conectado a BBDD")
	return nil
}

func GetInstance() *gorm.DB {
	return DB
}
