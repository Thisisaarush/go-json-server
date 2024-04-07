package main

import "net/http"

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, r, 200, struct{}{})
}