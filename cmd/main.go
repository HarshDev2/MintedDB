package main

import (
	"fmt"
	"github.com/harshdev2/db/handlers"
	"net/http"
	"os"
)

func main() {
	_, err := os.Stat("data")

	if err == nil {

	} else {

		err := os.Mkdir("data", 0755)
		if err != nil {
			fmt.Println("Something went wrong!")
		} else {

			http.HandleFunc("/status", handlers.StatusHandler)
			http.HandleFunc("/api/create", handlers.CreateHandler)

			// Start the server
			fmt.Println("Server started on port 8080")
			http.ListenAndServe(":8080", nil)
		}
	}
}
