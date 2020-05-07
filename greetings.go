package main

import (
    "fmt"
    "log"
    "net/http"
)

func greetings(message string) string {
    return "<b>" + message + "</b>"
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, greetings("Code.education Rocks!"))
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8000", nil))
}