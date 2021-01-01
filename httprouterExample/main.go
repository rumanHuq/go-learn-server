package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"

	"github.com/julienschmidt/httprouter"
)

func goVersion(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	out, _ := exec.Command("/usr/local/go/bin/go", "version").Output()
	io.WriteString(w, string(out))
}

func readFile(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	out, _ := exec.Command("/bin/cat", params.ByName("name")+".txt").Output()
	fmt.Fprintf(w, string(out))
}

func main() {
	router := httprouter.New()

	router.GET("/api/v1/go-version", goVersion)
	router.GET("/api/v1/show-file/:name", readFile)
	log.Fatal(http.ListenAndServe(":8080", router))
}
