/**
 * @author Chi-Chen Chang (ccchang915@gmail.com)
 */

package repositories

import (
	"simple-todo-list/db"
	"simple-todo-list/models"
)

func FindAllTasks() (tasks []models.Task, err error) {
	err = db.Session.Find(&tasks).Error
	return
}

func FindTaskById(id int) (task models.Task, err error) {
	err = db.Session.First(&task, id).Error
	return
}

func CreatTask(task models.Task) (err error) {
	err = db.Session.Create(&models.Task{Title: task.Title, Description: task.Description, Completed: task.Completed}).Error
	return
}

func DeleteTaskById(id int) (err error) {
	err = db.Session.Delete(&models.Task{}, id).Error
	return
}

func UpdateTask(task models.Task) (err error) {
	err = db.Session.Model(&task).Where("id = ?", task.Id).Updates(map[string]interface{}{"title": task.Title, "description": task.Description, "completed": task.Completed}).Error
	return
}
