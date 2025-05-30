package main

import (
	"database/sql"
	"log"
	"github.com/koushik/task-management-system/pkg/api"
	"github.com/koushik/task-management-system/pkg/repository"
	"github.com/koushik/task-management-system/pkg/service"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://user:password@localhost:5432/tasks_db?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewTaskRepository(db)
	svc := service.NewTaskService(repo)
	handler := api.NewTaskHandler(svc)

	r := gin.Default()

	r.POST("/tasks", handler.CreateTask)
	r.DELETE("/tasks/:id", handler.DeleteTask)
	r.PUT("/tasks/:id", handler.UpdateTask)

	r.Run(":8080")
}
