/**
 * @author Chi-Chen Chang (ccchang915@gmail.com)
 */

package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-todo-list/models"
	"simple-todo-list/repositories"
	"strconv"
	"time"
)

// GetTasks godoc
//	@Summary		Get all tasks
//	@Description	get a list of tasks
//	@Tags			task
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	models.Task
//	@Router			/todos [get]
func GetTasks(ctx *gin.Context) {
	tasks, err := repositories.FindAllTasks()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "tasks not found",
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, tasks)
}

// GetTask godoc
//	@Summary		Get a task
//	@Description	get task by id
//	@Tags			task
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Task ID"
//	@Success		200	{object}	models.Task
//	@Router			/todos/{id} [get]
func GetTask(ctx *gin.Context) {
	id := ctx.Params.ByName("id")

	taskId, err := strconv.Atoi(id)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	task, err := repositories.FindTaskById(taskId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "task not found",
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, task)
}

// CreateTask godoc
//	@Summary		Create a new task
//	@Description	create a new task
//	@Tags			task
//	@Accept			json
//	@Param			task	body	models.Task	true	"Create Task"
//	@Success		201
//	@Router			/todo [post]
func CreateTask(ctx *gin.Context) {
	var task models.Task

	err := ctx.BindJSON(&task)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	task.CreatedAt = time.Now()
	err = repositories.CreatTask(task)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusCreated)
}

// DeleteTask godoc
//	@Summary		Delete a task
//	@Description	delete task by id
//	@Tags			task
//	@Accept			json
//	@Param			id	path	int	true	"Task ID"
//	@Success		204
//	@Router			/todos/{id} [delete]
func DeleteTask(ctx *gin.Context) {
	id := ctx.Params.ByName("id")

	taskId, err := strconv.Atoi(id)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	err = repositories.DeleteTaskById(taskId)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusNoContent)
}

//UpdateTask godoc
//	@Summary		Update task
//	@Description	Update task by ID
//	@Tags			task
//	@Accept			json
//	@Produce		json
//	@Param			id		path	int			true	"Task ID"
//	@Param			task	body	models.Task	true	"Task Data"
//	@Success		204
//	@Router			/todos/{id} [put]
func UpdateTask(ctx *gin.Context) {
	id := ctx.Params.ByName("id")

	taskId, err := strconv.Atoi(id)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	var task models.Task

	err = ctx.BindJSON(&task)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	_, err = repositories.FindTaskById(taskId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "task not found",
		})
		return
	}

	task.Id = taskId
	err = repositories.UpdateTask(task)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusNoContent)
}
