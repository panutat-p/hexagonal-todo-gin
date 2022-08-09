package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/panutat-p/hexagonal-todo-gin/core/domain"
	"github.com/panutat-p/hexagonal-todo-gin/core/service"
	"github.com/panutat-p/hexagonal-todo-gin/repository"
	"github.com/panutat-p/hexagonal-todo-gin/router"
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

	err = db.AutoMigrate(&domain.Todo{})
	if err != nil {
		panic("Failed to migrate ElephantSQL")
	}

	todoHandler := service.NewTodoHandler(repository.NewGormStore(db))

	r := router.NewGinRouter() // decoupling technique -> wrap router

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
		})
	})

	r.POST("/todo/new", todoHandler.CreateNewTask) // we can use TodoHandler directly

	err = r.Run(fmt.Sprintf(":%s", os.Getenv("PORT"))) // block
	if err != nil {
		log.Println("🟥 Cannot start web server")
		return
	}
}
