package main

import (
	"encoding/json"
	"log"
	"net/http"
	"users-api/configs"
	"users-api/routes"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	configs.ConnectDB()
	routes.UserRoute(router) //add this

	router.HandleFunc("/status", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(map[string]string{"statusCode": "200", "success": "true"})
	}).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
