package main

import (
	"fmt"
	"net/http"

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

	http.HandleFunc("/generate-plan", handler.GeneratePlan)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("failed to start the server ", err)
	}
}
