package postgresql

import (
	"Projects/WordAnalytics/internal/counter"
	"Projects/WordAnalytics/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Sites struct {
	//ID        int `gorm:"PrimaryKey: ID"`
	gorm.Model
	Url  string
	Info []counter.Word `gorm:"foreign_key:WordID"`
	//CreatedAt string
	//UpdatedAt string
	//DeletedAt string
}

type Handler struct {
	DB *gorm.DB
}

func Conn() *gorm.DB {
	log := logger.GetLogger()
	dsn := "host=localhost user=postgres password=root dbname=telebot_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Errorf("Open postgres failed: &e", err)
	}

	db.AutoMigrate(&Sites{})

	return db
}

func (h Handler) Insert(site Sites) {
	log := logger.GetLogger()

	if result := h.DB.Select("Url", "Info", "CreatedAt").Create(&Sites{}); result != nil {
		log.Errorf("Inserting is failed")
	}
}
