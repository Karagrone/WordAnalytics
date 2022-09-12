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

type DataBase struct {
	DB *sql.DB
}

func BotRun() {
	log := logger.GetLogger()

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = false

	log.Infof("Authorized on account %s", bot.Self.UserName)

	checkUpdates(bot)
}

func checkUpdates(bot *tgbotapi.BotAPI) {
	log := logger.GetLogger()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {

		switch update.Message.Text {
		case "/start":
			log.Infof("[%s] %s", update.Message.From.UserName, update.Message.Text)

			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Hi, i can read your url, and count words. Also i can say amount of your words. If you want continue type /getUrl"))
		case "/getUrl":
			log.Infof("[%s] %s", update.Message.From.UserName, update.Message.Text)

			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Type url and word through a space"))
		default:
			log.Infof("[%s] %s", update.Message.From.UserName, update.Message.Text)

			arr := strings.Split(update.Message.Text, " ")

			if len(arr) != 2 {
				log.Infof("[%s] %s", update.Message.From.UserName, update.Message.Text)

				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Wrong format! Come on one more time"))
			}
			if parser.IsUrl(arr[0]) == false {
				log.Infof("[%s] %s", update.Message.From.UserName, update.Message.Text)

				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "This is not a link! Come on one more time"))
			}

			url := arr[0]
			word := arr[1]

			result := findResult(url, word)

			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Number of words: %d", result)))
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

	postgresql.Insert(url, jsonObj, db)
	parsed := postgresql.Select(db)
	result := parser.FindResult(parsed, word)

	return result
}
