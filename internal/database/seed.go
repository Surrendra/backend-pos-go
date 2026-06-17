package database

import (
	seeder "BackendPOS/internal/seeder/authentication"
	"log"

	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	log.Println("Seeding database....")
	seeder.SeedUser(db)

	log.Println("database seeding complete")
}
