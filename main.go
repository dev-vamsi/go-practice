package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyz1234567890_/-")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func main() {
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
