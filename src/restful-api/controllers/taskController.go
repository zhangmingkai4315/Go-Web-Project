package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	// "log"
	"net/http"
	"restful-api/common"
	"restful-api/data"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var dataResource TaskResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "invalid task data", 500)
		return
	}
	task := &dataResource.Data
	context := NewContext()
	defer context.Close()

	c := context.DbCollection("tasks")
	repo := &data.TaskRepository{c}
	repo.Create(task)
	if j, err := json.Marshal(TaskResource{Data: *task}); err != nil {
		common.DisplayAppError(w, err, "An Unexpected error has occurred", 500)
		return
	} else {
		w.Header().Set("Content-Type", "appliction/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(j)
	}
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	defer context.Close()

	c := context.DbCollection("tasks")
	repo := &data.TaskRepository{c}

	tasks := repo.GetAll()
	j, err := json.Marshal(TasksResource{Data: tasks})
	if err != nil {
		common.DisplayAppError(w, err, "An Unexpected error has occurred", 500)
		return
	} else {
		w.Header().Set("Content-Type", "appliction/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}

func GetTaskById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	context := NewContext()
	defer context.Close()

	c := context.DbCollection("tasks")
	repo := &data.TaskRepository{c}

	task, err := repo.GetById(id)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNoContent)
			return
		} else {
			common.DisplayAppError(w, err, "An Unexpected error has occurred", 500)
			return
		}
	}
	j, err := json.Marshal(task)
	if err != nil {
		common.DisplayAppError(w, err, "An Unexpected error has occurred", 500)
		return
	} else {
		w.Header().Set("Content-Type", "appliction/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(j)
	}
}

func GetTasksByUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	user := vars["id"]
	context := NewContext()
	defer context.Close()

	c := context.DbCollection("tasks")
	repo := &data.TaskRepository{c}

	tasks := repo.GetByUser(user)

	j, err := json.Marshal(TasksResource{Data: tasks})
	if err != nil {
		common.DisplayAppError(w, err, "An Unexpected error has occurred", 500)
		return
	} else {
		w.Header().Set("Content-Type", "appliction/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(j)
	}
}

func UpdateTaskById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := bson.ObjectIdHex(vars["id"])

	var dataResource TaskResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid task data", 500)
		return
	}
	task := &dataResource.Data
	task.Id = id
	context := NewContext()
	defer context.Close()

	c := context.DbCollection("tasks")
	repo := &data.TaskRepository{c}

	err = repo.Update(task)
	if err != nil {
		common.DisplayAppError(w, err, "An Unexpected error has occurred", 500)
		return
	} else {
		// w.Header().Set("Content-Type", "appliction/json")
		w.WriteHeader(http.StatusOK)
		// w.Write(j)
	}
}

func DeleteTaskById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	context := NewContext()
	defer context.Close()

	c := context.DbCollection("tasks")
	repo := &data.TaskRepository{c}

	err := repo.Delete(id)
	if err != nil {
		common.DisplayAppError(w, err, "An Unexpected error has occurred", 500)
		return
	} else {
		// w.Header().Set("Content-Type", "appliction/json")
		w.WriteHeader(http.StatusOK)
		// w.Write(j)
	}
}
