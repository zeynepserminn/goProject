package postgres

import (
	"fmt"
	"goProject/internal/core/dto"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func InitDB() (*Database, error) {
	dsn := "host=localhost port=8080 user=postgres password=zeynoo.15 dbname=postgres sslmode=disable"

	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(fmt.Errorf("failed to connect database %w", err))
	}

	fmt.Println("Successfully connected to database")

	return &Database{db: db}, err

}
func (d *Database) AutoMigrate(models ...interface{}) error {
	err := d.db.AutoMigrate(models)
	if err != nil {
		fmt.Println("Failed to migrate database")
	}
	return nil

}

func (d *Database) GetInstance() *gorm.DB {
	return d.db
}
func PaginatedResult(p dto.PaginationRequest) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (p.Skip - 1) * p.Limit
		query := db.Offset(offset).Limit(p.Limit)
		if p.SortBy != "" {
			query = query.Order(p.SortBy + " " + p.OrderBy)
		}
		return query
	}

}
