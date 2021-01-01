package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func main() {
	newMux := http.NewServeMux()

	newMux.HandleFunc("/randomFlot", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, rand.Float32())
	})
	newMux.HandleFunc("/randomInt", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, rand.Int())
	})

	http.ListenAndServe(":8080", newMux)
}
