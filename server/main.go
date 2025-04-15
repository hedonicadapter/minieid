package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func dbConnection() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn
}

func api(conn *pgx.Conn) {
	r := gin.Default()
	r.GET("/api", func(c *gin.Context) {
		var id int64
		var name string

		err := conn.QueryRow(context.Background(), "select id, name from users").Scan(&id, &name)

		if err != nil {
			fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)

			// TODO: why os.Exit()? send some error status?
			os.Exit(1)
		}

		c.JSON(http.StatusOK, gin.H{
			"id":   id,
			"name": name,
		})
	})
	r.POST("/api", func(c *gin.Context) {
	})
	r.PUT("/api", func(c *gin.Context) {
	})
	r.DELETE("/api", func(c *gin.Context) {
		id := c.Query("id")
		if id == nil {
			c.JSON(http.StatusBadRequest)
		}

		err := conn.QueryRow(context.Background(), "DELETE FROM users WHERE id=$1", id)

		if err != nil {
			fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)

			// TODO: why os.Exit()? send some error status?
			os.Exit(1)
		}

		// TODO: should be conditional https://stackoverflow.com/a/2342589
		c.JSON(http.StatusOK)
	})

	r.Run()
}

func main() {
	conn := dbConnection()
	api(conn)

	defer conn.Close(context.Background())
}
