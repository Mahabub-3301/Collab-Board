package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)


type TaskHandler struct {
	store *TaskStore
}

func NewTaskHandler(store *TaskStore) *TaskHandler {
	return &TaskHandler {
		store : store,
	}
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *htt.Request) {
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

	h.store.Create(&t)

	writeJSON(w,http.StatusOK,t)
}

func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	data := h.store.GetALL()

	writeJSON(w,http.StatusOK,data)
}

func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	ids := strings.TrimPrefix(r.URL.Path, "/task/")
	idr, err := strconv.Atoi(ids)
	if err!=nil {
		writeError(w,http.StatusBadRequest,"Invalid ID")
		return
	}
	task, ok := h.store.Get(idr)
	if !ok {
		writeError(w,http.StatusBadRequest,"Task not found")
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
	if !er {
		writeError(w,http.StatusBadRequest,"Invalid JSON")
		return
	}
	if err := validate(t); err!=nil {
		writeError(w,http.StatusBadRequest,err.Error())
		return
	}


	task, ok := h.store.Update(id,&t)
	if !ok {
		writeError(w,http.StatusBadRequest,"Task not Found")
		return
	}
	writeJSON(w,http.StatusOK,task)
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request){
	ids := strings.TrimPrefix(r.URL.Path,"/task/")
	id, err := strconv.Atoi(ids)

	if !err {
		writeError(w,http.StatusBadRequest,"Invalid ID")
		return 
	}

	ok := h.store.Delete(id)
	if !ok {
		writeError(w,http.StatusNotFound,"Task Not Found")
		return 
	}

	w.WriteHeader(http.StatusNoContent)
}

