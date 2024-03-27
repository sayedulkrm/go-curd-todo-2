package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/sayedulkrm/go-curd-todo-2/models"
	"github.com/sayedulkrm/go-curd-todo-2/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Get All the tasks

func GetAllTasks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte("Get all tasks"))

}

// CreateNewTask creates a new task

func CreateNewTask(w http.ResponseWriter, r *http.Request) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	defer cancel()

	var task models.TodoList

	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		utils.LogError(r, err)
		utils.ErrorResponse(w, r, http.StatusBadRequest, "failed to decode")
		return
	}

	// Check for missing fields
	if task.Task == "" || task.Status == "" {
		utils.ErrorResponse(w, r, http.StatusBadRequest, "Task and Status fields are required")
		return
	}

	task.Id = primitive.NewObjectID()

	_, err = models.TodoCollection.InsertOne(ctx, task)

	if err != nil {
		utils.LogError(r, err)
		utils.ErrorResponse(w, r, http.StatusBadRequest, "failed to create task")
		return
	}

	responseData := struct {
		Success bool        `json:"success"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}{

		Success: true,
		Message: "Task created successfully",
		Data:    task,
	}

	// Send success response with additional data
	utils.SuccessResponse(w, http.StatusCreated, responseData)

}

// Update Task

func UpdateTask(w http.ResponseWriter, r *http.Request) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	defer cancel()

	var updatedTask models.TodoList

	err := json.NewDecoder(r.Body).Decode(&updatedTask)

	if err != nil {
		utils.LogError(r, err)
		utils.ErrorResponse(w, r, http.StatusBadRequest, "failed to decode")
		return
	}

	// Updated by specific feild

	// update := bson.M{
	// 	"$set": bson.M{
	// 		"fieldName": "newValue", // Specify the field name and its new value
	// 	},
	// }

	var foundTask models.TodoList
	err = models.TodoCollection.FindOne(ctx, bson.M{"_id": updatedTask.Id}).Decode(&foundTask)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			utils.ErrorResponse(w, r, http.StatusNotFound, "task not found")
			return
		}
		utils.LogError(r, err)
		utils.ErrorResponse(w, r, http.StatusInternalServerError, "failed to find task")
		return
	}

	updateData := bson.M{
		"$set": updatedTask, // Use the updatedTask directly, assuming it contains only the fields you want to update
	}

	_, err = models.TodoCollection.UpdateByID(ctx, updatedTask.Id, updateData)

	if err != nil {
		utils.LogError(r, err)
		utils.ErrorResponse(w, r, http.StatusBadRequest, "failed to Updated task")
		return
	}

	responseData := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{

		Success: true,
		Message: "Task Updated successfully",
	}

	// Send success response with additional data
	utils.SuccessResponse(w, http.StatusCreated, responseData)

}

// Delete task

func DeleteTask(w http.ResponseWriter, r *http.Request) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	defer cancel()

	var deleteTask struct {
		Id primitive.ObjectID `json:"_id"`
	}

	err := json.NewDecoder(r.Body).Decode(&deleteTask)

	if err != nil {
		utils.LogError(r, err)
		utils.ErrorResponse(w, r, http.StatusBadRequest, "failed to decode")
		return
	}

	var foundTask models.TodoList
	err = models.TodoCollection.FindOne(ctx, bson.M{"_id": deleteTask.Id}).Decode(&foundTask)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			utils.ErrorResponse(w, r, http.StatusNotFound, "task not found")
			return
		}
		utils.LogError(r, err)
		utils.ErrorResponse(w, r, http.StatusInternalServerError, "failed to find task")
		return
	}

	_, err = models.TodoCollection.DeleteOne(ctx, bson.M{"_id": deleteTask.Id})

	if err != nil {
		utils.LogError(r, err)
		utils.ErrorResponse(w, r, http.StatusBadRequest, "failed to Deleted task")
		return
	}

	responseData := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{

		Success: true,
		Message: "Task Deleted successfully",
	}

	// Send success response with additional data
	utils.SuccessResponse(w, http.StatusCreated, responseData)

}
