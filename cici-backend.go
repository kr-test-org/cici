package main

import (
	"fmt"
	"net/http"
)

var BuildTime = "No Build Time Provided"

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello, my name is cici~\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func buildTime(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "build time :%s", BuildTime)
}

func main() {

	http.HandleFunc("/", buildTime)
	http.HandleFunc("/welcome", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil)
}
