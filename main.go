package main

import (
	"REST_API_GO/database"
	"REST_API_GO/utils"
	"log"
	"net/http"
)

func main() {
	utils.Connection()
	database.Migrate()
	log.Println("Listening on port 5000")
	http.ListenAndServe(":5000", nil)
}