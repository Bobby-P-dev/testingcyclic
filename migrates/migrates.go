package main

import (
	"github.com/Bobby-P-dev/testingcyclic/database"
	"github.com/Bobby-P-dev/testingcyclic/initiallizers"
	"github.com/Bobby-P-dev/testingcyclic/models"
)

func init() {
	initiallizers.LoadEnvVar()
	database.ConnectToDB()
}

func main() {
	database.DB.AutoMigrate(&models.Product{})
}
