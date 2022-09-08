package telegram

import (
	"Projects/WordAnalytics/internal/counter"
	"Projects/WordAnalytics/internal/parser"
	"Projects/WordAnalytics/pkg/logger"
	"Projects/WordAnalytics/pkg/postgresql"
	"database/sql"
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

const token = "5732743787:AAGjb6KK2x_iJd8yGuWuJDbXQz6WJ_MBNck"

var UrlChan string

type DataBase struct {
	DB *sql.DB
}

func NewStore(db *sql.DB) *DataBase {
	return &DataBase{DB: db}
}

func Bot() {
	logg := logger.GetLogger()

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		logg.Fatal(err)
	}

	bot.Debug = false

	logg.Infof("Authorized on account %s", bot.Self.UserName)

	checkUpdates(bot)
}

func checkUpdates(bot *tgbotapi.BotAPI) {
	logg := logger.GetLogger()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {

		switch update.Message.Text {
		case "/start":
			logg.Infof("[%s] %s", update.Message.From.UserName, update.Message.Text)

			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Привет, я умею считать слова на любом сайте! Введи /getUrl чтобы продолжить"))
		case "/getUrl":
			logg.Infof("[%s] %s", update.Message.From.UserName, update.Message.Text)

			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Вставь свою ссылку, и через пробел напиши слово"))
		default:
			logg.Infof("[%s] %s", update.Message.From.UserName, update.Message.Text)

			arr := strings.Split(update.Message.Text, " ")

			if len(arr) != 2 {
				logg.Infof("[%s] %s", update.Message.From.UserName, update.Message.Text)

				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Неверный формат! Давай ещё раз"))
			}
			if parser.IsUrl(arr[0]) == false {
				logg.Infof("[%s] %s", update.Message.From.UserName, update.Message.Text)

				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Это не ссылка! Давай ещё раз"))
			}

			url := arr[0]
			word := arr[1]

			result := findResult(url, word)

			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Количество слов: %d", result)))
		}
	}
}

func findResult(url, word string) int {
	log := logger.GetLogger()

	str := parser.Parse(url)
	objects := counter.Count(str)
	jsonObj, _ := json.Marshal(objects)

	for i, el := range objects {
		fmt.Println(i, el)
	}
	db, err := postgresql.Connect()
	if err != nil {
		log.Fatal("failed to connect ")
	}
	log.Info("Connected successful")
	storage := NewStore(db)

	postgresql.Insert(url, jsonObj, storage.DB)
	id := postgresql.SelectfromWords(storage.DB)
	parsed := postgresql.Select(db, id)
	result := parser.FindResult(parsed, word)

	return result
}
