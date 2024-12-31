package main

import (
	"log"
	"net/http"

	"github.com/Tejsvapandey1/goCRUD.git/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	http.Handle("/",r)
	log.Fatal(http.ListenAndServe("localhost:8000",r))
}