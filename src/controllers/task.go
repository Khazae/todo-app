package controllers

import (
	"example/hello/src/services"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllTasks(c *fiber.Ctx) error {
	startups, err := services.FetchAllTasks()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Tasks fetched successfully", "status": "success", "data": startups})
}

func GetTaskByIDHandler(c *fiber.Ctx) error {
    taskID := c.Params("id")
    if taskID == "" {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": "failed", "message": "Task ID is required", "data": nil})
    }


    id, err := strconv.ParseUint(taskID, 10, 64)
    if err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": "failed", "message": "Invalid Task ID", "data": nil})
    }


    task, err := services.FindById(uint(id))
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"status": "failed", "message": "Failed to get task", "data": nil})
    }

    return c.Status(http.StatusOK).JSON(fiber.Map{"status": "success", "message": "Task retrieved successfully", "data": task})
}

func CreateTask(c *fiber.Ctx) error {
	var input services.Task
	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
            "status": "failed", 
            "message": err.Error(), 
            "data": nil,
        })
	}

	savedStartup, err := input.Create()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": "failed", "message": err.Error(), "data": nil})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"status": "success", "message": "Startup saved successfully", "data": savedStartup})
}


func UpdateTaskById(c *fiber.Ctx) error {
    var updatedTask services.Task

    taskID := c.Params("id")
    if taskID == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed", "message": "Task ID is required", "data": nil})
    }

    id, err := strconv.ParseUint(taskID, 10, 64)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed", "message": "Invalid Task ID", "data": nil})
    }


    if err := c.BodyParser(&updatedTask); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed", "message": "Invalid JSON data", "data": nil})
    }


    deletedTask, err := updatedTask.Update(uint(id), &updatedTask.Task)

    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed", "message": "Failed to update task", "data": nil})
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Task updated successfully", "data": deletedTask})
}

func DeleteTaskById(c *fiber.Ctx) error {
	var input services.Task
    taskID := c.Params("id")
    if taskID == "" {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": "failed", "message": "Task ID is required", "data": nil})
    }

 
    id, err := strconv.ParseUint(taskID, 10, 64)
    if err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": "failed", "message": "Invalid Task ID", "data": nil})
    }


    deletedTask, err := input.Delete(uint(id))
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"status": "failed", "message": "Failed to delete task", "data": nil})
    }

    return c.Status(http.StatusOK).JSON(fiber.Map{"status": "success", "message": "Task deleted successfully", "data": deletedTask})
}

func ChangeTaskStatus(c *fiber.Ctx) error {
    var updatedTask services.Task

    taskID := c.Params("id")
    if taskID == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed", "message": "Task ID is required", "data": nil})
    }

    id, err := strconv.ParseUint(taskID, 10, 64)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed", "message": "Invalid Task ID", "data": nil})
    }


    if err := c.BodyParser(&updatedTask); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed", "message": "Invalid JSON data", "data": nil})
    }


    changeTask, err := updatedTask.ChangeTaskStatus(uint(id), &updatedTask.Task)

    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed", "message": "Failed to update task", "data": nil})
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Task updated successfully", "data": changeTask})
}


