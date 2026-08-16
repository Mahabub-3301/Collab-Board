package main

import "sync"

type TaskStore struct {
	mu sync.RWMutex
	tasks 	map[int]*Task 
	nextID 	int
}

func NewTaskStore() *TaskStore {
	return &TaskStore{
		tasks: make(map[int]*Task),
		nextID: 1,
	}
}

func (s *TaskStore) Create(t *Task) *Task {
	s.mu.Lock()
	defer s.mu.Unlock()

	t.ID = s.nextID
	s.tasks[s.nextID] = t 
	s.nextID++
	return t 
}

func (s *TaskStore) GetALL() map[int]*Task {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.tasks
}

func (s *TaskStore) Get(id int) (*Task,bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	task, ok := s.tasks[id]

	return task, ok
}

func (s *TaskStore) Update(id int, t *Task) (*Task, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _,ok := s.tasks[id]; !ok {
		return nil, false
	}

	t.ID = id 
	s.tasks[id] = t 

	return t, true
}

func (s *TaskStore) Delete(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _,ok := s.tasks[id]; !ok {
		return false
	}

	delete(s.tasks, id)
	return true
}




