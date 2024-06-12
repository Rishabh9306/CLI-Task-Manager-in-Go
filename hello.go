package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	Description string
	Completed   bool
}

type TaskManager struct {
	Tasks []Task
}

const (
	dataFile = "tasks.json"
)

func main() {
	manager := TaskManager{}
	manager.LoadTasks()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nTask Manager")
		fmt.Println("1. Add task")
		fmt.Println("2. List tasks")
		fmt.Println("3. Complete task")
		fmt.Println("4. Delete task")
		fmt.Println("5. Edit task")
		fmt.Println("6. Save and Exit")
		fmt.Print("Choose an option: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		input = strings.TrimSpace(input)

		switch input {
		case "1":
			manager.AddTask(reader)
		case "2":
			manager.ListTasks()
		case "3":
			manager.CompleteTask(reader)
		case "4":
			manager.DeleteTask(reader)
		case "5":
			manager.EditTask(reader)
		case "6":
			manager.SaveTasks()
			fmt.Println("Tasks saved. Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid input")
		}
	}
}

func (manager *TaskManager) AddTask(reader *bufio.Reader) {
	fmt.Print("Enter the task description: ")
	description, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	description = strings.TrimSpace(description)
	manager.Tasks = append(manager.Tasks, Task{Description: description})
	fmt.Println("Task added")
}

func (manager *TaskManager) ListTasks() {
	if len(manager.Tasks) == 0 {
		fmt.Println("No tasks")
	} else {
		fmt.Println("Tasks:")
		for i, task := range manager.Tasks {
			status := "Incomplete"
			if task.Completed {
				status = "Complete"
			}
			fmt.Printf("%d. %s [%s]\n", i+1, task.Description, status)
		}
	}
}

func (manager *TaskManager) CompleteTask(reader *bufio.Reader) {
	if len(manager.Tasks) == 0 {
		fmt.Println("No tasks to complete")
		return
	}
	fmt.Print("Enter the task number to complete: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	index, err := getIndex(input, len(manager.Tasks))
	if err != nil {
		fmt.Println("Invalid task number:", err)
		return
	}

	manager.Tasks[index].Completed = true
	fmt.Println("Task marked as complete")
}

func (manager *TaskManager) DeleteTask(reader *bufio.Reader) {
	if len(manager.Tasks) == 0 {
		fmt.Println("No tasks to delete")
		return
	}
	fmt.Print("Enter the task number to delete: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	index, err := getIndex(input, len(manager.Tasks))
	if err != nil {
		fmt.Println("Invalid task number:", err)
		return
	}

	manager.Tasks = append(manager.Tasks[:index], manager.Tasks[index+1:]...)
	fmt.Println("Task deleted")
}

func (manager *TaskManager) EditTask(reader *bufio.Reader) {
	if len(manager.Tasks) == 0 {
		fmt.Println("No tasks to edit")
		return
	}
	fmt.Print("Enter the task number to edit: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	index, err := getIndex(input, len(manager.Tasks))
	if err != nil {
		fmt.Println("Invalid task number:", err)
		return
	}

	fmt.Print("Enter the new task description: ")
	description, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	manager.Tasks[index].Description = strings.TrimSpace(description)
	fmt.Println("Task description updated")
}

func (manager *TaskManager) SaveTasks() {
	file, err := os.Create(dataFile)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(manager.Tasks)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}
}

func (manager *TaskManager) LoadTasks() {
	file, err := os.Open(dataFile)
	if err != nil {
		fmt.Println("No existing tasks found")
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&manager.Tasks)
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}
}

func getIndex(input string, length int) (int, error) {
	input = strings.TrimSpace(input)
	index, err := strconv.Atoi(input)
	if err != nil || index < 1 || index > length {
		return -1, fmt.Errorf("invalid index")
	}
	return index - 1, nil
}
