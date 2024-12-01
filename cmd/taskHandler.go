package cmd

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type GlobalHandler struct {
	serv *task.Service
}

type CreateRequest struct {
	Title 		string `json:"title"`
	Description string `json:"description"`
	Stage 		string `json:"stage"`
	CreatedBy 	string `json:"created_by"`
	AssignedTo 	string `json:"assigned_to"`
	Tags 		[]string `json:"tags"`
}

func NewTaskHandler(serv *task.Service) *TaskHandler {
	return &TaskHandler{
		serv: serv,
	}
}

func (taskHandler *TaskHandler) CreateTask() gin.HandlerFunc {
	return func(context *gin.Context) {
		newTask := &CreateRequest{}
		err := context.ShouldBindJSON(newTask)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		task := domain.Task {  // instance of domain.Task
			Title: 		 newTask.Title,
			Description: newTask.Description,
			Stage: 		 newTask.Stage,
			CreatedBy: 	 newTask.CreatedBy,
			AssignedTo:  newTask.AssignedTo,
			Tags: 		 newTask.Tags,
		}
		t, err := taskHandler.serv.CreateTask(context, task)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		context.JSON(http.StatusCreated, t)
	}
}
