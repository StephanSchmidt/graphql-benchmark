package web

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"inkmi/internal/sqlxbench"
	"math/rand"
	"net/http"
)

func SqlxPage(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := rand.Intn(100) + 1
		tasks := []sqlxbench.SqlxTask{}
		err := db.Select(&tasks,
			"select t.id Id,t.title Title,u.name Name,s.status Status from tasks t, users u, status s where t.user_id=u.id and t.status_id=s.id and t.user_id=$1", user)
		if err != nil {
			fmt.Println("ERROR!")
			panic(err)
		}
		p := sqlxbench.SqlxTaskList{}
		p.Tasks = tasks
		return c.Render(http.StatusOK, "sqlx.html", p)
	}
}
