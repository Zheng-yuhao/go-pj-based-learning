package helper

import (
	"encoding/json"
	"net/http"
)

func RenderJson(w http.ResponseWriter, result interface{}) {
	resp, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}
