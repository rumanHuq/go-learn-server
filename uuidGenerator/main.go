package main

import (
	"crypto/rand"
	"fmt"
	"net/http"
)

type uuid struct{}

func getRandomID(w http.ResponseWriter, r *http.Request) {
	c := 10
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, fmt.Sprintf("%x", b))
}

func (u *uuid) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		getRandomID(w, r)
		return
	}
	http.NotFound(w, r)
}

func main() {
	mux := &uuid{}
	http.ListenAndServe(":8080", mux)
}
