package db

import (
	"github.com/andreiz53/go-auth-starter/types"
	"gorm.io/gorm"
)

func CreateTables(db *gorm.DB) {
	if err := db.AutoMigrate(&types.User{}); err != nil {
		panic("failed to create users table")
	}
}
