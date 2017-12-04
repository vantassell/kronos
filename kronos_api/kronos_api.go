package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"github.com/gorilla/mux"
	"gopkg.in/robfig/cron.v2"
)

// models
type Task struct {
	ID        int    `json:"id,omitempty"`
	Image     *Image `json:"image"`
	Frequency string `json:"frequency"`
}

type Image struct {
	Tag string `json:"tag"`
}

// members
var scheduledTasks []Task
var _cron *cron.Cron

func main() {
	setup()
}

// initial setup
func setup() {
	_cron = cron.New()
	_cron.Start()

	router := mux.NewRouter()
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks", createTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// lists all scheduled tasks
func getTasks(w http.ResponseWriter, r *http.Request) {
	// return result
	json.NewEncoder(w).Encode(scheduledTasks)
}

// creates a scheduled task
func createTask(w http.ResponseWriter, r *http.Request) {
	// unpack request
	//	params := mux.Vars(r)
	var task Task
	_ = json.NewDecoder(r.Body).Decode(&task)

	// create task
	id, err := _cron.AddFunc(task.Frequency, func() {
		out, _ := exec.Command("docker", "run", task.Image.Tag).Output()
		fmt.Printf("task #%v - output: %s", task.ID, out)
	})
	check(err)

	// update local model
	task.ID = int(id)
	scheduledTasks = append(scheduledTasks, task)

	// return result
	json.NewEncoder(w).Encode(scheduledTasks)
}

// deletes a running task
func deleteTask(w http.ResponseWriter, r *http.Request) {
	// unpack request
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
	}

	// delete matching cron
	for index, item := range scheduledTasks {
		if item.ID == id {
			var entryId cron.EntryID
			entryId = cron.EntryID(id)

			if _cron.Entry(entryId).Job != nil {
				_cron.Remove(entryId)
				scheduledTasks = append(scheduledTasks[:index], scheduledTasks[index+1:]...)
			}
			break
		}
	}

	// return result
	json.NewEncoder(w).Encode(scheduledTasks)
}

// error chcker
func check(e error) {
	if e != nil {
		log.Fatal(e)
		panic(e)
	}
}
