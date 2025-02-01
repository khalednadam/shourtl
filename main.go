package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ShorterRequest struct {
	Name string
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world"))
	})

	/*
	   TODO: Complete
	   - Generate a unique id for the URL
	   - Store it in DB
	*/
	router.HandleFunc("POST /shorten", func(w http.ResponseWriter, r *http.Request) {
		requestData := ShorterRequest{}
		err := json.NewDecoder(r.Body).Decode(&requestData)
		if err != nil {
			fmt.Print("Error: ")
			panic(err)
		}
		w.Write([]byte(requestData.Name))
	})

	server := http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	fmt.Println("Listenting on port", server.Addr)
	server.ListenAndServe()
}
