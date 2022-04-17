package main

import (
	"log"
	"net/http"

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
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/middleware"
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

	//Repositories
	roleRepo := role.NewRoleRepository(db)
	userRepo := user.NewUserRepository(db)
	statusRepo := status.NewStatusRepository(db)
	categoryRepo := category.NewCategoryRepository(db)
	productRepo := product.NewProductRepository(db)
	cartRepo := cart.NewCartRepository(db)
	orderRepo := order.NewOrderRepository(db)
	productcategorymapRepo := productcategorymap.NewProductCategoryMapRepository(db)
	userrolemapRepo := userrolemap.NewUserRoleMapRepository(db)
	//Services
	jwtService := services.NewJWTService()
	authService := services.NewAuthService(userRepo)
	roleService := services.NewRoleService(roleRepo)
	userService := services.NewUserService(userRepo)
	statusService := services.NewStatusService(statusRepo)
	categoryService := services.NewCategoryService(categoryRepo)
	productService := services.NewProductService(productRepo)
	cartService := services.NewCartService(cartRepo)
	orderService := services.NewOrderService(orderRepo)
	productcategorymapService := services.NewProductCategoryMapService(productcategorymapRepo)
	userrolemapService := services.NewUserRoleMapService(userrolemapRepo)
	//Handlers
	authHandler := handlers.NewAuthHandler(authService, jwtService, roleService, userrolemapService)
	roleHandler := handlers.NewRoleHandler(roleService, userService, jwtService)
	userHandler := handlers.NewUserHandler(userService, roleService, jwtService)
	statusHandler := handlers.NewStatusHandler(statusService, jwtService, roleService)
	categoryHandler := handlers.NewCategoryHandler(categoryService, productService, roleService, jwtService)
	productHandler := handlers.NewProductHandler(productService, categoryService, roleService, jwtService)
	cartHandler := handlers.NewCartHandler(cartService, statusService, productService, roleService, jwtService, orderService)
	orderHandler := handlers.NewOrderHandler(orderService, cartService, productService, userService, statusService, roleService, jwtService)
	productcategorymapHandler := handlers.NewProductCategoryMapHandler(productcategorymapService, productService, categoryService, roleService, jwtService)

	router.StaticFS("/uploads", http.Dir("../../uploads"))
	authRoutes := router.Group("api/auth")
	{
		authRoutes.POST("/login", authHandler.Login)
		authRoutes.POST("/register", authHandler.Register)
	}
	roleRoutes := router.Group("api/roles", middleware.AuthorizeJWT(jwtService))
	{
		roleRoutes.GET("/", roleHandler.GetAllRoles)
		roleRoutes.GET("/:id", roleHandler.GetRole)
		roleRoutes.GET("/:id/users", roleHandler.GetRoleWithUsers)
		roleRoutes.POST("/", roleHandler.CreateRole)
		roleRoutes.PUT("/", roleHandler.UpdateRole)
		roleRoutes.DELETE("/:id", roleHandler.DeleteRole)
	}
	userRoutes := router.Group("api/users")
	{
		userRoutes.GET("/", userHandler.GetAllUsers, middleware.AuthorizeJWT(jwtService))
		userRoutes.GET("/:id", userHandler.GetUser)
		userRoutes.GET("/:id/roles", userHandler.GetUserWithRoles)
		userRoutes.POST("/", userHandler.CreateUser, middleware.AuthorizeJWT(jwtService))
		userRoutes.PUT("/", userHandler.UpdateUser, middleware.AuthorizeJWT(jwtService))
		userRoutes.DELETE("/:id", userHandler.DeleteUser, middleware.AuthorizeJWT(jwtService))
	}
	statusRoutes := router.Group("api/status")
	{
		statusRoutes.GET("/", statusHandler.GetAllStatus)
		statusRoutes.GET("/:id", statusHandler.GetStatus)
		statusRoutes.POST("/", statusHandler.CreateStatus, middleware.AuthorizeJWT(jwtService))
		statusRoutes.PUT("/", statusHandler.UpdateStatus, middleware.AuthorizeJWT(jwtService))
		statusRoutes.DELETE("/:id", statusHandler.DeleteStatus, middleware.AuthorizeJWT(jwtService))
	}
	categoryRoutes := router.Group("api/categories")
	{
		categoryRoutes.GET("/", categoryHandler.GetAllCategories)
		categoryRoutes.GET("/pagign", categoryHandler.GetAllCategoriesPaging)
		categoryRoutes.GET("/:id", categoryHandler.GetCategory)
		categoryRoutes.GET("/:id/products", categoryHandler.GetCategoryWithProducts)
		categoryRoutes.POST("/", categoryHandler.CreateCategory, middleware.AuthorizeJWT(jwtService))
		categoryRoutes.POST("/file", categoryHandler.UploadCsvFile, middleware.AuthorizeJWT(jwtService))
		categoryRoutes.PUT("/", categoryHandler.UpdateCategory, middleware.AuthorizeJWT(jwtService))
		categoryRoutes.DELETE("/:id", categoryHandler.DeleteCategory, middleware.AuthorizeJWT(jwtService))
	}
	productRoutes := router.Group("api/products")
	{
		productRoutes.GET("/", productHandler.GetAllProducts)
		productRoutes.GET("/pagign", productHandler.GetAllProductsPaging)
		productRoutes.GET("/:id", productHandler.GetProduct)
		productRoutes.GET("/:id/categories", productHandler.GetProductWithCategories)
		productRoutes.POST("/", productHandler.CreateProduct, middleware.AuthorizeJWT(jwtService))
		productRoutes.PUT("/", productHandler.UpdateProduct, middleware.AuthorizeJWT(jwtService))
		productRoutes.DELETE("/:id", productHandler.DeleteProduct, middleware.AuthorizeJWT(jwtService))
	}
	productCategoryMapRoutes := router.Group("api/productcategorymaps")
	{
		productCategoryMapRoutes.GET("/", productcategorymapHandler.GetAllProductCategoryMaps)
		productCategoryMapRoutes.GET("/:id", productcategorymapHandler.GetProductCategoryMap)
		productCategoryMapRoutes.POST("/", productcategorymapHandler.CreateProductCategoryMap, middleware.AuthorizeJWT(jwtService))
		productCategoryMapRoutes.PUT("/", productcategorymapHandler.UpdateProductCategoryMap, middleware.AuthorizeJWT(jwtService))
		productCategoryMapRoutes.DELETE("/:id", productcategorymapHandler.DeleteProductCategoryMap, middleware.AuthorizeJWT(jwtService))
	}
	cartRoutes := router.Group("api/carts", middleware.AuthorizeJWT(jwtService))
	{
		cartRoutes.GET("/", cartHandler.GetAllCarts)
		cartRoutes.GET("/:id", cartHandler.GetCart)
		cartRoutes.GET("/user/:id", cartHandler.GetCartsByUserID)
		cartRoutes.GET("/user/:id/in_progress", cartHandler.GetCartsByUserIDInProgress)
		cartRoutes.POST("/", cartHandler.CreateCart)
		cartRoutes.PUT("/", cartHandler.UpdateCart)
		cartRoutes.DELETE("/:id", cartHandler.DeleteCart)
	}
	orderRoutes := router.Group("api/orders", middleware.AuthorizeJWT(jwtService))
	{
		orderRoutes.GET("/", orderHandler.GetAllOrders)
		orderRoutes.GET("/:id", orderHandler.GetOrder)
		orderRoutes.GET("/user/:id", orderHandler.GetOrderByUser)
		orderRoutes.POST("/", orderHandler.CreateOrder)
		orderRoutes.PUT("/", orderHandler.UpdateOrder)
		orderRoutes.DELETE("/:id", orderHandler.DeleteOrder)
	}

	router.Run(":8080")

}
