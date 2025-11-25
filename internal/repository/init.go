package repository

import (
	"github.com/yudin-artem/avito_test_2025/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"fmt"
)

func DBInit() (*gorm.DB, error) {
	dns, err := config.Get("DATABASE_URL")
	if err != nil {
		return nil, fmt.Errorf("faild conn %w", err)
	}
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("faild open %w", err)
	}

	return db, nil
}