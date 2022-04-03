package main

import (
	"log"

	postgres "github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/store/common/db"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/store/domain/cart"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/store/domain/category"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/store/domain/order"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/store/domain/product"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/store/domain/productcategorymap"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/store/domain/role"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/store/domain/status"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/store/domain/user"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/store/domain/userrolemap"
	"github.com/joho/godotenv"
)

func main() {
	//Set enviroment variables
	err := godotenv.Load("../../env/settings.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := postgres.NewPsqlDB()
	if err != nil {
		log.Fatal("Postgres cannot init", err)
	}
	log.Println("Postgres connected")

	//Connection DB and migrations
	roleRepository := role.NewRoleRepository(db)
	roleRepository.Migration()
	userRepository := user.NewUserRepository(db)
	userRepository.Migration()
	categoryRepository := category.NewCategoryRepository(db)
	categoryRepository.Migration()
	productRepository := product.NewProductRepository(db)
	productRepository.Migration()
	statusRepository := status.NewStatusRepository(db)
	statusRepository.Migration()
	cartRepository := cart.NewCartRepository(db)
	cartRepository.Migration()
	orderRepository := order.NewOrderRepository(db)
	orderRepository.Migration()
	userRoleMapRepository := userrolemap.NewUserRoleMapRepository(db)
	userRoleMapRepository.Migration()
	productCategoryMapRepository := productcategorymap.NewProductCategoryMapRepository(db)
	productCategoryMapRepository.Migration()
	log.Println("Migrations done")
}
