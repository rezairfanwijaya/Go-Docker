package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	env, err := godotenv.Read(".env")
	if err != nil {
		log.Fatalf("ERROR : %v", err)
	}

	author := env["AUTHOR"]
	port := env["PORT"]
	address := fmt.Sprintf("localhost:%v", port)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		response := "Hello"
		if author != "" {
			response = response + " from " + author
		}

		w.Write([]byte(response))
	})

	if port == "" {
		log.Fatal("PORT env is required")
	}

	fmt.Printf("server run on %v\n", address)

	if err := http.ListenAndServe(address, mux); err != nil {
		log.Fatal(err)
	}
}
