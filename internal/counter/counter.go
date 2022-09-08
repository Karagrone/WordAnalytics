package counter

import (
	"Projects/WordAnalytics/pkg/logger"
	"log"
	"regexp"
	"strings"
	"unicode"
)

type Word struct {
	WordID int
	Word   string
	Count  int
}

func Count(str string) []Word {
	logging := logger.GetLogger()

	logging.Info("Formatting text")
	str = formatText(str)
	strArr := strings.Split(str, " ")
	var strMap = map[string]int{}

	logging.Info("Checking result")
	for _, value := range strArr {
		if isWord(value) {
			if _, isExist := strMap[value]; isExist {
				strMap[value] += 1
			} else {
				strMap[value] = 1
			}
		}

	}

	logging.Info("Result out")
	return toObjectArr(strMap)
}

func toObjectArr(strMap map[string]int) []Word {
	var arr []Word
	for value, key := range strMap {
		if value == "" {
			continue
		} else {
			arr = append(
				arr,
				Word{
					Word:  value,
					Count: key,
				})
		}
	}

	return arr
}

func isWord(str string) bool {
	for _, value := range str {
		if unicode.IsDigit(rune(value)) {
			return false
		}
	}
	return true
}

func formatText(str string) string {
	str = strings.ReplaceAll(str, "/", "")
	str = strings.ReplaceAll(str, ",", " ")
	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		log.Fatal(err)
	}
	str = re.ReplaceAllString(str, " ")

	return str
}
