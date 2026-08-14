package main

type TaskStore struct {
	tasks 	map[int]*Task 
	nextID 	int
}

func NewTaskStore() *TaskStore {
	return &TaskStore{
		tasks: make(map[int]*Task)
		nextID: 1
	}
}

func (s *TaskStore) Create(t *Task) *Task {
	t.ID = s.nextID
	s.tasks[s.nextID] = t 
	s.nextID++
	return t 
}

func (s *TaskStore) GetALL() map[int]*Task {
	return s.tasks
}

func (s *TaskStore) Get(id int) (*Task,bool) {
	task, ok = s.tasks[id]

	return task, ok
}

func (s *TaskStore) Update(id int, t *Task) (*Task, bool) {
	if _,ok := s.tasks[id]; ok!=nil {
		return nil, false
	}

	t.ID = id 
	s.tasks[id] = t 

	return t, true
}

func (s *TaskStore) Delete(id int) bool {
	if _,ok := s.tasks[id]; !ok {
		return false
	}

	delete(s.tasks, id)
	return true
}




