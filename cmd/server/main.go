package main

import (
	"log"

	postgres "github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/internal/store/common/db"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/internal/store/domain/cart"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/internal/store/domain/category"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/internal/store/domain/order"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/internal/store/domain/product"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/internal/store/domain/productcategorymap"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/internal/store/domain/role"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/internal/store/domain/status"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/internal/store/domain/user"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/internal/store/domain/userrolemap"
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
	if !db.Migrator().HasTable(&user.User{}) {
		db.AutoMigrate(&user.User{}, &role.Role{}, &userrolemap.UserRoleMap{}, &category.Category{}, &product.Product{}, &status.Status{}, &order.Order{}, &productcategorymap.ProductCategoryMap{}, &cart.Cart{})
		//Add Seed data
		roleRepo := role.NewRoleRepository(db)
		roleRepo.Seed()
		userRepo := user.NewUserRepository(db)
		userRepo.Seed()
		statusRepo := status.NewStatusRepository(db)
		statusRepo.Seed()
		// userrolemapRepo := userrolemap.NewUserRoleMapRepository(db)
		// userrolemapRepo.Seed()
		log.Println("Migrations done")
	}
	log.Println("DB connected")

}
