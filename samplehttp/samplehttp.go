package main

import (
    "fmt"
    "os"
    "net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
    fmt.Fprintf(w, "hello from %s",name)
}

func headers(w http.ResponseWriter, req *http.Request) {

    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func main() {

    http.HandleFunc("/hello", hello)
    http.HandleFunc("/browser", headers)

    http.ListenAndServe(":80", nil)
}
