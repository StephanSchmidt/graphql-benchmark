package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dosco/graphjin/core"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	e := echo.New()
	e.Debug = true
	pool := getPool()
	defer pool.Close()

	dbConn, err := sql.Open("pgx", "postgres://graphjin:graphjin@localhost:5432/graphjin")
	if err != nil {
		fmt.Println(err)
	}
	dbConn.SetMaxOpenConns(20)
	dbConn.SetMaxIdleConns(10)
	dbConn.SetConnMaxLifetime(time.Hour)
	defer dbConn.Close()
	conf := &core.Config{
		Production:       true,
		DisableAllowList: true,
	}
	gj, err := core.NewGraphJin(conf, dbConn)
	gj.IsProd()

	db, err := gorm.Open(postgres.Open("postgres://graphjin:graphjin@localhost:5432/graphjin"), &gorm.Config{})
	sqlDB, err := db.DB()
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(20)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)
	defer sqlDB.Close()

	xdb, err := sqlx.Open("pgx", "postgres://graphjin:graphjin@localhost:5432/graphjin")
	if err != nil {
		panic(err)
	}
	xdb.DB.SetMaxOpenConns(20) // The default is 0 (unlimited)
	xdb.DB.SetMaxIdleConns(10) // defaultMaxIdleConns = 2
	xdb.DB.SetConnMaxLifetime(time.Hour)
	defer xdb.Close()

	configRenderer(e)
	configRoutes(e, pool, gj, db, xdb)
	WriteFiles()

	// Start server
	go func() {
		if err := e.Start("localhost:1323"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended
	// for signal.Notify

	shutdownTime := 1 * time.Second
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTime)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

}

func getPool() *pgxpool.Pool {
	dsn := "postgres://graphjin:graphjin@localhost:5432/graphjin?pool_max_conns=20"
	pool, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return pool
}

func WriteFiles() {
	f, err := os.Create("data.sql")

	if err != nil {
		os.Exit(1)
	}

	defer f.Close()

	maxusers := 100
	maxstatus := 7
	maxtasks := 10000

	for i := 1; i <= maxstatus; i++ {
		sql1 := fmt.Sprintf("INSERT INTO status (id, status) VALUES (%d, 'status-%d');\n", i, i)
		sql2 := fmt.Sprintf("INSERT INTO statuses (id, status) VALUES (%d, 'status-%d');\n", i, i)
		f.WriteString(sql1)
		f.WriteString(sql2)
	}

	for i := 1; i <= maxusers; i++ {
		sql := fmt.Sprintf("INSERT INTO users (id, name) VALUES (%d, 'user-%d');\n", i, i)
		f.WriteString(sql)
	}

	for i := 1; i <= maxtasks; i++ {
		sql := fmt.Sprintf("INSERT INTO tasks (id, title, user_id, status_id) VALUES (%d, 'task-%d', %d, %d);\n", i, i, (i%maxusers)+1, (i%maxstatus)+1)
		f.WriteString(sql)
	}
}
