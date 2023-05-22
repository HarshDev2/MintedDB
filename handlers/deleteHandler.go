package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/harshdev2/db/utils"
)

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
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

	currentPath, err := utils.GetCurrentPath();
	
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
	}
	
	filePath := currentPath + "data/" + collectionName + "/data.json"

	_, error := os.Stat(filePath)
	if error != nil {
		http.Error(w, "Collection not found", http.StatusInternalServerError)
		return
	}

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		http.Error(w, "Failed to open data file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	fileData, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Failed to read data from file", http.StatusInternalServerError)
		return
	}

	var data []interface{}
	err = json.Unmarshal(fileData, &data)
	if err != nil {
		http.Error(w, "Failed to parse existing data", http.StatusInternalServerError)
		return
	}

	// Check if properties and values exist in any object of data
	var solutions []map[string]interface{}
	var filterData []map[string]interface{}
	for _, obj := range data {
		objMap, ok := obj.(map[string]interface{})
		if !ok {
			continue
		}

		match := true
		for key, value := range jsonData["data"].(map[string]interface{}) {
			if objValue, ok := objMap[key]; !ok || objValue != value {
				match = false
				break
			}
		}

		if match {
			solutions = append(solutions, objMap)
		} else {
			filterData = append(filterData, objMap)
		}
	}

	// Handle the results
	if len(solutions) == 1 {
		filteredDataJSON, _ := json.Marshal(filterData)
		err :=  os.WriteFile(filePath, filteredDataJSON, 0644)
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK);
		}
	} else if len(solutions) > 1 {
		http.Error(w, "Object not found", http.StatusInternalServerError)
	} else {
		http.Error(w, "Object not found", http.StatusInternalServerError)
	}
}
