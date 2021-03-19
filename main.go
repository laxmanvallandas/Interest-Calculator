package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/laxmanvallandas/assignment/pkg/handler"
)

var (
	//build commit id
	Build string
	//Name service name
	Name string
	//Version service version
	Version string
)

func main() {
	fmt.Println("Service ", Name, "started. Version", Version, " Build Date: ", Build)

	r := mux.NewRouter()

	r.HandleFunc("/generate-plan", handler.GeneratePlan).
		Methods("POST").
		Schemes("http")

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
