package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type RequestBody struct {
	Id   *int   `json:"id,omitempty"`
	Name string `json:"name"`
}

func dbConnection() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn
}

func api(conn *pgx.Conn) {
	r := gin.Default() // gin router

	r.GET("/api/users", func(c *gin.Context) {
		id := c.Query("id")
		if id == "" {
			c.String(http.StatusBadRequest, "No id provided")
			return
		}
		var resultId int64
		var resultName string

		err := conn.QueryRow(context.Background(), "SELECT * FROM users WHERE id=$1", id).Scan(&resultId, &resultName)

		if err != nil {
			fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)

			c.String(http.StatusInternalServerError, "Something failed")
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"id":   resultId,
			"name": resultName,
		})
	})
	r.POST("/api/users", func(c *gin.Context) {
		var resultId int64
		var resultName string
		var requestBody RequestBody

		err := json.NewDecoder(c.Request.Body).Decode(&requestBody)
		if err != nil {
			fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
			c.String(http.StatusInternalServerError, "Something failed")
			return
		}

		fmt.Println(requestBody.Name)

		err2 := conn.QueryRow(context.Background(), "INSERT INTO users(name) VALUES ($1) RETURNING *", requestBody.Name).Scan(&resultId, &resultName)

		if err2 != nil {
			fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err2)
			c.String(http.StatusInternalServerError, "Something failed")
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"id":   resultId,
			"name": resultName,
		})
	})
	r.PUT("/api/users", func(c *gin.Context) {
		var resultId int64
		var resultName string
		var requestBody RequestBody

		err := json.NewDecoder(c.Request.Body).Decode(&requestBody)
		if err != nil {
			fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
			c.String(http.StatusInternalServerError, "Something failed")
			return
		}

		fmt.Println(requestBody.Name)

		err2 := conn.QueryRow(context.Background(), "INSERT INTO users(id, name) VALUES ($1, $2) ON CONFLICT (id) DO UPDATE SET name = $2 RETURNING *", requestBody.Id, requestBody.Name).Scan(&resultId, &resultName)

		if err2 != nil {
			fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err2)
			c.String(http.StatusInternalServerError, "Something failed")
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"id":   resultId,
			"name": resultName,
		})
	})
	r.DELETE("/api/users", func(c *gin.Context) {
		id := c.Query("id")
		if id == "" {
			c.String(http.StatusBadRequest, "No id provided")
			return
		}

		var deletedId string
		err := conn.QueryRow(context.Background(), "DELETE FROM users WHERE id=$1;", id).Scan(&deletedId)

		if err != nil || deletedId == "" {
			fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
			c.String(http.StatusInternalServerError, "Something failed")
			return
		}

		// TODO: should be conditional https://stackoverflow.com/a/2342589
		c.String(http.StatusOK, deletedId)
	})

	r.Run()
}

func main() {
	conn := dbConnection()
	defer conn.Close(context.Background())

	api(conn)
}

// go
// OOP in go
// go ORM
// services
// DTO gin
// queue
