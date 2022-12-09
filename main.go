package main

import (
	"go-docker/db"
	"go-docker/route"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	dbConn, err := db.Connection(".env.prod")
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	r := gin.Default()

	route.Router(dbConn, r)

	if err := r.Run("0.0.0.0:8080"); err != nil {
		log.Fatal(err)
	}
}
