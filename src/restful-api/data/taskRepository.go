package data

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"restful-api/models"
	"time"
)

type TaskRepository struct {
	C *mgo.Collection
}

func (r *TaskRepository) Create(task *models.Task) error {
	obj_id := bson.NewObjectId()
	task.Id = obj_id
	task.CreatedOn = time.Now()
	task.Status = "Created"
	err := r.C.Insert(&task)
	return err
}
func (r *TaskRepository) Update(task *models.Task) error {
	err := r.C.Update(bson.M{"_id": task.id}, bson.M{"$set": bson.M{
		"name":        task.Name,
		"description": task.Description,
		"due":         task.Due,
		"status":      task.Status,
		"tags":        task.Tags,
	}})
	return err
}

func (r *TaskRepository) Delete(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
}
func (r *TaskRepository) GetAll() []models.Task {
	var task []models.Task
	iter := r.C.Find(nil).Iter()
	result := models.Task{}
	for iter.Next(&result) {
		tasks = append(tasks, result)
	}
	return tasks
}
func (r *TaskRepository) GetById(user string) []models.Task {
	var task = models.Task{}
	err = r.C.FindId(bson.ObjectIdHex(id)).One(&task)
	if err != nil {
		return "", err
	}
	return task, nil
}
func (r *TaskRepository) GetByUser(user string) []models.Task {
	var tasks []models.Task
	iter := r.C.Find(bson.M{"createdby": user}).Iter()
	result := models.Task{}
	for iter.Next(&result) {
		tasks = append(tasks, result)
	}
	return tasks
}
