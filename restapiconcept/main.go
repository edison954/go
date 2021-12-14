package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

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

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(tasks)
}

func getTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	taskId, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid Id")
		return
	}

	for _, task := range tasks {
		if task.Id == taskId {
			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(task)
		}
	}

}

func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask task

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Inserte datos validos para la tarea")
	}

	json.Unmarshal(reqBody, &newTask)

	newTask.Id = len(tasks) + 1

	tasks = append(tasks, newTask)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)

}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Id de tarea invalido")
	}

	for index, task := range tasks {
		if task.Id == taskId {
			tasks = append(tasks[:index], tasks[index+1:]...)
			fmt.Fprintf(w, "La tarea con el id %v ha sido removida satisfactoriamente ", taskId)
		}
	}

}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my api Edison")
}

func main() {
	//fmt.Println("Hello Edison")
	router := mux.NewRouter().StrictSlash(true) //	StrictSlash para que no permita urls erroneas /allTasks

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks", createTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", getTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", router))

}
