package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func articleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "Category is: %v\n", vars["category"])
	fmt.Fprintf(w, "id is: %v\n", vars["id"])
}

func articleHandlerQuery(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "Category is: %v\n", queryParams["id"][0])
	fmt.Fprintf(w, "id is: %v\n", queryParams["category"][0])
}

func main() {
	router := mux.NewRouter()

	/* 😱 pattern matching as route */
	router.HandleFunc("/articles/{category}/{id:[0-9]+}", articleHandler)
	/* route with query */
	router.HandleFunc("/articleswithquery", articleHandlerQuery)

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
