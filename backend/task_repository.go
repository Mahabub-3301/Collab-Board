package main

type TaskRepository interface {
        Create(task *Task) *Task
        GetALL()        map[int]*Task
        Get(id int)     (*Task,bool)
        Update(id int, task *Task) (*Task,bool)
        Delete(id int) bool
}
