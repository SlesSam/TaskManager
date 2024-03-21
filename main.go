package main

import (
	"taskmanager/internal/tasks"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	//TODO
	dsn := "host=localhost user=adminuser password=admin dbname=mydatabase port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&tasks.Task{})

	// cmd.Execute(db)
}
