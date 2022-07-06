package main

import (
	"context"
	"fmt"
	_ "github.com/berkayersoyy/e-commerce-go/docs"
	authService "github.com/berkayersoyy/e-commerce-go/internal/application/services/auth"
	categoryService "github.com/berkayersoyy/e-commerce-go/internal/application/services/category"
	productService "github.com/berkayersoyy/e-commerce-go/internal/application/services/product"
	userService "github.com/berkayersoyy/e-commerce-go/internal/application/services/user"
	"github.com/joho/godotenv"

	categoryRepository "github.com/berkayersoyy/e-commerce-go/internal/infrastructure/repositories/category"
	productRepository "github.com/berkayersoyy/e-commerce-go/internal/infrastructure/repositories/product"

	userRepository "github.com/berkayersoyy/e-commerce-go/internal/infrastructure/repositories/user"
	authHandler "github.com/berkayersoyy/e-commerce-go/internal/presentation/auth"
	categoryHandler "github.com/berkayersoyy/e-commerce-go/internal/presentation/category"

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

	categoryRepo := categoryRepository.ProvideCategoryRepository()
	categorySvc := categoryService.ProvideCategoryService(categoryRepo)
	categoryApi := categoryHandler.ProvideCategoryHandler(categorySvc)

	err := userRepo.CreateTable(ctx)
	if err != nil {
		log.Fatalf("Error on creating users table, %s", err)
	}

	authSvc := authService.ProvideAuthService()
	authApi := authHandler.ProvideAuthHandler(authSvc, userSvc)

	router := gin.Default()

	//products
	products := router.Group("/v1/products")
	products.Use(middlewares.AuthorizeJWTMiddleware(authSvc))

	products.GET("", productApi.GetAllProducts)
	products.GET("/getbycategoryid/:id", productApi.GetProductByCategoryID)
	products.POST("", productApi.AddProduct)
	products.GET("/:id", productApi.GetProductByID)
	products.DELETE("/:id", productApi.DeleteProduct)
	products.PUT("/:id", productApi.UpdateProduct)

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
	categories.GET("", categoryApi.GetAllCategories)
	categories.POST("", categoryApi.AddCategory)
	categories.GET("/:id", categoryApi.GetCategoryByID)
	categories.DELETE("/:id", categoryApi.DeleteCategory)
	categories.PUT("/:id", categoryApi.UpdateCategory)

	auth := router.Group("/v1/auth")
	auth.POST("/login", authApi.Login)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
