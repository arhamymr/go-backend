package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB;

func InitDatabase() {
	dsn := "host=ep-blue-limit-309264-pooler.ap-southeast-1.postgres.vercel-storage.com user=default password=zLfXyUm6g5Ax dbname=verceldb port=5432"
	var err error;
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("failed to connect databases")
		return
	}

	fmt.Println("Database initialized successfully")
}
