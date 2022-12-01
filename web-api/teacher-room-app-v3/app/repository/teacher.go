package handler

import (
	"a21hc3NpZ25tZW50/app/model"

	"gorm.io/gorm"
)

type TeacherRepo struct {
	db *gorm.DB
}

func NewTeacherRepo(db *gorm.DB) TeacherRepo {
	return TeacherRepo{db}
}

func (u *TeacherRepo) AddTeacher(teacher model.Teacher) error {
	if teacher.Name != "" && teacher.FieldOfStudy != "" && teacher.Age != 0 {
		result := u.db.Create(&teacher)
		return result.Error
	}

	return nil // TODO: replace this
}

func (u *TeacherRepo) ReadTeacher() ([]model.ViewTeacher, error) {
	results := []model.ViewTeacher{}
	data := []model.Teacher{}
	result := u.db.Raw("SELECT * FROM teachers").Scan(&data)
	if result.Error != nil {
		return results, result.Error
	}
	for _, x := range data {
		res := model.ViewTeacher{}
		res.Name = x.Name
		res.FieldOfStudy = x.FieldOfStudy
		res.Age = x.Age
		results = append(results, res)
	}
	return results, result.Error // TODO: replace this
}

func (u *TeacherRepo) UpdateName(id uint, name string) error {
	result := u.db.Model(&model.Teacher{}).Where("id = ?", id).Update("name", name)
	return result.Error // TODO: replace this
}

func (u *TeacherRepo) DeleteTeacher(id uint) error {
	result := u.db.Delete(&model.Teacher{}, id)
	return result.Error // TODO: replace this
}
