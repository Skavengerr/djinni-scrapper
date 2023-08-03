package scrapper

import (
	"crypto/md5"
	"encoding/hex"
)

func isVacancyExists(existingVacancies []Vacancy, newVacancy Vacancy) bool {
	for _, vacancy := range existingVacancies {
		if vacancy.Title == newVacancy.Title && vacancy.URL == newVacancy.URL {
			return true
		}
	}
	return false
}

func generateHash(title, url string) string {
	data := title + url
	hasher := md5.New()
	hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil))
}
