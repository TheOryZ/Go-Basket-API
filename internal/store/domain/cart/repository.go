package cart

import "gorm.io/gorm"

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{db: db}
}
func (r *CartRepository) Migration() {
	r.db.AutoMigrate(&Cart{})
}

//Create a new cart
func (r *CartRepository) Create(cart *Cart) error {
	return r.db.Create(cart).Error
}

//Update a cart
func (r *CartRepository) Update(cart *Cart) error {
	return r.db.Save(cart).Error
}

//Delete a cart
func (r *CartRepository) Delete(cart *Cart) error {
	return r.db.Delete(cart).Error
}

//Delete a cart by id
func (r *CartRepository) DeleteByID(id uint) error {
	cart := Cart{}
	cart.ID = id
	return r.db.Delete(&cart).Error
}

//Find all carts
func (r *CartRepository) FindAll() ([]Cart, error) {
	var carts []Cart
	err := r.db.Find(&carts).Error
	return carts, err
}

//Find a cart by id
func (r *CartRepository) FindByID(id uint) (*Cart, error) {
	cart := Cart{}
	err := r.db.First(&cart, id).Error
	return &cart, err
}

//Find carts by user id
func (r *CartRepository) FindByUserID(userID uint) ([]Cart, error) {
	var carts []Cart
	err := r.db.Where("user_id = ? AND deleted_at = ?", userID, nil).Find(&carts).Error
	return carts, err
}

//Find carts by user id with status //TODO: Chack this
func (r *CartRepository) FindByUserIDInProgress(userID, statusID uint) ([]Cart, error) {
	var carts []Cart
	err := r.db.Where("user_id = ? AND status_id = ? AND deleted_at = ?", userID, statusID, nil).Find(&carts).Error
	return carts, err
}
