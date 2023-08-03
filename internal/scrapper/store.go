package scrapper

import (
	"encoding/json"
	"io/ioutil"
)

func saveDataToFile(filename string, data []Vacancy) error {
	dataBytes, err := json.Marshal(data)
	if err != nil || dataBytes == nil {
		return err
	}

	return ioutil.WriteFile(filename, dataBytes, 0644)
}

func readVacanciesFromFile(filename string) ([]Vacancy, error) {
	dataBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var vacancies []Vacancy
	err = json.Unmarshal(dataBytes, &vacancies)
	if err != nil {
		return nil, err
	}

	return vacancies, nil
}
