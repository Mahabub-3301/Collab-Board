package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)


type TaskHandler struct {
	service *TaskService
}

func NewTaskHandler(service *TaskService) *TaskHandler {
	return &TaskHandler {
		service : service,
	}
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var t Task 
	
	err:= json.NewDecoder(r.Body).Decode(&t)

	if err!=nil {
		writeError(w,http.StatusBadRequest, "invalid JSON")
		return 
	}

	if err := validate(t); err!=nil {
		writeError(w,http.StatusBadRequest,err.Error())
		return  
	}

	h.service.CreateTask(&t)

	writeJSON(w,http.StatusOK,t)
}

func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	data := h.service.GetTasks()

	writeJSON(w,http.StatusOK,data)
}

func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	ids := strings.TrimPrefix(r.URL.Path, "/task/")
	idr, err := strconv.Atoi(ids)
	if err!=nil {
		writeError(w,http.StatusBadRequest,"Invalid ID")
		return
	}
	task, err := h.service.GetTask(idr)
	if err != nil {
		writeError(w,http.StatusBadRequest,err.Error())
		return 
	}
	writeJSON(w,http.StatusOK,task)
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	ids := strings.TrimPrefix(r.URL.Path,"/task/")
	id, err := strconv.Atoi(ids)
	if err!=nil {
		writeError(w,http.StatusBadRequest,"Invalid ID")
		return
	}
	var t Task
	er := json.NewDecoder(r.Body).Decode(&t)
	if er!=nil {
		writeError(w,http.StatusBadRequest,"Invalid JSON")
		return
	}
	if err := validate(t); err!=nil {
		writeError(w,http.StatusBadRequest,err.Error())
		return
	}


	task, err := h.service.UpdateTask(id,&t)
	if err != nil {
		writeError(w,http.StatusBadRequest,err.Error())
		return
	}
	writeJSON(w,http.StatusOK,task)
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request){
	ids := strings.TrimPrefix(r.URL.Path,"/task/")
	id, err := strconv.Atoi(ids)

	if err!=err {
		writeError(w,http.StatusBadRequest,"Invalid ID")
		return 
	}

	ok := h.service.DeleteTask(id)
	if ok != nil {
		writeError(w,http.StatusNotFound,ok.Error())
		return 
	}

	w.WriteHeader(http.StatusNoContent)
}

