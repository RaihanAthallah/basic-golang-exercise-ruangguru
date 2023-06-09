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
	if result := t.db.Create(&data); result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}

func (t TeacherRepo) Query() ([]model.Teacher, error) {
	rows, err := t.db.Limit(10).Find(&model.Teacher{}).Rows()
	if err != nil {
		return nil, err
	}
	var teachers []model.Teacher
	for rows.Next() {
		t.db.ScanRows(rows, &teachers)
	}
	return teachers, nil // TODO: replace this
}

func (t TeacherRepo) Update(id uint, name string) error {
	if result := t.db.Table("teachers").Where("id = ?", id).Update("name", name); result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}

func (t TeacherRepo) Delete(id uint) error {
	teacher := model.Teacher{}
	result := t.db.Where("id = ?", id).Delete(&teacher)
	if result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}
