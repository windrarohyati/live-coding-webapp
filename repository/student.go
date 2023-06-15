package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type StudentRepository interface {
	FetchByID(id int) (*model.Student, error)
	Store(student *model.Student) error
	Delete(id int) error
}

type studentRepository struct {
	db *gorm.DB
}

func NewStudentRepo(db *gorm.DB) *studentRepository {
	return &studentRepository{db}
}

func (s *studentRepository) FetchByID(id int) (*model.Student, error) {
	var student model.Student
	err := s.db.Where("id = ?", id).First(&student).Error
	if err != nil {
		return nil, err
	}

	return &student, nil
}

func (s *studentRepository) Store(student *model.Student) error {
	err := s.db.Create(student).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *studentRepository) Delete(id int) error {
	student := model.Student{}
	if result := s.db.Where("id = ?", id).Delete(&student); result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}
