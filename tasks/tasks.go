package task

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// Esta es como la clase Task o entidad Task
type Task struct {
	ID       int    `json:"id"`
	NAME     string `json:"name"`
	COMPLETE bool   `json:"complete"`
}

// Esta funcion lista las tareas
func ListTask(tasks []Task) {
	if (len(tasks)) == 0 {
		fmt.Println("No tasks")
	}

	for _, task := range tasks {

		status := " "

		if task.COMPLETE {
			status = "✓"
		}

		fmt.Printf("[%s] %d %s", status, task.ID, task.NAME)
	}
}

// Esta funcion agrega una tarea
func AddTask(tasks []Task, name string) []Task {
	newTask := Task{
		ID:       len(tasks) + 1,
		NAME:     name,
		COMPLETE: false,
	}
	tasks = append(tasks, newTask)
	return tasks
}

// Esta funcion guarda una tarea en el archivo Tasks.json
func SaveTask(file *os.File, tasks []Task) {

	bytes, err := json.Marshal(tasks)

	if err != nil {
		panic(err)
	}

	_, err = file.Seek(0, 0)

	if err != nil {
		panic(err)
	}

	err = file.Truncate(0)

	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(file)
	_, err = writer.Write(bytes)

	if err != nil {
		panic(err)
	}

	err = writer.Flush()

	if err != nil {
		panic(err)
	}

	fmt.Println("Saving tasks")
}

// Esta funcion completa una tarea
func CompleteTask(tasks []Task, id int) []Task {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].COMPLETE = true
			break
		}
	}
	return tasks
}

func DeleteTask(tasks []Task, id int) []Task {
	for i, task := range tasks {
		if task.ID == id {
			return append(tasks[:i], tasks[i+1:]...)
		}
	}
	return tasks
}
