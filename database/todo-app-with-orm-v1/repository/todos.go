package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return TodoRepository{db}
}

func (u *TodoRepository) AddTodo(todo model.Todo) error {
	result := u.db.Create(&todo)
	return result.Error // TODO: replace this
}

func (u *TodoRepository) ReadTodo() ([]model.Todo, error) {
	results := []model.Todo{}
	result := u.db.Raw("SELECT * FROM todos WHERE deleted_at is NULL").Take(&results)
	return results, result.Error // TODO: replace this
}

func (u *TodoRepository) UpdateDone(id uint, status bool) error {
	result := u.db.Model(&model.Todo{}).Where("id = ?", id).Update("done", status)
	return result.Error // TODO: replace this
}

func (u *TodoRepository) DeleteTodo(id uint) error {
	result := u.db.Delete(&model.Todo{}, id)
	return result.Error // TODO: replace this
}
