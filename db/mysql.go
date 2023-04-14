/**
 * @author Chi-Chen Chang (ccchang915@gmail.com)
 */

package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"simple-todo-list/models"
	"strconv"
)

var Session *gorm.DB

func InitializeMySQL() error {
	username := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	endpoint := os.Getenv("MYSQL_HOST")
	port, err := strconv.Atoi(os.Getenv("MYSQL_PORT"))
	database := os.Getenv("MYSQL_DATABASE")

	Session, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True", username, password, endpoint, port, database)), &gorm.Config{})
	if err != nil {
		return err
	}

	err = Session.AutoMigrate(&models.Task{})
	if err != nil {
		return err
	}

	return nil
}
