package handlers

import (
	"encoding/json"
	"github.com/harshdev2/db/utils"
	"io"
	"net/http"
	"os"
)

func FindHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a variable of type map[string]interface{} to store the parsed JSON data
	var jsonData map[string]interface{}

	// Parse the JSON data into the 'jsonData' variable
	err = json.Unmarshal(body, &jsonData)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if jsonData["collection"] == nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
	} else if jsonData["queryData"] == nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
	} else {
		filePath := "data/" + jsonData["collection"].(string) + "/data.json"

		file, err := os.Open(filePath)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}

		// Read the file contents
		fileContent, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, "Failed to parse json", http.StatusInternalServerError)
		}

		defer file.Close()

		if err == nil {
			var jsonArray []interface{}
			err = json.Unmarshal(fileContent, &jsonArray)
			if err != nil {
				http.Error(w, "Failed to parse json", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			data, err := utils.MatchObjects(jsonData["queryData"], jsonArray, false)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			} else {
				if data == nil {
					data = []interface{}{}
				}
			
				responseData, err := json.Marshal(data)
				if err != nil {
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
					return
				}
			
				// Write the response body
				w.Write(responseData)
			}
		}
	}
}
