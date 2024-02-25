package main

import (
	"example/hello/src/models"
	"example/hello/src/routes"
	"example/hello/src/utils"
	"log"
)

func main() {
 utils.LoadEnv()
 models.SetupDataBase()
 models.AutoMigrateModels()
 app := routes.SetupRoutes()

 err := app.Listen(":3003")
 if err != nil {
	 log.Fatalf("Failed to start server: %v", err)
 }
}