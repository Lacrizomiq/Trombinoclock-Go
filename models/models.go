package models

import (
	"encoding/json"
	"fmt"
	"os"
)

type Promos struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	GitHub string `json:"github_organization"`
}

type Students struct {
	ID                int    `json:"id"`
	Promo             string `json:"promo"`
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	ProfilePictureUrl string `json:"profile_picture_url"`
}

func GetPromos(filename string) ([]Promos, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	var promos []Promos
	err = json.Unmarshal(bytes, &promos)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	return promos, nil
}

func GetStudents(filename string) ([]Students, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	var students []Students
	err = json.Unmarshal(bytes, &students)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	return students, nil
}
