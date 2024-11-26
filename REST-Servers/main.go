package main

import (
	"go-pj-based-learning/REST-Servers/helper"
	"go-pj-based-learning/REST-Servers/internal/taskstore"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type taskServer struct {
	store *taskstore.TaskStore
}

// taskServerにstore(=DB)情報をラップする
func NewTaskServer() *taskServer {
	store := taskstore.New()
	return &taskServer{store: store}
}

// get task handler
func (ts *taskServer) getTaskHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling get tasks at %s\n", req.URL.Path)

	// id, err := strconv.Atoi(req.PathValue("id"))

	id, _ := strconv.Atoi(mux.Vars(req)["id"])

	task, err := ts.store.GetTask(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	helper.RenderJson(w, task)
}

// get all tasks
func (ts *taskServer) getAllTasksHandler(w http.ResponseWriter, req *http.Request) {

}

// routing adding
func main() {
	// standard library
	// mux := http.NewServeMux()
	// server := NewTaskServer()

	// mux.HandleFunc("Get /task/{id}/", server.getTaskHandler)

	// using gorilla/mux
	router := mux.NewRouter()
	router.StrictSlash(true)
	server := NewTaskServer()

	router.HandleFunc("/task/{id:[0-9]+}/", server.getTaskHandler).Methods("GET")
	router.HandleFunc("/task/{id:[0-9]+}/", server.getAllTasksHandler).Methods("GET")

	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
