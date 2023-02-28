package plant

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Plant, error)
	FindByID(ID int) (Plant, error)
	Create(plant Plant) (Plant, error)
	Update(plant Plant) (Plant, error)
	Delete(plant Plant) (Plant, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Plant, error) {
	var plants []Plant

	err := r.db.Find(&plants).Error

	return plants, err
}

func (r *repository) FindByID(ID int) (Plant, error) {
	var plant Plant

	err := r.db.Find(&plant, ID).Error

	return plant, err
}

func (r *repository) Create(plant Plant) (Plant, error) {
	err := r.db.Create(&plant).Error

	return plant, err
}

func (r *repository) Update(plant Plant) (Plant, error) {
	err := r.db.Save(&plant).Error

	return plant, err
}

func (r *repository) Delete(plant Plant) (Plant, error) {
	err := r.db.Delete(&plant).Error

	return plant, err
}