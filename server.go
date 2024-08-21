package main

import (
	"log"
	"net/http"
)

func startServer() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /students", createStudentsHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func createStudentsHandler(w http.ResponseWriter, r *http.Request) {

	error := CreateStudents()
	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusCreated)
}
