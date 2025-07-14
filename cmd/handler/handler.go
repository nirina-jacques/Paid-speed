package handler

import (
	"encoding/json"
	"net/http"
)

func (cf *ServerConfig) HomePage(w http.ResponseWriter, r *http.Request) {
	message := map[string]string{"text": "Bienvenue sur la page d'accueil!"}
	jsonResponse, _ := json.Marshal(message)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
