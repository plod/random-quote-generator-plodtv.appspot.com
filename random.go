package random

import (
    "fmt"
    "net/http"
)

func init() {
    http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
    data, _ := Asset("random.html")
    fmt.Fprint(w, string(data))
}
//go-bindata -prefix "./static/" -pkg random ./static/random.html