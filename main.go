package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/panutat-p/hexagonal-todo-gin/todo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load ENV")
	}
	dsn := os.Getenv("DATABASE_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect ElephantSQL")
	}
	fmt.Println("🟩 Connected to ElephantSQL")

	err = db.AutoMigrate(&todo.Todo{})
	if err != nil {
		panic("Failed to migrate ElephantSQL")
	}

	todoHandler := todo.NewHandler(todo.NewGormStore(db))

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
		})
	})

	r.POST("/todo/new", todo.NewGinHandler(todoHandler.CreateNewTask)) // convert AdapterHandler to GinHandler

	err = r.Run(fmt.Sprintf(":%s", os.Getenv("PORT"))) // block
	if err != nil {
		log.Println("🟥 Cannot start web server")
		return
	}
}
