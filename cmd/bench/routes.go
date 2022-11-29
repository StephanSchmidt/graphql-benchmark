package main

import (
	"github.com/dosco/graphjin/core"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	gormweb "inkmi/internal/gormbench/web"
	gql "inkmi/internal/graphqlbench/web"
	pgx "inkmi/internal/pgxbench/web"
	sqlxweb "inkmi/internal/sqlxbench/web"
)

func configRoutes(e *echo.Echo, pool *pgxpool.Pool, gj *core.GraphJin, db *gorm.DB, xdb *sqlx.DB) {
	e.GET("/pgx", pgx.PgxPage(pool)).Name = "Pgx"
	e.GET("/sqlx", sqlxweb.SqlxPage(xdb)).Name = "Sqlx"
	e.GET("/graphql", gql.GraphQlPage(pool, gj)).Name = "Graphql"
	e.GET("/gorm", gormweb.GormPage(db)).Name = "Gorm"
}
