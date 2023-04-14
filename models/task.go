/**
 * @author Chi-Chen Chang (ccchang915@gmail.com)
 */

package models

import (
	"time"
)

type Task struct {
	Id          int       `gorm:"type:bigint(20) auto_increment;primary_key;" json:"id"`
	Title       string    `gorm:"type:varchar(100) NOT NULL;" json:"title"`
	Description string    `gorm:"type:varchar(200) NOT NULL;" json:"description"`
	Completed   bool      `gorm:"default:false" json:"completed"`
	CreatedAt   time.Time `gorm:"type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
}
