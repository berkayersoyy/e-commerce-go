package main

import (
	"context"
	"fmt"
	_ "github.com/berkayersoyy/e-commerce-go/docs"
	authService "github.com/berkayersoyy/e-commerce-go/internal/application/services/auth"
	productService "github.com/berkayersoyy/e-commerce-go/internal/application/services/product"
	userService "github.com/berkayersoyy/e-commerce-go/internal/application/services/user"
	productRepository "github.com/berkayersoyy/e-commerce-go/internal/infrastructure/repositories/product"
	userRepository "github.com/berkayersoyy/e-commerce-go/internal/infrastructure/repositories/user"
	authHandler "github.com/berkayersoyy/e-commerce-go/internal/presentation/auth"
	"github.com/berkayersoyy/e-commerce-go/internal/presentation/middlewares"
	productHandler "github.com/berkayersoyy/e-commerce-go/internal/presentation/product"
	userHandler "github.com/berkayersoyy/e-commerce-go/internal/presentation/user"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sethvargo/go-retry"
	"log"
	"time"
)

//version app_version
var version = "dev"

func setup(ctx context.Context) *gin.Engine {
	productRepo := productRepository.ProvideProductRepository()
	productSvc := productService.ProvideProductService(productRepo)
	productApi := productHandler.ProvideProductHandler(productSvc)

	userRepo := userRepository.ProvideUserRepository(time.Second * 30)
	userSvc := userService.ProvideUserService(userRepo)
	userApi := userHandler.ProvideUserHandler(userSvc)

	err := userRepo.CreateTable(ctx)
	if err != nil {
		log.Fatalf("Error on creating users table, %s", err)
	}

	authSvc := authService.ProvideAuthService()
	authApi := authHandler.ProvideAuthHandler(authSvc, userSvc)

	router := gin.Default()

	//products
	products := router.Group("/v1")
	products.Use(middlewares.AuthorizeJWTMiddleware(authSvc))

	products.GET("/products", productApi.GetAllProducts)
	products.POST("/products", productApi.AddProduct)
	products.GET("/products/:id", productApi.GetProductByID)
	products.DELETE("/products/:id", productApi.DeleteProduct)
	products.PUT("/products/:id", productApi.UpdateProduct)

	usersDynamoDb := router.Group("/v1")
	usersDynamoDb.GET("/users/getbyuuid/:uuid", userApi.FindByUUID)
	usersDynamoDb.GET("/users/getbyusername/:username", userApi.FindByUsername)
	usersDynamoDb.POST("/users", userApi.Insert)
	usersDynamoDb.DELETE("/users/:uuid", userApi.Delete)
	usersDynamoDb.PUT("/users", userApi.Update)

	auth := router.Group("/v1")
	auth.POST("/login", authApi.Login)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}

func main() {
	fmt.Printf("Version: %s", version)
	ctx := context.TODO()

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
