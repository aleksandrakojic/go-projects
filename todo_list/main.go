package main

import (
	"fmt"
	"net/http"
)


var shortGolang = "Learn Go: A Beginner's Guide"
var fullGolang = "Master Go: A Comprehensive Guide"
var rewards = "Earn $100 for completing a task"


var taskItems = []string{shortGolang, fullGolang, rewards}

func main() {

	http.HandleFunc("/hello-go", helloUser)
	http.HandleFunc("/show-tasks", showTasks)

	http.ListenAndServe(":8080", nil)
}

func helloUser(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, welcome to our Todolist App!")
}

func showTasks(w http.ResponseWriter, r *http.Request) {
    for _, task := range taskItems {
		fmt.Fprintln(w, task)
	}
    
}


func addTask(taskItems []string, newTask string) ([]string) {
	taskItems = append(taskItems, newTask)
    fmt.Printf("Task '%s' has been added.\n", newTask)
    return taskItems
}