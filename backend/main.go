package main

import (
	"fmt"
	"net/http"


)


func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w,"Server is Healthy")
}
func setUpRoutes(taskHandler *TaskHandler) http.Handler {

		mux := http.NewServeMux()
		
		mux.HandleFunc("/health", healthHandler)


		mux.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case http.MethodGet :
				taskHandler.GetTasks(w,r)

			case http.MethodPost :
				taskHandler.CreateTask(w,r)

			default:
				w.Header().Set("Allow","GET, POST")
				writeError(w,http.StatusMethodNotAllowed,"Method Not Allowed")
			}
		})
	
		mux.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case http.MethodGet :
				taskHandler.GetTask(w,r)

			case http.MethodPut :
				taskHandler.UpdateTask(w,r)

			case http.MethodDelete:
				taskHandler.DeleteTask(w,r)

			default:
				w.Header().Set("Allow","GET, PUT, DELETE")
				writeError(w,http.StatusMethodNotAllowed,"Method not Allowed")
			}
		})

		return mux
	}

func main() {
	

	store 		:= NewTaskStore()
	service 	:= NewTaskService(store)
	taskHandler := NewTaskHandler(service)
	Router 		:= setUpRoutes(taskHandler)

	finalHandler := Chain(
		Router,
		RecoveryMiddleware,
		RequestIDMiddleware,
		LoggerMiddleware,
	)

	port := 9090
	add := fmt.Sprintf(":%d",port)


	fmt.Printf("server is running on http://localhost:%d\n",port)

	err := http.ListenAndServe(add,finalHandler)
	if err!=nil {
		panic(err)
	}
}
