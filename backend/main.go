package main

import (
	"fmt"
	"net/http"


)


func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w,"Server is Healthy")
}


func main() {

	store := NewTaskStore()
	service := NewTaskService(store)
	TaskHandler := NewTaskHandler(service)


	http.HandleFunc("/health", healthHandler)


	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet :
			TaskHandler.GetTasks(w,r)

		case http.MethodPost :
			TaskHandler.CreateTask(w,r)

		default:
			w.Header().Set("Allow","GET, POST")
			writeError(w,http.StatusMethodNotAllowed,"Method Not Allowed")
		}
	})

	http.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet :
			TaskHandler.GetTask(w,r)

		case http.MethodPut :
			TaskHandler.UpdateTask(w,r)

		case http.MethodDelete:
			TaskHandler.DeleteTask(w,r)

		default:
			w.Header().Set("Allow","GET, PUT, DELETE")
			writeError(w,http.StatusMethodNotAllowed,"Method not Allowed")
		}
	})


	port := 9090
	add := fmt.Sprintf(":%d",port)


	fmt.Printf("server is running on http://localhost:%d\n",port)

	err := http.ListenAndServe(add,Logger{})
	if err!=nil {
		panic(err)
	}
}
