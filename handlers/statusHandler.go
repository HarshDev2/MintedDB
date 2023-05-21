package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"github.com/harshdev2/db/utils"
)

type Response struct {
	Collections []string `json:"collections"`
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	// Define the data directory path
	dataDir, err := utils.GetCurrentPath();
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	// Check if the data directory exists
	if _, err := os.Stat(dataDir + "data"); os.IsNotExist(err) {
		http.Error(w, "Data directory does not exist", http.StatusInternalServerError)
		return
	}

	// Get all directories (no files) in the data directory
	directories, err := getDirectories(dataDir + "data")
	if err != nil {
		http.Error(w, "Failed to retrieve directories", http.StatusInternalServerError)
		return
	}

	// Check for the presence of data.json in each directory
	var collections []string
	for _, dir := range directories {
		dataJSONPath := filepath.Join(dataDir+"data", dir, "data.json")
		if _, err := os.Stat(dataJSONPath); err == nil {
			collections = append(collections, dir)
		}
	}

	// Prepare the response
	response := Response{
		Collections: collections,
	}

	// Convert the response to JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	// Set the response content type and write the JSON response
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonResponse)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}

func getDirectories(path string) ([]string, error) {
	var directories []string

	// Open the directory
	dir, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer dir.Close()

	// Read the directory entries
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return nil, err
	}

	// Iterate over the file infos and check for directories
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			directories = append(directories, fileInfo.Name())
		}
	}

	return directories, nil
}
