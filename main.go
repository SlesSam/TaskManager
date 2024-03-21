package main

import (
	"taskmanager/cmd"
	"taskmanager/internal/tasks"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=localhost user=adminuser password=admin dbname=mytasksdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&tasks.Task{})

	cmd.Comands(db)
}
