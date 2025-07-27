package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, Galua! Auto commit!!!")
}

func main() {
    http.HandleFunc("/", handler)

    fmt.Println("Starting server...")
    err := http.ListenAndServe(":80", nil)
    if err != nil {
        fmt.Println(err)
    }
}