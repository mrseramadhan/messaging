package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "./models"
	routes "./routes"
	"github.com/gorilla/mux"
)

func startServer() {
	port := os.Getenv("web_port")
	prefix := os.Getenv("prefix")
	fmt.Println("Server started at " + port + "...")
	r := mux.NewRouter().StrictSlash(true)
	// Routes
	routes.ApiRoutes(prefix, r)

	//Start Server on the port set in your .env file
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	startServer()
}
