// app/main.go

package main

import (
	"PBI/database" // Sesuaikan dengan nama proyek Anda

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDatabase()
	database.AutoMigrate()

	r := gin.Default()

	// Routes and other logic here...

	r.Run(":8080")
}
