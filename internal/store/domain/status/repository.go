package status

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type statusRepository struct {
	db *gorm.DB
}

//interface
type IStatusRepository interface {
	Migration()
	Create(status *Status) error
	Update(status *Status) error
	Delete(status *Status) error
	DeleteByID(id uuid.UUID) error
	FindAll() ([]Status, error)
	FindByID(id uuid.UUID) (*Status, error)
	FindByName(name string) (*Status, error)
	Seed() error
}

var StatusRepository IStatusRepository = &statusRepository{}

func NewStatusRepository(db *gorm.DB) *statusRepository {
	return &statusRepository{db: db}
}
func (r *statusRepository) Migration() {
	r.db.AutoMigrate(&Status{})
}

//Create a new status
func (r *statusRepository) Create(status *Status) error {
	return r.db.Create(status).Error
}

//Update a status
func (r *statusRepository) Update(status *Status) error {
	return r.db.Save(status).Error
}

//Delete a status
func (r *statusRepository) Delete(status *Status) error {
	return r.db.Delete(status).Error
}

//Delete a status by id
func (r *statusRepository) DeleteByID(id uuid.UUID) error {
	status := Status{}
	status.ID = id
	return r.db.Delete(&status).Error
}

//Find all statuses
func (r *statusRepository) FindAll() ([]Status, error) {
	var statuses []Status
	err := r.db.Find(&statuses).Error
	return statuses, err
}

//Find a status by id
func (r *statusRepository) FindByID(id uuid.UUID) (*Status, error) {
	status := Status{}
	err := r.db.First(&status, id).Error
	return &status, err
}

//Find a status by name
func (r *statusRepository) FindByName(name string) (*Status, error) {
	status := Status{}
	err := r.db.Where("name = ?", name).First(&status).Error
	return &status, err
}

//Seed a status
func (r *statusRepository) Seed() error {
	statuses := []Status{
		{Name: "Pending", CreatedAt: "2020-01-01 00:00:00", UpdatedAt: "2020-01-01 00:00:00"},
		{Name: "In Progress", CreatedAt: "2020-01-01 00:00:00", UpdatedAt: "2020-01-01 00:00:00"},
		{Name: "Completed", CreatedAt: "2020-01-01 00:00:00", UpdatedAt: "2020-01-01 00:00:00"},
		{Name: "Canceled", CreatedAt: "2020-01-01 00:00:00", UpdatedAt: "2020-01-01 00:00:00"},
	}
	for _, status := range statuses {
		err := r.db.FirstOrCreate(&status, status).Error
		if err != nil {
			return err
		}
	}
	return nil
}
