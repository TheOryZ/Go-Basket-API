package services

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/internal/store/domain/order"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/dtos"
	"github.com/gofrs/uuid"
)

//OrderService is an interface for OrderService
type OrderService interface {
	Create(model dtos.OrderCreateDTO) error
	Update(model dtos.OrderUpdateDTO) error
	Delete(model dtos.OrderUpdateDTO) error
	DeleteByID(id uuid.UUID) error
	FindAll() ([]dtos.OrderListDTO, error)
	FindByID(id uuid.UUID) (dtos.OrderListDTO, error)
	FindByUserID(userid uuid.UUID) ([]dtos.OrderListDTO, error)
	FindByUserIDInProgress(userid, statusid uuid.UUID) ([]dtos.OrderListDTO, error)
}

//orderService is a struct for OrderService
type orderService struct {
	orderRepository order.IOrderRepository
}

//NewOrderService is a constructor for OrderService
func NewOrderService(orderRepository order.IOrderRepository) OrderService {
	return &orderService{orderRepository: orderRepository}
}

//Create a new order
func (r *orderService) Create(model dtos.OrderCreateDTO) error {
	orderEntity := order.Order{}
	orderEntity.ID = uuid.Must(uuid.NewV4())
	orderEntity.UserID = model.UserID
	orderEntity.StatusID = model.StatusID
	err := r.orderRepository.Create(&orderEntity)
	if err != nil {
		return err
	}
	return nil
}

//Update a order
func (r *orderService) Update(model dtos.OrderUpdateDTO) error {
	orderEntity := order.Order{}
	orderEntity.ID = model.ID
	orderEntity.UserID = model.UserID
	orderEntity.StatusID = model.StatusID
	err := r.orderRepository.Update(&orderEntity)
	if err != nil {
		return err
	}
	return nil
}

//Delete a order
func (r *orderService) Delete(model dtos.OrderUpdateDTO) error {
	orderEntity := order.Order{}
	orderEntity.ID = model.ID
	orderEntity.UserID = model.UserID
	orderEntity.StatusID = model.StatusID
	err := r.orderRepository.Delete(&orderEntity)
	if err != nil {
		return err
	}
	return nil
}

//Delete a order by id
func (r *orderService) DeleteByID(id uuid.UUID) error {
	orderEntity := order.Order{}
	orderEntity.ID = id
	err := r.orderRepository.Delete(&orderEntity)
	if err != nil {
		return err
	}
	return nil
}

//FindAll orders
func (r *orderService) FindAll() ([]dtos.OrderListDTO, error) {
	orders, err := r.orderRepository.FindAll()
	if err != nil {
		return nil, err
	}
	var orderList []dtos.OrderListDTO
	for _, order := range orders {
		orderList = append(orderList, dtos.OrderListDTO{
			ID:       order.ID,
			User:     dtos.UserListDTO{ID: order.UserID},
			Product:  dtos.ProductListDTO{ID: order.ProductID},
			Quantity: order.Quantity,
			Price:    order.Price,
			Status:   dtos.StatusListDTO{ID: order.StatusID},
		})
	}
	return orderList, nil
}

//FindByID orders
func (r *orderService) FindByID(id uuid.UUID) (dtos.OrderListDTO, error) {
	orderEntity, err := r.orderRepository.FindByID(id)
	if err != nil {
		return dtos.OrderListDTO{}, err
	}
	return dtos.OrderListDTO{
		ID:       orderEntity.ID,
		User:     dtos.UserListDTO{ID: orderEntity.UserID},
		Product:  dtos.ProductListDTO{ID: orderEntity.ProductID},
		Quantity: orderEntity.Quantity,
		Price:    orderEntity.Price,
		Status:   dtos.StatusListDTO{ID: orderEntity.StatusID},
	}, nil
}

//FindByUserID orders
func (r *orderService) FindByUserID(userid uuid.UUID) ([]dtos.OrderListDTO, error) {
	orders, err := r.orderRepository.FindByUserID(userid)
	if err != nil {
		return nil, err
	}
	var orderList []dtos.OrderListDTO
	for _, order := range orders {
		orderList = append(orderList, dtos.OrderListDTO{
			ID:       order.ID,
			User:     dtos.UserListDTO{ID: order.UserID},
			Product:  dtos.ProductListDTO{ID: order.ProductID},
			Quantity: order.Quantity,
			Price:    order.Price,
			Status:   dtos.StatusListDTO{ID: order.StatusID},
		})
	}
	return orderList, nil
}

//FindByUserIDInProgress orders
func (r *orderService) FindByUserIDInProgress(userid, statusid uuid.UUID) ([]dtos.OrderListDTO, error) {
	orders, err := r.orderRepository.FindByUserIDInProgress(userid, statusid)
	if err != nil {
		return nil, err
	}
	var orderList []dtos.OrderListDTO
	for _, order := range orders {
		orderList = append(orderList, dtos.OrderListDTO{
			ID:       order.ID,
			User:     dtos.UserListDTO{ID: order.UserID},
			Product:  dtos.ProductListDTO{ID: order.ProductID},
			Quantity: order.Quantity,
			Price:    order.Price,
			Status:   dtos.StatusListDTO{ID: order.StatusID},
		})
	}
	return orderList, nil
}
