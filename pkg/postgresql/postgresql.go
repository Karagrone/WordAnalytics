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

//func psqlInfo(host, user, password, dbname string, port int) string {
//	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
//
//	return psqlInfo
//}

func Connecting() (db *sql.DB) {
	//psqlInfo := psqlInfo(host, dbname, user, password, port)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	log.Print("connection successful")

	return db
}

func Insert(db *sql.DB) {
	sqlstatment := `INSERT INTO sites (url,info, created_at, updated_at, deleted_at)
                    VALUES ($1, $2, $3, $4, $5) `

	_, err := db.Exec(sqlstatment, "https://habr.com/ru/post/654569/", "package : 1", "2022-08-22", nil, nil)
	if err != nil {
		panic(err)
	}
	log.Print("Inserting is successful")
}

func Delete(db *sql.DB) {
	sqlstatment := `DELETE FROM sites
                    WHERE id = $1`
	res, err := db.Exec(sqlstatment, 2)
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	log.Printf("Deleted successful, %d", count)
}

func Update(db *sql.DB) {
	sqlstatment := `UPDATE sites
                    SET url = $2
                    WHERE id = $1`
	res, err := db.Exec(sqlstatment, 3, "google.com")
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	log.Printf("Updated successful, %d", count)
}
