package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
		return
	}

	// Read the request body
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
	} else if jsonData["data"] == nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
	} else {
		collectionName := jsonData["collection"].(string)
		err := os.MkdirAll("./data/"+collectionName, 0755)
		if err != nil {
			http.Error(w, "Failed to create collection directory", http.StatusInternalServerError)
			return
		}
	
		filePath := "./data/" + collectionName + "/data.json"
		file, err := os.Create(filePath)
		if err != nil {
			http.Error(w, "Failed to create data file", http.StatusInternalServerError)
			return
		}
		defer file.Close()
	
		dataBytes, err := json.Marshal(jsonData["data"])
		if err != nil {
			http.Error(w, "Failed to serialize data", http.StatusInternalServerError)
			return
		}
	
		_, err = file.Write(dataBytes)
		if err != nil {
			http.Error(w, "Failed to write data to file", http.StatusInternalServerError)
			return
		}
	}
	
}
