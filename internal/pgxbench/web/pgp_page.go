package web

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
	"inkmi/internal/pgxbench"
	"math/rand"
	"net/http"
)

func PgxPage(pool *pgxpool.Pool) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := rand.Intn(100) + 1
		rows, err := pool.Query(context.Background(),
			"select t.id,t.title,u.name,s.status from tasks t, users u, status s where t.user_id=$1 and t.user_id=u.id and t.status_id=s.id", user)
		if err != nil {
			panic(err)
		}
		p := pgxbench.PgxTaskList{}
		tasks := make([]pgxbench.PgxTask, 0)
		for rows.Next() {
			var id int64
			var title string
			var status string
			var user string
			err := rows.Scan(&id, &title, &user, &status)
			if err != nil {
				fmt.Println(err)
			}
			t := pgxbench.PgxTask{
				Id:     id,
				Title:  title,
				User:   user,
				Status: status,
			}
			tasks = append(tasks, t)
		}
		p.Tasks = tasks
		return c.Render(http.StatusOK, "pgx.html", p)
	}
}
