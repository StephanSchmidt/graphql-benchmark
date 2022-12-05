package web

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
	"inkmi/internal"
	"math/rand"
	"net/http"
)

func render(c echo.Context, pool *pgxpool.Pool, template string, query string, args ...any) error {
	row, err := pool.Query(
		context.Background(),
		query, args...)
	if err != nil {
		panic(err)
	}
	defer row.Close()
	row.Next()
	var json map[string]interface{}
	if err := row.Scan(&json); err != nil {
		panic(err)
	}
	return c.Render(http.StatusOK, template, json)
}

func JsonPgxPage(pool *pgxpool.Pool) echo.HandlerFunc {
	query := internal.Replacer(
		`
			WITH tasks AS $(
				SELECT
					t.id id,
					t.title title,
					u.name name,
					s.status status
					FROM tasks t, users u, status s
					WHERE t.user_id=u.id AND t.status_id=s.id
					AND t.user_id=$1
			)
			$$(
				'tasks',  (SELECT * from tasks)
			)
		`)
	return func(c echo.Context) error {
		user := rand.Intn(100) + 1
		return render(c, pool, "jsonpgx.html", query, user)
	}
}
