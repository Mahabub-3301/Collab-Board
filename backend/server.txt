package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"strings"
	"strconv"
	"errors"
)

type Logger struct{}

type APIError struct{
	Error string `json:"error"`
}

type Task struct{
	ID	int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Status	string `json:"status"`
}

var Tasks = make(map[int]*Task)
var nextID = 1

func validate(t Task) error {
	if strings.TrimSpace(t.Title) == "" {
		return errors.New("Title is required")
	}
	switch t.Status {
	case "todo", "in-progress", "done":

	default:
		return errors.New("Invalid Status. Must be todo, in-progress or done")
	}
	return nil
}

func writeJSON(w  http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(status)
	
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w,status,APIError{Error:message})
}



func createTask(w http.ResponseWriter, r *http.Request) {
	var t Task
	json.NewDecoder(r.Body).Decode(&t)
	err := validate(t)
	if err != nil {
		writeError(w,http.StatusBadRequest,err.Error())
		return
	}
	t.ID = nextID
	Tasks[nextID] = &t
	nextID++
	writeJSON(w, http.StatusCreated, t)
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	writeJSON(w,http.StatusOK,Tasks)
}

func getTask(w http.ResponseWriter, r *http.Request) {
	ids := strings.TrimPrefix(r.URL.Path, "/Tasks/")
	idr, err := strconv.Atoi(ids)
	if err!= nil {
		writeError(w,http.StatusBadRequest,"Invalid ID")
		return
	}

	Task, ok := Tasks[idr]
	if ok {
		writeJSON(w,http.StatusOK,Task)
		return
	}
	writeError(w,http.StatusNotFound,"Task Not Found")
}

func updateTasks(w http.ResponseWriter, r *http.Request) {
	ids := strings.TrimPrefix(r.URL.Path, "/Tasks/")
	idr, err := strconv.Atoi(ids)
	if err!= nil {
		writeError(w,http.StatusBadRequest,"Invalid ID")
		return
	}
	var t Task
	er := json.NewDecoder(r.Body).Decode(&t)
	if er!=nil {
		writeError(w,http.StatusBadRequest,"Invalid Json")
		return
	}
	if err:= validate(t); err != nil {
		writeError(w,http.StatusBadRequest,err.Error())
		return
	}
	t.ID = idr
	Tasks[idr] = &t
	writeJSON(w,http.StatusOK,t)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	ids := strings.TrimPrefix(r.URL.Path, "/Tasks/")
	idr, err := strconv.Atoi(ids)
	if err!= nil {
		writeError(w,http.StatusBadRequest,"Invalid ID")
		return
	}

	if _, ok := Tasks[idr]; ok {
		delete(Tasks,idr)
		w.WriteHeader(http.StatusNoContent)
		return
	}
	writeError(w,http.StatusNotFound,"Task not Found")
}

func (l Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path, r.Method)
	http.DefaultServeMux.ServeHTTP(w,r)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w,"Server is Healthy")
}


func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet :
			getTasks(w,r)
		case http.MethodPost :
			createTask(w,r)
		default:
			w.Header().Set("Allow","GET, POST")
			writeError(w,http.StatusMethodNotAllowed,"Method Not Allowed")
		}
	})

	http.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet :
			getTask(w,r)
		case http.MethodPut :
			updateTasks(w,r)
		case http.MethodDelete:
			deleteTask(w,r)
		default:
			w.Header().Set("Allow","GET, PUT, DELETE")
			writeError(w,http.StatusMethodNotAllowed,"Method not Allowed")
		}
	})


	err := http.ListenAndServe(":8080",Logger{})

	if err!=nil {
		panic(err)
	}
}
