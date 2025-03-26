package config

import (
	"dairy-management-backend/entities"
	"fmt"
	"os"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ✅ Now returns an error
func ConnectDB() error {
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		"DairyManagement",
	)

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	DB = db
	fmt.Println("✅ Database connected successfully")
	return nil
}

// ✅ MigrateDB now correctly handles uninitialized DB
func MigrateDB() error {
	if DB == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	err := DB.AutoMigrate(&entities.Cow{}, &entities.User{}, &entities.MilkCollection{}) // Add all entities here
	if err != nil {
		return fmt.Errorf("failed to run migrations: %v", err)
	}

	fmt.Println("✅ Database migrated successfully")
	return nil
}
