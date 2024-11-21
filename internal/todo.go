package todo

import (
	"log"
)

type Task struct {
	Name      string
	Completed bool
}

var Tasks []Task

func AppendTask(name string) {
	Tasks = append(Tasks, Task{Name: name, Completed: false})
	log.Printf("Task successfully added")
}
