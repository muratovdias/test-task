package handler

import (
	"encoding/json"
	"net/http"
	"salt/internal/service"
)

func InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/generate-salt", generateSalt)
	return mux
}

func generateSalt(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	salt := service.GetSalt()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{"salt": salt}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
