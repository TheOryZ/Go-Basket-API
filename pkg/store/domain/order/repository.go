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
