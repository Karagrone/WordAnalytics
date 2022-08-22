package main

import (
	"Projects/WordAnalytics/internal/counter"
	"Projects/WordAnalytics/internal/parser"
	"Projects/WordAnalytics/pkg/logger"
	"Projects/WordAnalytics/pkg/postgresql"
	"fmt"
)

func main() {
	logging := logger.GetLogger()

	logging.Info("Set URL for parser")
	str := parser.Parse("https://habr.com/ru/post/654569/")
	objects := counter.Count(str)

	for value := range objects {
		fmt.Println(objects[value])
	}

	db := postgresql.Connecting()

	postgresql.Update(db)

}
