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
		return
	} else if jsonData["data"] == nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	collectionName := jsonData["collection"].(string)
	err = os.MkdirAll("./data/"+collectionName, 0755)
	if err != nil {
		http.Error(w, "Failed to create collection directory", http.StatusInternalServerError)
		return
	}

	filePath := "./data/" + collectionName + "/data.json"

	_, err = os.Stat(filePath)
	if err != nil {
		file, err := os.Create(filePath)
		if err != nil {
			http.Error(w, "Failed to create data file", http.StatusInternalServerError)
			return
		}

		dataBytes, err := json.Marshal([]interface{}{jsonData["data"]})
		if err != nil {
			http.Error(w, "Failed to serialize data", http.StatusInternalServerError)
			return
		}

		_, err = file.Write(dataBytes)
		if err != nil {
			http.Error(w, "Failed to write data to file", http.StatusInternalServerError)
			return
		}

		defer file.Close()
	} else {
		file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			http.Error(w, "Failed to open data file", http.StatusInternalServerError)
			return
		}
		defer file.Close()
	
		existingData, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, "Failed to read data from file", http.StatusInternalServerError)
			return
		}
	
		var existingArray []interface{}
		err = json.Unmarshal(existingData, &existingArray)
		if err != nil {
			http.Error(w, "Failed to parse existing data", http.StatusInternalServerError)
			return
		}
	
		existingArray = append(existingArray, jsonData["data"])
	
		newDataBytes, err := json.Marshal(existingArray)
		if err != nil {
			http.Error(w, "Failed to serialize data", http.StatusInternalServerError)
			return
		}
	
		err = file.Truncate(0)
		if err != nil {
			http.Error(w, "Failed to truncate data file", http.StatusInternalServerError)
			return
		}
	
		_, err = file.Seek(0, 0)
		if err != nil {
			http.Error(w, "Failed to seek data file", http.StatusInternalServerError)
			return
		}
	
		_, err = file.Write(newDataBytes)
		if err != nil {
			http.Error(w, "Failed to write data to file", http.StatusInternalServerError)
			return
		}
	}
	
	
	

}
