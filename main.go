package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"khalednadam/shourtl/cmd/api"
	"khalednadam/shourtl/config"
	"khalednadam/shourtl/db"
	"khalednadam/shourtl/service"
	"log"
	"net/http"

	"github.com/go-sql-driver/mysql"
)

type ShorterRequest struct {
	Name string
}

func main() {
	router := api.NewAPIServer(":3000")

	db, err := db.NewMySQLStorage(mysql.Config{
		User:      config.Envs.DBUser,
		Passwd:    config.Envs.DBPassword,
		Addr:      config.Envs.DBAddress,
		DBName:    config.Envs.DBName,
		Net:       "tcp",
		ParseTime: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	initStorage(db)

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

	service.TestRoutes(router)
	err = router.Run()
	if err != nil {
		log.Fatal(err)
	}
}

// Starts the DB connection
func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB: Successfully connected")
}
