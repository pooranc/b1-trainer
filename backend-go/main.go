package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pooranc/b1-trainer/backend-go/db"
	"github.com/pooranc/b1-trainer/backend-go/handlers"
	"github.com/pooranc/b1-trainer/backend-go/middleware"
)

func main() {
	db.Connect()

	r := mux.NewRouter()

	handlers.RegisterCardRoutes(r)
	handlers.RegisterSessionRoutes(r)

	log.Println("Server starting on port 8081...")
	log.Fatal(http.ListenAndServe(":8081", middleware.CORS(r)))
}
