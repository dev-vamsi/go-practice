package main

import "fmt"

type Task struct {
	description string
	id          string
}

type TasksTracker struct {
	arr []Task
}

func NewTracker() TasksTracker {
	return TasksTracker{}
}

func (t *TasksTracker) addTask(task string) {
	t.arr = append(t.arr, Task{
		description: task,
		id:          randSeq(8),
	})
}

func (t *TasksTracker) deleteTaskByID(id string) {
	var arr []Task
	for _, value := range t.arr {
		if value.id != id {
			arr = append(arr, value)
		}
	}
	t.arr = arr
}

func (t *TasksTracker) printTasks() {
	if len(t.arr) == 0 {
		fmt.Println("No tasks...")
		return
	}
	for _, value := range t.arr {
		fmt.Printf("%v: %v\n", value.id, value.description)
	}
}
