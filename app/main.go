package main

import (
	"PBI/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDatabase()
	database.AutoMigrate()

	r := gin.Default()

	r.Run(":8080")
}
