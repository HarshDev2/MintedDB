package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
	} else {
		body, err := io.ReadAll(r.Body)

		if err == nil {
			fmt.Println(body)
		}
	}

}
