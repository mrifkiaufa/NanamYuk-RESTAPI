package userplants

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]UserPlants, error)
	FindByID(ID int) (UserPlants, error)
	Create(userPlants UserPlants) (UserPlants, error)
	Update(userPlants UserPlants) (UserPlants, error)
	Delete(userPlants UserPlants) (UserPlants, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]UserPlants, error) {
	var userPlants []UserPlants

	err := r.db.Find(&userPlants).Error

	return userPlants, err
}

func (r *repository) FindByID(ID int) (UserPlants, error) {
	var userPlants UserPlants

	err := r.db.Find(&userPlants, ID).Error

	return userPlants, err
}

func (r *repository) Create(userPlants UserPlants) (UserPlants, error) {
	err := r.db.Create(&userPlants).Error

	return userPlants, err
}

func (r *repository) Update(userPlants UserPlants) (UserPlants, error) {
	err := r.db.Save(&userPlants).Error

	return userPlants, err
}

func (r *repository) Delete(userPlants UserPlants) (UserPlants, error) {
	err := r.db.Delete(&userPlants).Error

	return userPlants, err
}