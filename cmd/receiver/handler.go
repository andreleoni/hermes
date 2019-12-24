package main

import (
	"fmt"
	"hermes/internal/storage"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	storage.SetupQueue()

	router := mux.NewRouter()

	router.HandleFunc("/{queue}", ReceiverHandler).Methods("POST")

	fmt.Println("Starting server...")

	http.Handle("/", router)

	http.ListenAndServe(":8090", nil)
}

func ReceiverHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	w.WriteHeader(http.StatusOK)

	fmt.Println("Params: ", vars)

	fmt.Println("Reading test: ", storage.QueueRead(vars["queue"]))
	fmt.Println("Write test: ", storage.QueueWrite(vars["queue"], "teste"))

	fmt.Fprintf(w, "Queue", vars["queue"])
}
