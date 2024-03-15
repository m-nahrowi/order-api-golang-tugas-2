package main

import (
	"log"
	"fmt"
	"github.com/gin-gonic/gin"
	"main.go/controllers"
	"main.go/models"
	// "main.go/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "44342"
	dbname   = "order-api"
)

func main() {
	// Koneksi ke database
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Migrasi model ke database
	err = db.AutoMigrate(&models.Order{}, &models.Item{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Set DB di controller
	orderController := controllers.NewOrderController(db)

	// Inisialisasi Gin router
	router := gin.Default()

	// Routes
	router.POST("/orders", orderController.CreateOrder)
	router.GET("/orders/:orderID", orderController.GetOrder)
	router.PUT("/orders/:orderID", orderController.UpdateOrder)
	router.DELETE("/orders/:orderID", orderController.DeleteOrder)
	router.GET("/orders", orderController.GetAllOrder)

	// Jalankan server
	err = router.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
