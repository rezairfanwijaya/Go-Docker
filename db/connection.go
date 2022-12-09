package db

import (
	"errors"
	"fmt"
	"go-docker/helper"
	"go-docker/user"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection(path string) (*gorm.DB, error) {
	// get env
	env, err := helper.GetEnv(path)
	if err != nil {
		return nil, err
	}

	databaseUsername := env["DATABASE_USERNAME"]
	databasePassword := env["DATABASE_PASSWORD"]
	databaseHost := env["DATABASE_HOST"]
	databasePort := env["DATABASE_PORT"]
	databaseName := env["DATABASE_NAME"]

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", databaseUsername, databasePassword, databaseHost, databasePort, databaseName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		errMsg := fmt.Sprintf("dsn : %v", err)
		return db, errors.New(errMsg)
	}

	db.AutoMigrate(&user.User{})
	return db, nil
}
