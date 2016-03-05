package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", RenderJson)
	log.Fatal(http.ListenAndServe(":8084", router))
}

func RenderJson(w http.ResponseWriter, r *http.Request) {
	var fruits = map[string]map[string]string{
    "1": map[string]string{"Id": "1", "Name": "Mango", "Varieties": "Alphonso, Florigon, Van Dyke, Ivory"},
    "2": map[string]string{"Id": "2", "Name": "Apple", "Varieties": "Golden Supreme, Jonathan, Lady, Liberty"},
	}
	fruits_jm, _ := json.Marshal(fruits)
  w.Header().Set("Content-Type", "application/json")
  w.Write(fruits_jm)
}