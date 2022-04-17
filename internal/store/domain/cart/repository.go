package cart

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type cartRepository struct {
	db *gorm.DB
}

//interface
type ICartRepository interface {
	Migration()
	Create(cart *Cart) error
	Update(cart *Cart) error
	Delete(cart *Cart) error
	DeleteByID(id uuid.UUID) error
	FindAll() ([]Cart, error)
	FindByID(id uuid.UUID) (*Cart, error)
	FindByUserID(userID uuid.UUID) ([]Cart, error)
	FindByUserIDAndID(userID, id uuid.UUID) ([]Cart, error)
	FindByUserIDInProgress(userID, statusID uuid.UUID) ([]Cart, error)
}

var CartRepository ICartRepository = &cartRepository{}

func NewCartRepository(db *gorm.DB) *cartRepository {
	return &cartRepository{db: db}
}
func (r *cartRepository) Migration() {
	r.db.AutoMigrate(&Cart{})
}

//Create a new cart
func (r *cartRepository) Create(cart *Cart) error {
	return r.db.Create(cart).Error
}

//Update a cart
func (r *cartRepository) Update(cart *Cart) error {
	return r.db.Save(cart).Error
}

//Delete a cart
func (r *cartRepository) Delete(cart *Cart) error {
	return r.db.Delete(cart).Error
}

//Delete a cart by id
func (r *cartRepository) DeleteByID(id uuid.UUID) error {
	cart := Cart{}
	cart.ID = id
	return r.db.Delete(&cart).Error
}

//Find all carts
func (r *cartRepository) FindAll() ([]Cart, error) {
	var carts []Cart
	err := r.db.Find(&carts).Error
	return carts, err
}

//Find a cart by id
func (r *cartRepository) FindByID(id uuid.UUID) (*Cart, error) {
	cart := Cart{}
	err := r.db.First(&cart, id).Error
	return &cart, err
}

//Find carts by user id
func (r *cartRepository) FindByUserID(userID uuid.UUID) ([]Cart, error) {
	var carts []Cart
	err := r.db.Where("user_id = ? AND deleted_at is null", userID).Find(&carts).Error
	return carts, err
}

//Find carts by user id with status //TODO: Chack this
func (r *cartRepository) FindByUserIDInProgress(userID, statusID uuid.UUID) ([]Cart, error) {
	var carts []Cart
	err := r.db.Where("user_id = ? AND status_id = ? AND deleted_at is null", userID, statusID).Find(&carts).Error
	return carts, err
}

//FindByUserIDAndID find a cart by user id and id
func (r *cartRepository) FindByUserIDAndID(userID, id uuid.UUID) ([]Cart, error) {
	var carts []Cart
	err := r.db.Where("user_id = ? AND id = ? AND deleted_at is null", userID, id).Find(&carts).Error
	return carts, err
}
