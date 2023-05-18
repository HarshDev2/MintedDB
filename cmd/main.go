package main

import (
	"fmt"
	"github.com/harshdev2/db/handlers"
	"net/http"
	"os"
)

func main() {
	_, err := os.Stat("./data")

	if err != nil {
		err := os.Mkdir("./data", 0755)
		if err != nil {
			fmt.Println(err)
			return;
		}
	}

	http.HandleFunc("/status", handlers.StatusHandler)
	http.HandleFunc("/api/create", handlers.CreateHandler)
	http.HandleFunc("/api/find", handlers.FindHandler)
	
	// Start the server
	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)

}
