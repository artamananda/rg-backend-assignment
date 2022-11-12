package repo

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type TeacherRepo struct {
	db *gorm.DB
}

func NewTeacherRepo(db *gorm.DB) TeacherRepo {
	return TeacherRepo{db}
}

func (t TeacherRepo) Save(data model.Teacher) error {
	result := t.db.Create(&data)
	return result.Error // TODO: replace this
}

func (t TeacherRepo) Query() ([]model.Teacher, error) {
	teacher := []model.Teacher{}
	result := t.db.Raw("SELECT * FROM teachers").Take(&teacher)
	return teacher, result.Error // TODO: replace this
}

func (t TeacherRepo) Update(id uint, name string) error {
	teacher := model.Teacher{}
	result := t.db.Model(&teacher).Where("id = ?", id).Update("name", name)
	return result.Error // TODO: replace this
}

func (t TeacherRepo) Delete(id uint) error {
	teacher := model.Teacher{}
	result := t.db.Delete(&teacher, id)
	return result.Error // TODO: replace this
}
