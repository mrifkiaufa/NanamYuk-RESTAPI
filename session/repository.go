package session

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Session, error)
	FindByID(ID int) (Session, error)
	Create(session Session) (Session, error)
	Update(session Session) (Session, error)
	Delete(session Session) (Session, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Session, error) {
	var session []Session

	err := r.db.Find(&session).Error

	return session, err
}

func (r *repository) FindByID(ID int) (Session, error) {
	var session Session

	err := r.db.Find(&session, ID).Error

	return session, err
}

func (r *repository) Create(session Session) (Session, error) {
	err := r.db.Create(&session).Error

	return session, err
}

func (r *repository) Update(session Session) (Session, error) {
	err := r.db.Save(&session).Error

	return session, err
}

func (r *repository) Delete(session Session) (Session, error) {
	err := r.db.Delete(&session).Error

	return session, err
}