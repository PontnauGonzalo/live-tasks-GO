package cmd

import (
	"github.com/gin-gonic/gin"
)

type GlobalHandler struct {
	taskHandler TaskHandler
}

func NewGlobalHandler(taskHandler TaskHandler) *GlobalHandler {
	return &GlobalHandler{
		taskHandler: taskHandler,
	}
}

func (globalHandler *GlobalHandler) BuildRoutes(engine *gin.Engine) {
	globalHandler.BuildTaskRoutes(engine)
}

func (globalHandler *GlobalHandler) BuildTaskRoutes(engine *gin.Engine) {
	group := engine.Group("/tasks")
	group.POST("/", globalHandler.taskHandler.CreateTask())
}



