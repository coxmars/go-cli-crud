package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	task "github.com/coxmars/go-cli-crud/tasks"
)

// This is the main function, here we call other functions
func main() {
	file, err := os.OpenFile("Tasks.json", os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	var tasks []task.Task

	info, err := file.Stat()

	if err != nil {
		panic(err)
	}

	if info.Size() != 0 {
		bytes, err := io.ReadAll(file) // Da los datos en bytes
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(bytes, &tasks) // Convierte los bytes en un slice de tareas

		if err != nil {
			panic(err)
		}

	} else {
		tasks = []task.Task{}
	}

	if len(os.Args) < 2 {
		printUsage()
		return
	}

	switch os.Args[1] {
	case "list":
		task.ListTask(tasks)
		return
	case "add":
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("What is your task? ")
		name, _ := reader.ReadString('\n')
		strings.TrimSpace(name)
		tasks = task.AddTask(tasks, name)
		task.SaveTask(file, tasks)
		fmt.Println("Added task")
		return
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Missing task id")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Id should be a number")
			return
		}
		tasks = task.DeleteTask(tasks, id)
		task.SaveTask(file, tasks)
	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Missing task id")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Id should be a number")
			return
		}
		tasks = task.CompleteTask(tasks, id)
		task.SaveTask(file, tasks)
	default:
		printUsage()
		return
	}

}

// This is like the menu
func printUsage() {
	fmt.Println("Usage: go-cli-crud <list, add, complete, delete>")
}
