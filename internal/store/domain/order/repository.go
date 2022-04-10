package order

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

//interface
type IOrderRepository interface {
	Migration()
	Create(order *Order) error
	Update(order *Order) error
	Delete(order *Order) error
	DeleteByID(id uuid.UUID) error
	FindAll() ([]Order, error)
	FindByID(id uuid.UUID) (*Order, error)
	FindByUserID(userID uuid.UUID) ([]Order, error)
	FindByUserIDInProgress(userID, statusID uuid.UUID) ([]Order, error)
}

var OrderRepository IOrderRepository = &orderRepository{}

func NewOrderRepository(db *gorm.DB) *orderRepository {
	return &orderRepository{db: db}
}
func (r *orderRepository) Migration() {
	r.db.AutoMigrate(&Order{})
}

//Create a new order
func (r *orderRepository) Create(order *Order) error {
	return r.db.Create(order).Error
}

//Update a order
func (r *orderRepository) Update(order *Order) error {
	return r.db.Save(order).Error
}

//Delete a order
func (r *orderRepository) Delete(order *Order) error {
	return r.db.Delete(order).Error
}

//Delete a order by id
func (r *orderRepository) DeleteByID(id uuid.UUID) error {
	order := Order{}
	order.ID = id
	return r.db.Delete(&order).Error
}

//Find all orders
func (r *orderRepository) FindAll() ([]Order, error) {
	var orders []Order
	err := r.db.Find(&orders).Error
	return orders, err
}

//Find a order by id
func (r *orderRepository) FindByID(id uuid.UUID) (*Order, error) {
	order := Order{}
	err := r.db.First(&order, id).Error
	return &order, err
}

//Find orders by user id
func (r *orderRepository) FindByUserID(userID uuid.UUID) ([]Order, error) {
	var orders []Order
	err := r.db.Where("user_id = ? AND deleted_at = ?", userID, nil).Find(&orders).Error
	return orders, err
}

//Find orders by user id with status //TODO: Chack this
func (r *orderRepository) FindByUserIDInProgress(userID, statusID uuid.UUID) ([]Order, error) {
	var orders []Order
	err := r.db.Where("user_id = ? AND status_id = ? AND deleted_at = ?", userID, statusID, nil).Find(&orders).Error
	return orders, err
}
