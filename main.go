package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	env, err := godotenv.Read(".env")
	if err != nil {
		log.Fatalf("ERROR : %v", err)
	}

	author := env["AUTHOR"]
	port := env["PORT"]
	address := fmt.Sprintf("0.0.0.0:%v", port)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		response := "Hallo "
		if author != "" {
			response = response + author
		}

		c.JSON(http.StatusOK, map[string]string{
			"data": response,
		})
	})

	log.Println("server run on : ", address)

	r.Run(address)
}
