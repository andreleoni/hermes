package main

import (
	"bytes"
	"fmt"
	"hermes/internal/queue"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	queue.Setup()

	router := mux.NewRouter()

	router.HandleFunc("/{queue}", ReceiverHandler).Methods("POST")

	fmt.Println("Starting server...")

	http.Handle("/", router)

	http.ListenAndServe(":8090", nil)
}

func ReceiverHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	w.WriteHeader(http.StatusOK)

	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	stringBody := bytes.NewBuffer(body).String()

	fmt.Println("Request params: ", vars, "Body: ", stringBody)

	// fmt.Println("Reading test: ", queue.Read(vars["queue"]))

	fmt.Println("Writing on queue test: ", queue.Enqueue(vars["queue"], stringBody))

	fmt.Fprintf(w, "OK")
}
