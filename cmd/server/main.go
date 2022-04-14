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
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/handlers"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/services"
	"github.com/gin-gonic/gin"
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
		users, _ := userRepo.FindAll()
		roles, _ := roleRepo.FindAll()
		userrolemapRepo := userrolemap.NewUserRoleMapRepository(db)
		userrolemapRepo.Seed(users, roles)
		log.Println("Migrations done")
	}
	log.Println("DB connected")

	//gin server
	router := gin.Default()
	router.Use(gin.Logger())

	//Services
	roleRepo := role.NewRoleRepository(db)
	userRepo := user.NewUserRepository(db)
	statusRepo := status.NewStatusRepository(db)
	categoryRepo := category.NewCategoryRepository(db)
	jwtService := services.NewJWTService()
	authService := services.NewAuthService(userRepo)
	roleService := services.NewRoleService(roleRepo)
	userService := services.NewUserService(userRepo)
	statusService := services.NewStatusService(statusRepo)
	categoryService := services.NewCategoryService(categoryRepo)
	//Handlers
	authHandler := handlers.NewAuthHandler(authService, jwtService)
	roleHandler := handlers.NewRoleHandler(roleService)
	userHandler := handlers.NewUserHandler(userService)
	statusHandler := handlers.NewStatusHandler(statusService)
	categoryHandler := handlers.NewCategoryHandler(categoryService)
	authRoutes := router.Group("api/auth")
	{
		authRoutes.POST("/login", authHandler.Login)
		authRoutes.POST("/register", authHandler.Register)
	}
	roleRoutes := router.Group("api/roles")
	{
		roleRoutes.GET("/", roleHandler.GetAllRoles)
		roleRoutes.GET("/:id", roleHandler.GetRole)
		roleRoutes.POST("/", roleHandler.CreateRole)
		roleRoutes.PUT("/", roleHandler.UpdateRole)
		roleRoutes.DELETE("/:id", roleHandler.DeleteRole)
	}
	userRoutes := router.Group("api/users")
	{
		userRoutes.GET("/", userHandler.GetAllUsers)
		userRoutes.GET("/:id", userHandler.GetUser)
		userRoutes.POST("/", userHandler.CreateUser)
		userRoutes.PUT("/", userHandler.UpdateUser)
		userRoutes.DELETE("/:id", userHandler.DeleteUser)
	}
	statusRoutes := router.Group("api/status")
	{
		statusRoutes.GET("/", statusHandler.GetAllStatus)
		statusRoutes.GET("/:id", statusHandler.GetStatus)
		statusRoutes.POST("/", statusHandler.CreateStatus)
		statusRoutes.PUT("/", statusHandler.UpdateStatus)
		statusRoutes.DELETE("/:id", statusHandler.DeleteStatus)
	}
	categoryRoutes := router.Group("api/categories")
	{
		categoryRoutes.GET("/", categoryHandler.GetAllCategories)
		categoryRoutes.GET("/:id", categoryHandler.GetCategory)
		categoryRoutes.POST("/", categoryHandler.CreateCategory)
		categoryRoutes.PUT("/", categoryHandler.UpdateCategory)
		categoryRoutes.DELETE("/:id", categoryHandler.DeleteCategory)
	}

	router.Run(":8080")

}
