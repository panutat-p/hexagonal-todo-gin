package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"os/user"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
}

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
	fmt.Println("Connected to ElephantSQL")

	err = db.AutoMigrate(&user.User{})
	if err != nil {
		panic("Failed to migrate ElephantSQL")
	}

	//todoHandler := todo.NewHandler(todo.NewGormStore(db))

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "healthy",
		})
	})

	//r.POST("/todo/new", todoHandler.NewTask)

	err = r.Run(fmt.Sprintf(":%s", os.Getenv("PORT"))) // block
	if err != nil {
		log.Println("🟥 Cannot start web server")
		return
	}
}
