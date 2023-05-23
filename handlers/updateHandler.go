package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"

	"github.com/harshdev2/db/utils"
)

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
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
	} else if jsonData["dataToBeUpdated"] == nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	} else if jsonData["newData"] == nil {
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

	for i := 0; i < len(data); i++ {
		if obj, ok := data[i].(map[string]interface{}); ok {
			if newData, ok := jsonData["newData"].(map[string]interface{}); ok {
				if obj["_id"] == newData["_id"] {
					// Compare properties in jsonData["newData"] with data[i]
					for key, value := range newData {
						if key != "_id" {
							if oldValue, exists := obj[key]; exists {
								if !reflect.DeepEqual(oldValue, value) {
									// Update property value in data[i] with the value from jsonData["newData"]
									obj[key] = value
								}
							} else {
								// Add missing property from jsonData["newData"] to data[i]
								obj[key] = value
							}
						}
					}
				}
			}
		}
	}
	
	

	fmt.Println(data);

}
