package main

import (
	"log"
	"os"

	"github.com/HaAnhTu2/ecommerce_web.git/db"
	"github.com/HaAnhTu2/ecommerce_web.git/route"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	client := db.Connect()
	db := client.Database(os.Getenv("DB_NAME"))
	port := os.Getenv("PORT")
	r := gin.Default()
	route.Route(r, db)
	r.Run(":" + port)
}
