package postgresql

import (
	"Projects/WordAnalytics/pkg/logger"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "telebot_db"
)

type Data struct {
	Id        int
	Url       string
	Info      string
	CreatedAt string
}

var Url string

//func psqlInfo(host, user, password, dbname string, port int) string {
//	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
//
//	return psqlInfo
//}

func Connect() (*sql.DB, error) {
	log := logger.GetLogger()
	//psqlInfo := psqlInfo(host, dbname, user, password, port)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Errorf("Opening postgres failed: %e", err)
	}

	if err = db.Ping(); err != nil {
		log.Errorf("Opening postgres failed: %e", err)
	}

	log.Info("connection successful")

	return db, nil
}

func Insert(url string, objects []byte, db *sql.DB) {
	log := logger.GetLogger()
	insertToWords(db, objects)
	id := SelectfromWords(db)
	if id == 0 {
		log.Fatal("Can't select from table words")
	}

	sqlstatment := `INSERT INTO sites (url ,info ,created_at ,updated_at ,deleted_at)
                   VALUES ($1, $2, $3, $4, $5)`

	fmt.Println(url, objects)

	_, err := db.Exec(sqlstatment, url, id, "2022-09-06", nil, nil)
	if err != nil {
		log.Errorf("Can't insert to table sites", err)
	}
	log.Info("Inserting is successful")
}

func Select(db *sql.DB) []byte {
	log := logger.GetLogger()

	rows, err := db.Query("SELECT words FROM words;")
	if err != nil {
		log.Fatalf("Can't select from table words: %e", err)
	}
	defer rows.Close()

	for rows.Next() {
		var title []byte
		if err := rows.Scan(&title); err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(title))
		return title
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return nil
}

func insertToWords(db *sql.DB, objects []byte) {
	log := logger.GetLogger()
	sqlstatement := `INSERT INTO words (words)
                   VALUES ($1)`

	_, err := db.Exec(sqlstatement, objects)
	if err != nil {
		log.Errorf("Can't insert to table words: &e", err)
	}

	log.Info("inserted to words successfully")
}

func SelectfromWords(db *sql.DB) int {
	log := logger.GetLogger()
	var id int

	sqlstatement := `SELECT currval('words_id_seq')`
	row := db.QueryRow(sqlstatement)
	switch err := row.Scan(&id); err {
	case sql.ErrNoRows:
		log.Errorf("No rows were returned!")
	case nil:
		return id
	default:
		log.Fatal("selecting failed")
	}

	return 0
}
