package repository

import (
	"a21hc3NpZ25tZW50/model"
	"fmt"

	"gorm.io/gorm"
)

type StudentRepository interface {
	FetchAll() ([]model.Student, error)
	FetchByID(id int) (*model.Student, error)
	Store(s *model.Student) error
	Update(id int, s *model.Student) error
	Delete(id int) error
	FetchWithClass() (*[]model.StudentClass, error)
}

type studentRepoImpl struct {
	db *gorm.DB
}

func NewStudentRepo(db *gorm.DB) *studentRepoImpl {
	return &studentRepoImpl{db: db}
}

func (s *studentRepoImpl) FetchAll() ([]model.Student, error) {
	var students []model.Student
	s.db.Limit(10).Find(&[]model.Student{}).Scan(&students)
	return students, nil // TODO: replace this
}

func (s *studentRepoImpl) Store(student *model.Student) error {
	err := s.db.Create(&student)
	if err.Error != nil {
		return err.Error
	}
	return nil // TODO: replace this
}

func (s *studentRepoImpl) Update(id int, student *model.Student) error {
	err := s.db.Model(&model.Student{}).Where("id = ?", id).Updates(&model.Student{
		Name:    student.Name,
		Address: student.Address,
		ClassId: student.ClassId,
	})
	if err.Error != nil {
		return err.Error
	}
	return nil // TODO: replace this
}

func (s *studentRepoImpl) Delete(id int) error {
	err := s.db.Where("id = ?", id).Delete(&model.Student{})
	if err.Error != nil {
		return err.Error
	}
	return nil // TODO: replace this
}

func (s *studentRepoImpl) FetchByID(id int) (*model.Student, error) {
	var student model.Student
	err := s.db.Where("id = ?", id).First(&model.Student{}).Scan(&student)
	if err.Error != nil {
		return nil, err.Error
	}
	return &student, nil // TODO: replace this
}

func (s *studentRepoImpl) FetchWithClass() (*[]model.StudentClass, error) {
	var studentClass []model.StudentClass
	err := s.db.Limit(10).Table("students").Select("students.id, students.name, students.address, classes.name as class_name,classes.professor,classes.room_number").Joins("left join classes on students.class_id = classes.id").Scan(&studentClass)
	fmt.Printf("TOTAL ROWS AFFECTED ==> %v", err.RowsAffected)
	if err.RowsAffected == 0 {
		return &[]model.StudentClass{}, err.Error
	}
	return &studentClass, nil // TODO: replace this
}
