package db

import (
	"go-gin-todo/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error to load env")
	}

	conn := os.Getenv("POSTGRES_URL")

	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	Migrate(db)

	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Todo{})

	// data := models.Todo{}
	// if db.Find(&data).Error(){
	// 	seederTodo(db)
	// }
	var count int64
    result := db.Model(&models.Todo{}).Count(&count)

    if result.Error != nil {
        log.Fatal(result.Error)
    }

    if count == 0 {
        seederTodo(db)
    }
}

func seederTodo(db *gorm.DB){
	data := []models.Todo{
		{
			Title: "test",
			Completed: false,
			DueDate: "0000-00-00",
		},
	}
	for _, v := range data {
		db.Create(&v)
	}

}