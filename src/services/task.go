package services

import (
	"example/hello/src/models"
)

type Task struct {
    models.Task 
}


func FetchAllTasks() (*[]models.Task, error) {
	var tasks []models.Task
	err := models.Database.Find(&tasks).Error
	if err != nil {
	 return &[]models.Task{}, err
	}
	return &tasks, nil
}


func FindById(id uint) (*models.Task, error) {
    var task models.Task
    err := models.Database.First(&task, id).Error
    if err != nil {
        return nil, err
    }
    return &task, nil
}


func (task *Task) Create() (*models.Task, error) {
	err := models.Database.Model(&task.Task).Create(&task.Task).Error

	if err != nil {
		return nil, err
	}

	return &task.Task, nil
}



func (task *Task) Update(id uint, updatedTask *models.Task) (*models.Task, error) {

    existingTask, err := FindById(id)
    if err != nil {
        return nil, err
    }

	if updatedTask.Title != "" {
    existingTask.Title = updatedTask.Title
	}
	if updatedTask.Description != "" {
    existingTask.Description = updatedTask.Description
	}

    

    if err := models.Database.Save(existingTask).Error; err != nil {
        return nil, err
    }
    
    return existingTask, nil
}

func (task *Task) Delete(id uint) (*models.Task, error) {
    deletedTask, err := FindById(id)
    if err != nil {
        return nil, err
    }

    err = models.Database.Delete(&task, id).Error
    if err != nil {
        return nil, err
    }

    return deletedTask, nil
}

func (task *Task) ChangeTaskStatus(id uint, changeTask *models.Task) (*models.Task, error) {

    existingTask, err := FindById(id)
    if err != nil {
        return nil, err
    }

    existingTask.Status = changeTask.Status

    if err := models.Database.Save(existingTask).Error; err != nil {
        return nil, err
    }
    
    return existingTask, nil
}
