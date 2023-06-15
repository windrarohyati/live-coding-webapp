package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type CourseRepository interface {
	FetchByID(id int) (*model.Course, error)
	Store(course *model.Course) error
	Delete(id int) error
}

type courseRepository struct {
	db *gorm.DB
}

func NewCourseRepo(db *gorm.DB) *courseRepository {
	return &courseRepository{db}
}

func (s *courseRepository) FetchByID(id int) (*model.Course, error) {
	var course model.Course
	err := s.db.Where("id = ?", id).First(&course).Error
	if err != nil {
		return nil, err
	}

	return &course, nil
}

func (c *courseRepository) Store(course *model.Course) error {
	err := c.db.Create(course).Error
	if err != nil {
		return err
	}

	return nil
}

func (c *courseRepository) Delete(id int) error {
	course := model.Course{}
	if result := c.db.Where("id = ?", id).Delete(&course); result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}
