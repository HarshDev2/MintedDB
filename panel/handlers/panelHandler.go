package panelhandlers

import (
	"net/http"
	"os"
	"github.com/harshdev2/db/utils"
)

func PanelHandler(w http.ResponseWriter, r *http.Request) {
	path, err := utils.GetCurrentPath();
	if(err != nil){
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
	}
	file, err := os.Open(path + "panel/html/index.html")
	if err != nil {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}

	http.ServeContent(w, r, file.Name(), fileStat.ModTime(), file)
}
