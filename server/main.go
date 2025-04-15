package main

import (
	"context"
	"fmt"
	"os"
	// "net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func dbConnection() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	return conn

}

func api(conn *pgx.Conn) {
	r := gin.Default()
	r.GET("/api", func(c *gin.Context) {
		var name string
		var weight int64

		err := conn.QueryRow(context.Background(), "select name, weight from widgets where id=$1", 42).Scan(&name, &weight)

		if err != nil {
			fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(name, weight)
	})
	r.POST("/api", func(c *gin.Context) {
	})
	r.PUT("/api", func(c *gin.Context) {
	})
	r.DELETE("/api", func(c *gin.Context) {
	})

	r.Run()
}

func main() {
	conn := dbConnection()
	api(conn)
}

// c.JSON(http.StatusOK, gin.H{
// 	"message": "pong",
// })
