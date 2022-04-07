package status

import "gorm.io/gorm"

type StatusRepository struct {
	db *gorm.DB
}

func NewStatusRepository(db *gorm.DB) *StatusRepository {
	return &StatusRepository{db: db}
}
func (r *StatusRepository) Migration() {
	r.db.AutoMigrate(&Status{})
}

//Create a new status
func (r *StatusRepository) Create(status *Status) error {
	return r.db.Create(status).Error
}

//Update a status
func (r *StatusRepository) Update(status *Status) error {
	return r.db.Save(status).Error
}

//Delete a status
func (r *StatusRepository) Delete(status *Status) error {
	return r.db.Delete(status).Error
}

//Delete a status by id
func (r *StatusRepository) DeleteByID(id uint) error {
	status := Status{}
	status.ID = id
	return r.db.Delete(&status).Error
}

//Find all statuses
func (r *StatusRepository) FindAll() ([]Status, error) {
	var statuses []Status
	err := r.db.Find(&statuses).Error
	return statuses, err
}

//Find a status by id
func (r *StatusRepository) FindByID(id uint) (*Status, error) {
	status := Status{}
	err := r.db.First(&status, id).Error
	return &status, err
}

//Find a status by name
func (r *StatusRepository) FindByName(name string) (*Status, error) {
	status := Status{}
	err := r.db.Where("name = ?", name).First(&status).Error
	return &status, err
}

//Seed a status
func (r *StatusRepository) Seed() error {
	statuses := []Status{
		{Name: "Pending"},
		{Name: "In Progress"},
		{Name: "Completed"},
		{Name: "Canceled"},
	}
	for _, status := range statuses {
		err := r.db.FirstOrCreate(&status, status).Error
		if err != nil {
			return err
		}
	}
	return nil
}
