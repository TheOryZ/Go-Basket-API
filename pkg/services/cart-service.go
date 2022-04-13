package services

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/internal/store/domain/cart"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/dtos"
	"github.com/gofrs/uuid"
)

//CartService is an interface for CartService
type CartService interface {
	Create(model dtos.CartCreateDTO) error
	Update(model dtos.CartUpdateDTO) error
	Delete(model dtos.CartUpdateDTO) error
	DeleteByID(id uuid.UUID) error
	FindAll() ([]dtos.CartListDTO, error)
	FindByID(id uuid.UUID) (dtos.CartListDTO, error)
	FindByUserID(userid uuid.UUID) ([]dtos.CartListDTO, error)
	FindByUserIDInProgress(userid, statusid uuid.UUID) ([]dtos.CartListDTO, error)
}

//cartService is a struct for CartService
type cartService struct {
	cartRepository cart.ICartRepository
}

//NewCartService is a constructor for CartService
func NewCartService(cartRepository cart.ICartRepository) CartService {
	return &cartService{cartRepository: cartRepository}
}

//Create a new cart
func (r *cartService) Create(model dtos.CartCreateDTO) error {
	cartEntity := cart.Cart{}
	cartEntity.ID = uuid.Must(uuid.NewV4())
	cartEntity.UserID = model.UserID
	cartEntity.StatusID = model.StatusID
	err := r.cartRepository.Create(&cartEntity)
	if err != nil {
		return err
	}
	return nil
}

//Update a cart
func (r *cartService) Update(model dtos.CartUpdateDTO) error {
	cartEntity := cart.Cart{}
	cartEntity.ID = model.ID
	cartEntity.UserID = model.UserID
	cartEntity.StatusID = model.StatusID
	err := r.cartRepository.Update(&cartEntity)
	if err != nil {
		return err
	}
	return nil
}

//Delete a cart
func (r *cartService) Delete(model dtos.CartUpdateDTO) error {
	cartEntity := cart.Cart{}
	cartEntity.ID = model.ID
	cartEntity.UserID = model.UserID
	cartEntity.StatusID = model.StatusID
	err := r.cartRepository.Delete(&cartEntity)
	if err != nil {
		return err
	}
	return nil
}

//Delete a cart by id
func (r *cartService) DeleteByID(id uuid.UUID) error {
	err := r.cartRepository.DeleteByID(id)
	if err != nil {
		return err
	}
	return nil
}

//Find all carts
func (r *cartService) FindAll() ([]dtos.CartListDTO, error) {
	carts, err := r.cartRepository.FindAll()
	if err != nil {
		return nil, err
	}
	var cartList []dtos.CartListDTO
	for _, cart := range carts {

		cartList = append(cartList, dtos.CartListDTO{
			ID:       cart.ID,
			User:     dtos.UserListDTO{ID: cart.UserID, Name: cart.User.Name, Email: cart.User.Email},
			Product:  dtos.ProductListDTO{ID: cart.ProductID, Name: cart.Product.Name, Price: cart.Product.Price},
			Quantity: cart.Quantity,
			Price:    cart.Price,
			Status:   dtos.StatusListDTO{ID: cart.StatusID, Name: cart.Status.Name},
		})
	}
	return cartList, nil
}

//Find a cart by id
func (r *cartService) FindByID(id uuid.UUID) (dtos.CartListDTO, error) {
	cart, err := r.cartRepository.FindByID(id)
	if err != nil {
		return dtos.CartListDTO{}, err
	}
	return dtos.CartListDTO{
		ID:       cart.ID,
		User:     dtos.UserListDTO{ID: cart.UserID, Name: cart.User.Name, Email: cart.User.Email},
		Product:  dtos.ProductListDTO{ID: cart.ProductID, Name: cart.Product.Name, Price: cart.Product.Price},
		Quantity: cart.Quantity,
		Price:    cart.Price,
		Status:   dtos.StatusListDTO{ID: cart.StatusID, Name: cart.Status.Name},
	}, nil
}

//Find a cart by user id
func (r *cartService) FindByUserID(userid uuid.UUID) ([]dtos.CartListDTO, error) {
	carts, err := r.cartRepository.FindByUserID(userid)
	if err != nil {
		return nil, err
	}
	var cartList []dtos.CartListDTO
	for _, cart := range carts {

		cartList = append(cartList, dtos.CartListDTO{
			ID:       cart.ID,
			User:     dtos.UserListDTO{ID: cart.UserID, Name: cart.User.Name, Email: cart.User.Email},
			Product:  dtos.ProductListDTO{ID: cart.ProductID, Name: cart.Product.Name, Price: cart.Product.Price},
			Quantity: cart.Quantity,
			Price:    cart.Price,
			Status:   dtos.StatusListDTO{ID: cart.StatusID, Name: cart.Status.Name},
		})
	}
	return cartList, nil
}

//Find a cart by user id and status id
func (r *cartService) FindByUserIDInProgress(userid, statusid uuid.UUID) ([]dtos.CartListDTO, error) {
	carts, err := r.cartRepository.FindByUserIDInProgress(userid, statusid)
	if err != nil {
		return nil, err
	}
	var cartList []dtos.CartListDTO
	for _, cart := range carts {

		cartList = append(cartList, dtos.CartListDTO{
			ID:       cart.ID,
			User:     dtos.UserListDTO{ID: cart.UserID, Name: cart.User.Name, Email: cart.User.Email},
			Product:  dtos.ProductListDTO{ID: cart.ProductID, Name: cart.Product.Name, Price: cart.Product.Price},
			Quantity: cart.Quantity,
			Price:    cart.Price,
			Status:   dtos.StatusListDTO{ID: cart.StatusID, Name: cart.Status.Name},
		})
	}
	return cartList, nil
}
