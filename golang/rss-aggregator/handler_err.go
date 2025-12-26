package main

import (
	"log"
	"net/http"
)

func handlerErr(w http.ResponseWriter, r *http.Request) {
	log.Println("error route encounterd")
	respondWithError(w, 400, "Something went wrong")
}
