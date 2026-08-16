package main

import "errors"


type TaskService struct {
        repository TaskRepository
}

func NewTaskService(repository TaskRepository) *TaskService {
        return &TaskService{
                repository: repository,
        }
}

func (s *TaskService) CreateTask(task *Task) (*Task, error) {
        if err := validate(*task); err != nil {
                return nil, err
        }

        createdTask := s.repository.Create(task)

        return createdTask, nil
}

func (s *TaskService) GetTask(id int) (*Task, error) {
        task, ok := s.repository.Get(id)

        if !ok {
                return nil, errors.New("Task not found")
        }
        return task, nil
}

func (s *TaskService) GetTasks() map[int]*Task {
        return s.repository.GetALL()
}

func (s *TaskService) UpdateTask(id int, task *Task) (*Task, error) {
        if err := validate(*task); err != nil {
                return nil, err
        }

        updatedTask, ok := s.repository.Update(id, task)

        if !ok {
                return nil, errors.New("Task not Found")
        }

        return updatedTask, nil
}

func (s *TaskService) DeleteTask(id int) error {
        ok := s.repository.Delete(id)

        if !ok {
                return errors.New("task not found")
        }
        return nil
}

