package main

import (
	"bufio"
	"fmt"
	"os"
)

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
		id:          RandSeq(8),
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

func RunTasksTracker() {
	tracker := NewTracker()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Press 1 to print all your tasks")
		fmt.Println("Press 2 to add new task")
		fmt.Println("Press 3 to delete a task")
		fmt.Println("Press 0 to exit the program")
		fmt.Print("Enter your choice: ")
		if scanner.Scan() {
			choice := scanner.Text()

			switch choice {
			case "1":
				tracker.printTasks()
			case "2":
				fmt.Print("Enter task description: ")
				if scanner.Scan() {
					newTask := scanner.Text()
					tracker.addTask(newTask)
				}
			case "3":
				fmt.Print("Enter task id: ")
				if scanner.Scan() {
					taskId := scanner.Text()
					tracker.deleteTaskByID(taskId)
				}
			case "0":
				fmt.Println("Exiting... 👋")
				return
			default:
				fmt.Println("Invalid option 🤔")
			}
		}
	}
}
