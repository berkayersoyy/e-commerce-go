package providers

import (
	"context"
	"fmt"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/models"
	"github.com/jinzhu/gorm"
	"github.com/sethvargo/go-retry"
	"log"
	"os"
	"time"
)

//NewMysqlDb Init db
func NewMysqlDb() *gorm.DB {
	dsn := os.Getenv("MYSQL_DSN")
	ctx := context.Background()
	var db *gorm.DB
	var err error
	if err := retry.Fibonacci(ctx, 1*time.Second, func(ctx context.Context) error {
		db, err = gorm.Open("mysql", dsn)
		if err != nil {
			fmt.Println(err)
			return retry.RetryableError(err)
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}

	db.DB().SetMaxOpenConns(10)
	db.DB().SetMaxIdleConns(5)

	db.AutoMigrate(&models.Order{})
	db.AutoMigrate(&models.Category{})
	db.AutoMigrate(&models.OrderItem{})
	db.AutoMigrate(&models.Product{})

	return db
}
