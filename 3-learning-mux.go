package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/cars", Cars)
	router.HandleFunc("/mobiles", Mobiles)
	router.HandleFunc("/bags", Bags)
	router.HandleFunc("/", Home)
	log.Fatal(http.ListenAndServe(":8083", router))
}

func Cars(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You have requested for cars\n")
  fmt.Fprintf(w, "URL: %q", html.EscapeString(r.URL.Path))
}

func Mobiles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You have requested for mobiles\n")
  fmt.Fprintf(w, "URL: %q", html.EscapeString(r.URL.Path))
}

func Bags(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You have requested for bags\n")
  fmt.Fprintf(w, "URL: %q", html.EscapeString(r.URL.Path))
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome\n")
  fmt.Fprintf(w, "URL: %q", html.EscapeString(r.URL.Path))
}