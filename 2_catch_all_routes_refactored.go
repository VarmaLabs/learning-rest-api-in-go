package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
)

func main() {
  http.HandleFunc("/", HandlerFunc)
  log.Fatal(http.ListenAndServe(":8082", nil))
}

func HandlerFunc(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello Techies\n")
  fmt.Fprintf(w, "URL: %q", html.EscapeString(r.URL.Path))
}