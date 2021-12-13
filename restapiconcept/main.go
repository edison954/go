package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type task struct {
	Id      int    `json:Id`
	Name    string `json:Name`
	Content string `json:Content`
}

type allTasks []task

var tasks = allTasks{
	{
		Id:      1,
		Name:    "task one",
		Content: "some content",
	}, //ultima comma indica que continua la lista
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(tasks)
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my api Edison")
}

func main() {
	//fmt.Println("Hello Edison")
	router := mux.NewRouter().StrictSlash(true) //	StrictSlash para que no permita urls erroneas /allTasks

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/tasks", getTasks)

	log.Fatal(http.ListenAndServe(":3000", router))

}
