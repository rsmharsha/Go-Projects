package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// Printing dynamic table foe listing tasks
func printTasksTable(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	idWidth := len("ID")
	descWidth := len("Description")
	statusWidth := len("Status")

	// calculate max widths
	for _, t := range tasks {
		if len(fmt.Sprintf("%d", t.Id)) > idWidth {
			idWidth = len(fmt.Sprintf("%d", t.Id))
		}
		if len(t.Description) > descWidth {
			descWidth = len(t.Description)
		}
		if len(t.Status) > statusWidth {
			statusWidth = len(t.Status)
		}
	}

	printBorder := func() {
		fmt.Println(
			"+" + strings.Repeat("-", idWidth+2) +
				"+" + strings.Repeat("-", descWidth+2) +
				"+" + strings.Repeat("-", statusWidth+2) + "+",
		)
	}

	printBorder()
	fmt.Printf("| %-*s | %-*s | %-*s |\n", idWidth, "ID", descWidth, "Description", statusWidth, "Status")
	printBorder()

	for _, t := range tasks {
		fmt.Printf("| %-*d | %-*s | %-*s |\n",
			idWidth, t.Id,
			descWidth, t.Description,
			statusWidth, t.Status,
		)
	}

	printBorder()
}

// loading tasks - decoding JSON

func loadTasks() []Task {
	data, err := os.ReadFile("tasks.json")

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exist")
			return []Task{}
		} else {
			fmt.Println("Error: ", err)
			return []Task{}
		}
	}

	var tasks []Task

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		panic(err)
	}

	return tasks

}

// saving tasks - encoding JSON

func saveTasks(tasks []Task) {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		fmt.Println("Error marshaling:", err)
		return
	}

	err = os.WriteFile("tasks.json", data, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

}

// adding a new task

func addTask(description string) {
	tasks := loadTasks()

	newTask := Task{
		Id:          len(tasks) + 1,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, newTask)

	saveTasks(tasks)

	fmt.Printf("Task added successfully (ID: %d)\n", newTask.Id)

}

// listing all the tasks

func listTasks() {
	tasks := loadTasks()
	printTasksTable(tasks)
}

// deleting a task

func deleteTask(id int) {
	tasks := loadTasks()

	for i, task := range tasks {
		if task.Id == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			saveTasks(tasks)
			fmt.Printf("Task %d deleted successfully\n", id)
			return
		}
	}
	fmt.Printf("Task with ID %d not found\n", id)
}

// updating a task

func updateTask(id int, newDescription string) {
	tasks := loadTasks()

	for i, task := range tasks {
		if task.Id == id {
			tasks[i].Description = newDescription
			tasks[i].UpdatedAt = time.Now()
			saveTasks(tasks)
			fmt.Printf("Task %d updated successfully\n", id)
			return
		}
	}
	fmt.Printf("Task with ID %d not found\n", id)
}

// marking a task - changing status

func markTask(id int, status string) {
	tasks := loadTasks()

	for i, task := range tasks {
		if task.Id == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()
			saveTasks(tasks)
			fmt.Printf("Task %d marked successfully\n", id)
			return

		}
	}
	fmt.Printf("Task with ID %d not found\n", id)
}

func main() {
	fmt.Println("Welcome to Task-Tracker")

	if len(os.Args) < 2 {
		fmt.Println("Usage: task-cli <command>")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("usage: task-cli add <description>")
			os.Exit(1)
		}
		addTask(os.Args[2])

	case "list":
		if len(os.Args) == 3 {
			tasks := loadTasks()
			var filtered []Task
			for _, task := range tasks {
				if task.Status == os.Args[2] {
					filtered = append(filtered, task)
				}
			}
			printTasksTable(filtered)
		} else {
			listTasks()
		}

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("usage: task-cli delete <id>")
			os.Exit(1)
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid ID — must be a number")
			os.Exit(1)
		}
		deleteTask(id)

	case "update":
		if len(os.Args) < 4 {
			fmt.Println("usage: task-cli update <id> <description>")
			os.Exit(1)
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid ID — must be a number")
			os.Exit(1)
		}

		updateTask(id, os.Args[3])

	case "mark-in-progress":
		status := "in-progress"

		if len(os.Args) < 3 {
			fmt.Println("usage: task-cli marks-in-progress <id>")
			os.Exit(1)
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid ID — must be a number")
			os.Exit(1)
		}

		markTask(id, status)

	case "mark-done":
		status := "done"

		if len(os.Args) < 3 {
			fmt.Println("usage: task-cli marks-done <id>")
			os.Exit(1)
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid ID — must be a number")
			os.Exit(1)
		}

		markTask(id, status)

	default:
		fmt.Println("Unknown command:", command)

	}

}
