package postgresql

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "telebot_db"
)

func Connecting() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		panic(err)
	}

	log.Print("connection successful")

	sqlstatment := `INSERT INTO sites (url,info, created_at, updated_at, deleted_at)
                    VALUES ($1, $2, $3, $4, $5)
                    `

	_, err = db.Exec(sqlstatment, "https://habr.com/ru/post/654569/", "package : ", "2022-08-21", nil, nil)
	if err != nil {
		panic(err)
	}
	log.Print("Inserting is successful")
}

//func insert(db sql.DB) {
//	sqlstatment := `INSERT INTO sites (url,info, created_at, updated_at, deleted_at)
//                    VALUES ($1, $2, $3, $4, $5)
//                    `
//
//	_, err := db.Exec(sqlstatment, "https://habr.com/ru/post/654569/", "package : ", nil, nil, nil)
//	if err != nil {
//		panic(err)
//	}
//	log.Print("Inserting is successful")
//}
