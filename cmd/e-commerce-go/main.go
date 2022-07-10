package main

import (
	"context"
	"fmt"
	_ "github.com/berkayersoyy/e-commerce-go/docs"
	authHandler "github.com/berkayersoyy/e-commerce-go/internal/application/handlers/auth"
	categoryHandler "github.com/berkayersoyy/e-commerce-go/internal/application/handlers/category"
	orderHandler "github.com/berkayersoyy/e-commerce-go/internal/application/handlers/order"
	productHandler "github.com/berkayersoyy/e-commerce-go/internal/application/handlers/product"
	userHandler "github.com/berkayersoyy/e-commerce-go/internal/application/handlers/user"
	authService "github.com/berkayersoyy/e-commerce-go/internal/application/services/auth"
	categoryService "github.com/berkayersoyy/e-commerce-go/internal/application/services/category"
	orderService "github.com/berkayersoyy/e-commerce-go/internal/application/services/order"
	orderItemService "github.com/berkayersoyy/e-commerce-go/internal/application/services/order_item"
	productService "github.com/berkayersoyy/e-commerce-go/internal/application/services/product"
	userService "github.com/berkayersoyy/e-commerce-go/internal/application/services/user"
	"github.com/berkayersoyy/e-commerce-go/internal/infrastructure/http/middlewares"
	orderItemRepository "github.com/berkayersoyy/e-commerce-go/internal/infrastructure/repositories/order_item"
	"github.com/berkayersoyy/e-commerce-go/internal/providers"
	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	categoryRepository "github.com/berkayersoyy/e-commerce-go/internal/infrastructure/repositories/category"
	orderRepository "github.com/berkayersoyy/e-commerce-go/internal/infrastructure/repositories/order"
	productRepository "github.com/berkayersoyy/e-commerce-go/internal/infrastructure/repositories/product"

	userRepository "github.com/berkayersoyy/e-commerce-go/internal/infrastructure/repositories/user"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sethvargo/go-retry"
	"log"
	"time"
)

//version app_version
var version = "dev"

func setup(ctx context.Context) *gin.Engine {
	//mysql
	mysqlDb := providers.NewMysqlDb()
	dynamoDbSession, err := providers.NewSession()
	if err != nil {
		log.Fatal(err)
	}
	//dynamodb
	dynamoDb := providers.NewDynamoDb(dynamoDbSession)

	//redis
	redisDb := providers.NewRedisDb()

	productRepo := productRepository.ProvideProductRepository(mysqlDb)
	productSvc := productService.ProvideProductService(productRepo)
	productApi := productHandler.ProvideProductHandler(productSvc)

	userRepo := userRepository.ProvideUserRepository(time.Second*30, dynamoDb)
	userSvc := userService.ProvideUserService(userRepo)
	userApi := userHandler.ProvideUserHandler(userSvc)

	categoryRepo := categoryRepository.ProvideCategoryRepository(mysqlDb)
	categorySvc := categoryService.ProvideCategoryService(categoryRepo)
	categoryApi := categoryHandler.ProvideCategoryHandler(categorySvc)

	orderItemRepo := orderItemRepository.ProvideOrderItemRepository(mysqlDb)
	orderItemSvc := orderItemService.ProvideOrderItemService(orderItemRepo)

	orderRepo := orderRepository.ProvideOrderRepository(mysqlDb)
	orderSvc := orderService.ProvideOrderService(orderRepo, productSvc, orderItemSvc)
	orderApi := orderHandler.ProvideOrderHandler(orderSvc)

	err = userRepo.CreateTable(ctx)
	if err != nil {
		log.Fatalf("Error on creating users table, %s", err)
	}

	authSvc := authService.ProvideAuthService(redisDb)
	authApi := authHandler.ProvideAuthHandler(authSvc, userSvc)

	router := gin.Default()

	//products
	products := router.Group("/v1/products")
	products.Use(middlewares.AuthorizeJWTMiddleware(authSvc))

	products.GET("", productApi.GetAllProducts)
	products.GET("/getbycategoryid/:id", productApi.GetProductsByCategoryID)
	products.POST("", productApi.CreateProduct)
	products.GET("/:id", productApi.GetProductByID)
	products.DELETE("/:id", productApi.DeleteProduct)
	products.PUT("/:id", productApi.UpdateProduct)
	//TODO elasticsearch, redis cache, message broker, notification, payment, basket
	//users
	users := router.Group("/v1/users")
	users.GET("/getbyuuid/:uuid", userApi.FindByUUID)
	users.GET("/getbyusername/:username", userApi.FindByUsername)
	users.POST("", userApi.Insert)
	users.DELETE("/:uuid", userApi.Delete)
	users.PUT("", userApi.Update)

	//categories
	categories := router.Group("/v1/categories")
	categories.Use(middlewares.AuthorizeJWTMiddleware(authSvc))

	categories.Use(middlewares.AuthorizeJWTMiddleware(authSvc))
	categories.GET("", categoryApi.GetAllCategories)
	categories.POST("", categoryApi.CreateCategory)
	categories.GET("/:id", categoryApi.GetCategoryByID)
	categories.DELETE("/:id", categoryApi.DeleteCategory)
	categories.PUT("/:id", categoryApi.UpdateCategory)

	//orders
	orders := router.Group("/v1/orders")
	orders.Use(middlewares.AuthorizeJWTMiddleware(authSvc))

	orders.GET("", orderApi.GetAllOrders)
	orders.POST("", orderApi.CreateOrder)
	orders.GET("/:id", orderApi.GetOrderByID)
	orders.GET("/getbyuserid/:id", orderApi.GetOrdersByUserID)
	orders.DELETE("/:id", orderApi.DeleteOrder)
	orders.PUT("/:id", orderApi.UpdateOrder)

	auth := router.Group("/v1/auth")
	auth.POST("/login", authApi.Login)

	//If env development or test activate swagger
	if version == "dev" || version == "test" {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	return router
}

// @title e-commerce-go swagger
// @version 1.0
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @schemes http
func main() {
	fmt.Printf("Version: %s", version)
	ctx := context.TODO()

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	r := setup(ctx)
	if err := retry.Fibonacci(ctx, 1*time.Second, func(ctx context.Context) error {
		err := r.Run()
		if err != nil {
			fmt.Println(err)
			return retry.RetryableError(err)
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
}
