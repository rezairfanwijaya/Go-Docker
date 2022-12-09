package main

import (
	"fmt"
	"go-docker/route"
	"log"
	"os/user"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	env, err := godotenv.Read(".env.prod")
	if err != nil {
		errMsg := fmt.Sprintf("ERROR GET ENV : %v", err)
		log.Fatal(errMsg)
		return
	}

	databaseUsername := env["DATABASE_USERNAME"]
	databasePassword := env["DATABASE_PASSWORD"]
	databaseHost := env["DATABASE_HOST"]
	databasePort := env["DATABASE_PORT"]
	databaseName := env["DATABASE_NAME"]

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", databaseUsername, databasePassword, databaseHost, databasePort, databaseName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		errMsg := fmt.Sprintf("ERROR CONNECTION : %v", err)
		log.Fatal(errMsg)
		return
	}

	db.AutoMigrate(&user.User{})

	r := gin.Default()

	route.Router(db, r)

	if err := r.Run("0.0.0.0:8080"); err != nil {
		log.Fatal(err)
	}
}
