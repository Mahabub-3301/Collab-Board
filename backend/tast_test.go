
package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"sync"
)

func TestTaskStoreConcurrentCreate(t *testing.T){
	store := NewTaskStore()
	const numTasks = 100

	var wg sync.WaitGroup

	wg.Add(numTasks)


	for i:=0;i<numTasks;i++ {
		go func() {
			defer wg.Done()

			task := &Task{
				Title: "something...",
				Status: "todo",
			}

			store.Create(task)
		}()
	}

	wg.Wait()

	tasks := store.GetALL()

	if len(tasks) != numTasks {
		t.Errorf(
			"expected %d tasks, got %d",
			numTasks,
			len(tasks),
		)
	}
}
				

func TestGetTasks(t *testing.T) {

	store := NewTaskStore()
	handler := NewTaskHandler(store)

	request := httptest.NewRequest(
		http.MethodGet,
		"/tasks",
		nil,
	)

	response := httptest.NewRecorder()
	handler.GetTasks(response, request)

	if response.Code != http.StatusOK {
		t.Errorf(
			"expected status %d, got %d",
			http.StatusOK,
			response.Code,
		)
	}
}




func TestValidate(t *testing.T) {

	tests := []struct {
		name	string
		task	Task 
		wantErr	bool
	}{
		{
			name: "valid Task",
			task: Task{
				Title: "Learn Go",
				Status: "todo",
			},
			wantErr: false,
		},
		{
			name: "empty title",
			task: Task{
				Title: "",
				Status: "todo",
			},
			wantErr: true,
		},
		{
			name: "Invalid status",
			task: Task{
				Title: "learn go",
				Status: "babna",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func (t *testing.T) {
			err:= validate(tt.task)
			if (err !=nil) != tt.wantErr {
				t.Errorf(
					"validate() error = %v, wantErr = %v",
					err,
					tt.wantErr,
				)
			}
		})
	}
}


