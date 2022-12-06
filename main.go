package main

import (
	"go-docker/db"
	"go-docker/route"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	dbConnection, err := db.Connection(".env")
	if err != nil {
		log.Fatalf("ERROR CONNECTION : %v", err)
	}

	r := gin.Default()

	route.Router(dbConnection, r)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
