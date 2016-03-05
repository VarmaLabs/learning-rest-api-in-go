package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"io/ioutil"
)

type Fruit struct {
	Id 			int
	Name    string
	Varieties []string
}

var mango = Fruit{
    Id:   1,
    Name: "Mango",
    Varieties: []string{"Alphonso", "Florigon", "Van Dyke", "Ivory"}}

var apple = Fruit{
		Id:   2,
    Name:   "Apple",
    Varieties: []string{"Golden Supreme", "Jonathan", "Lady", "Liberty"}}

var fruits = map[string]Fruit{
	"1":  mango,
	"2": apple,
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/fruits", Index)
	router.HandleFunc("/fruits/{Id}", Show)
	router.HandleFunc("/fruits/", Create).Methods("POST")
	router.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8086", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in index")
	fruits_jm, _ := json.Marshal(fruits)
  w.Header().Set("Content-Type", "application/json")
  w.Write(fruits_jm)
}

func Show(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in show")
	vars := mux.Vars(r)
  Id := vars["Id"]
	fruit_jm, _ := json.Marshal(fruits[Id])

  w.Header().Set("Content-Type", "application/json")
  w.Write(fruit_jm)
}

func Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in create")
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println(body)

  w.Header().Set("Content-Type", "application/json")
  data_jm, _ := json.Marshal("hello")
  w.Write(data_jm)
}