package routes

import (
	"example/hello/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func startupsGroupRouter(baseRouter fiber.Router) {
	startups := baseRouter.Group("/tasks")

	startups.Get("/all", controllers.GetAllTasks)
	startups.Get("/task/:id", controllers.GetTaskByIDHandler)
	startups.Post("/create", controllers.CreateTask)
	startups.Patch("/task/:id", controllers.UpdateTaskById)
	startups.Put("/update-status/:id", controllers.ChangeTaskStatus)
	startups.Delete("/delete/:id", controllers.DeleteTaskById)
}

func SetupRoutes() *fiber.App {
	app := fiber.New()

	versionRouter := app.Group("/api/v1")
	startupsGroupRouter(versionRouter)

	return app
}
