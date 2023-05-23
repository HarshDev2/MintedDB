package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/harshdev2/db/handlers"
	"github.com/harshdev2/db/panel/handlers"
	"github.com/harshdev2/db/utils"
)

func main() {
	currentPath, err := utils.GetCurrentPath()

	if err != nil {
		fmt.Println("Something went wrong, try running the exe again.")
		return
	}

	_, err = os.Stat(currentPath + "data")
	if err != nil {
		err := os.Mkdir(currentPath+"data", 0755)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	_, err = os.Stat(currentPath + "data/users.json")
	if err != nil {
		userFile, err := os.Create(currentPath + "data/users.json")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer userFile.Close()
	}

	http.HandleFunc("/api/status", handlers.StatusHandler)
	http.HandleFunc("/api/create", handlers.CreateHandler)
	http.HandleFunc("/api/find", handlers.FindHandler)
	http.HandleFunc("/api/find-many", handlers.FindManyHandler)
	http.HandleFunc("/api/delete", handlers.DeleteHandler)
	http.HandleFunc("/api/update", handlers.UpdateHandler)
	http.HandleFunc("/", panelhandlers.PanelHandler)

	// Start the server
	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}
