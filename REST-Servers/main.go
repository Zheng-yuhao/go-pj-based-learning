package main

import (
	"go-pj-based-learning/REST-Servers/helper"
	"go-pj-based-learning/REST-Servers/internal/taskstore"
	"log"
	"net/http"
	"strconv"
)

type taskServer struct {
	store *taskstore.TaskStore
}

func NewTaskServer() *taskServer {
	store := taskstore.New()
	return &taskServer{store: store}
}

// get task handler
func (ts *taskServer) getTaskHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling get tasks at %s\n", req.URL.Path)

	id, err := strconv.Atoi(req.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	task, err := ts.store.GetTask(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	helper.RenderJson(w, task)
}

// routing adding
func main() {
	mux := http.NewServeMux()
	server := NewTaskServer()

	mux.HandleFunc("Get /task/{id}/", server.getTaskHandler)
}
