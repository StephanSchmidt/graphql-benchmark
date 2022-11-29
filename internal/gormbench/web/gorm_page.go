package web

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"inkmi/internal/gormbench"
	"math/rand"

	//	"math/rand"
	"net/http"
)

func GormPage(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := rand.Intn(100) + 1
		var tasks []gormbench.Task
		db.Where("user_id = ?", userId).Preload("User").Preload("Status").Find(&tasks)
		taskList := gormbench.GormTaskList{
			Tasks: tasks,
		}
		return c.Render(http.StatusOK, "gorm.html", taskList)
	}
}
