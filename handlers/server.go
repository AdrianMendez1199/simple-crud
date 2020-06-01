package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func StartServer() {
	router := mux.NewRouter()

	router.HandleFunc("/prueba", Prueba)
	http.Handle("/", router)

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "4000"
	}
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}

func Prueba(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
