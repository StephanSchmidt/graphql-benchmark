package web

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dosco/graphjin/core"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
	"math/rand"
	"net/http"
)

func GraphQlPage(pool *pgxpool.Pool, gj *core.GraphJin) echo.HandlerFunc {
	return func(c echo.Context) error {
		query := `
query GetTasks {
  tasks(limit: 1000,  where: { user_id: $userId } ){
    id
    title
	user {
		id
		name
	}
    status {
		id
		status
    }
  }
}`
		user := rand.Intn(100) + 1
		vars := json.RawMessage(fmt.Sprintf(`{ "userId" : %d }`, user))
		res, err := gj.GraphQL(context.Background(), query, vars, nil)
		if err != nil {
			panic(err)
		}
		var dat map[string]interface{}
		if err := json.Unmarshal(res.Data, &dat); err != nil {
			panic(err)
		}
		return c.Render(http.StatusOK, "graphql.html", dat)
	}
}
