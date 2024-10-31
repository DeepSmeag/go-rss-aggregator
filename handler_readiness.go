package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct {
		Ready   bool
		Message string
	}{Ready: true, Message: "Yes I'm ready"})
}
