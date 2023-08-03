package scrapper

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	djinniURL = "https://djinni.co/jobs/?all-keywords=&any-of-keywords=&exclude-keywords=&primary_keyword=Golang&employment=remote"
	douURL    = "https://jobs.dou.ua/vacancies/?remote&category=Golang"
)

type Vacancy struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

func StartScraping(chatID int64, bot *tgbotapi.BotAPI) {
	fileName := fmt.Sprintf("%d_vacancies.json", chatID)
	existingVacancies, err := readVacanciesFromFile(fileName)
	if err != nil {
		log.Println("Error loading data:", err)
	}

	urls := map[string]string{
		"djinni": djinniURL,
		"dou":    douURL,
	}

	newVacancies := GetVacanciesByUrl(urls)

	var uniqueVacancies []Vacancy
	for _, newVacancy := range newVacancies {
		if existingVacancies == nil || !isVacancyExists(existingVacancies, newVacancy) {
			uniqueVacancies = append(uniqueVacancies, newVacancy)
		}
	}

	// TODO: use hash IDs for saving data
	//hashVacanciesMap := make(map[string]Vacancy)

	//for _, vacancy := range uniqueVacancies {
	//	hash := generateHash(vacancy.Title, vacancy.URL)
	//	hashVacanciesMap[hash] = vacancy
	//}

	//for hash, vacancy := range hashVacanciesMap {
	//	fmt.Printf("ID: %s, Title: %s, URL: %s\n", hash, vacancy.Title, vacancy.URL)
	//}

	combinedVacancies := append(existingVacancies, uniqueVacancies...)

	err = saveDataToFile(fileName, combinedVacancies)
	if err != nil {
		log.Println("Error saving data:", err)
	}

	for _, vacancy := range uniqueVacancies {
		text := fmt.Sprintf("%s\n%s", vacancy.Title, vacancy.URL)

		msg := tgbotapi.NewMessage(chatID, text)
		_, sendErr := bot.Send(msg)
		if sendErr != nil {
			log.Println("Error sending message:", sendErr)
		}
	}
}
