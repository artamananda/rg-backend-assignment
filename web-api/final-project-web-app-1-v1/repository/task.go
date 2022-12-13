package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"
	"errors"

	"gorm.io/gorm"
)

type TaskRepository interface {
	GetTasks(ctx context.Context, id int) ([]entity.Task, error)
	StoreTask(ctx context.Context, task *entity.Task) (taskId int, err error)
	GetTaskByID(ctx context.Context, id int) (entity.Task, error)
	GetTasksByCategoryID(ctx context.Context, catId int) ([]entity.Task, error)
	UpdateTask(ctx context.Context, task *entity.Task) error
	DeleteTask(ctx context.Context, id int) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db}
}

func (r *taskRepository) GetTasks(ctx context.Context, id int) ([]entity.Task, error) {
	results := []entity.Task{}
	result := r.db.WithContext(ctx).Where("user_id = ?", id).Find(&results)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return []entity.Task{}, nil
		}
		return []entity.Task{}, result.Error
	}
	return results, nil // TODO: replace this
}

func (r *taskRepository) StoreTask(ctx context.Context, task *entity.Task) (taskId int, err error) {
	result := r.db.WithContext(ctx).Create(&task)
	if result.Error != nil {
		return 0, result.Error
	}
	return task.ID, nil // TODO: replace this
}

func (r *taskRepository) GetTaskByID(ctx context.Context, id int) (entity.Task, error) {
	results := entity.Task{}
	result := r.db.WithContext(ctx).First(&results, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return entity.Task{}, nil
		}
		return entity.Task{}, result.Error
	}
	return results, nil // TODO: replace this
}

func (r *taskRepository) GetTasksByCategoryID(ctx context.Context, catId int) ([]entity.Task, error) {
	results := []entity.Task{}
	result := r.db.WithContext(ctx).Where("category_id = ?", catId).Find(&results)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return []entity.Task{}, nil
		}
		return []entity.Task{}, result.Error
	}
	return results, nil // TODO: replace this
}

func (r *taskRepository) UpdateTask(ctx context.Context, task *entity.Task) error {
	result := r.db.WithContext(ctx).Updates(&task)
	return result.Error // TODO: replace this
}

func (r *taskRepository) DeleteTask(ctx context.Context, id int) error {
	result := r.db.WithContext(ctx).Delete(&entity.Task{}, id)
	return result.Error // TODO: replace this
}
