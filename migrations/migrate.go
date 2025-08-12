package database

import (
	"log"

	"codelabs-backend-fiber/internal/user/domain"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	log.Println("🚀 Starting migration...")

	err := db.Exec(`
		DO $$
		BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_role') THEN
				CREATE TYPE user_role AS ENUM ('admin', 'user');
			END IF;
		END
		$$;
	`).Error
	if err != nil {
		log.Fatalf("❌ Failed to create user_role enum: %v", err)
	} else {
		log.Println("✅ Enum type 'user_role' created or already exists.")
	}

	err = db.AutoMigrate(&domain.User{})
	if err != nil {
		log.Fatalf("❌ AutoMigrate failed: %v", err)
	} else {
		log.Println("✅ AutoMigrate successful.")
	}
}
