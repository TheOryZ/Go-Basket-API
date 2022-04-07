package order

import "gorm.io/gorm"

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}
func (r *OrderRepository) Migration() {
	r.db.AutoMigrate(&Order{})
}

//Create a new order
func (r *OrderRepository) Create(order *Order) error {
	return r.db.Create(order).Error
}

//Update a order
func (r *OrderRepository) Update(order *Order) error {
	return r.db.Save(order).Error
}

//Delete a order
func (r *OrderRepository) Delete(order *Order) error {
	return r.db.Delete(order).Error
}

//Delete a order by id
func (r *OrderRepository) DeleteByID(id uint) error {
	order := Order{}
	order.ID = id
	return r.db.Delete(&order).Error
}

//Find all orders
func (r *OrderRepository) FindAll() ([]Order, error) {
	var orders []Order
	err := r.db.Find(&orders).Error
	return orders, err
}

//Find a order by id
func (r *OrderRepository) FindByID(id uint) (*Order, error) {
	order := Order{}
	err := r.db.First(&order, id).Error
	return &order, err
}

//Find orders by user id
func (r *OrderRepository) FindByUserID(userID uint) ([]Order, error) {
	var orders []Order
	err := r.db.Where("user_id = ? AND deleted_at = ?", userID, nil).Find(&orders).Error
	return orders, err
}

//Find orders by user id with status //TODO: Chack this
func (r *OrderRepository) FindByUserIDInProgress(userID, statusID uint) ([]Order, error) {
	var orders []Order
	err := r.db.Where("user_id = ? AND status_id = ? AND deleted_at = ?", userID, statusID, nil).Find(&orders).Error
	return orders, err
}
