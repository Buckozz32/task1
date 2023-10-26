package main

import (
	"encoding/json"
	"io/ioutil"

	"net/http"

	"github.com/gin-gonic/gin"
)

// Структура для хранения данных о человеке
type Nationalities struct {
	Name          string            `json:"name"`
	Surname       string            `json:"surname"`
	Patronymic    string            `json:"patronymic"`
	Nationalities []NationalityData `json:"nationalities"`
}

// Структура для хранения данных о национальности
type NationalityData struct {
	CountryID   string  `json:"country_id"`
	Probability float64 `json:"probability"`
}


func getNationality(name string) ([]NationalityData, error) {
	response, err := http.Get("https://api.nationalize.io/?name=" + name)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Nationalities []NationalityData `json:"country"`
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result.Nationalities, nil
}


func AddPerson1(c *gin.Context) {
	var person Person
	err := c.ShouldBindJSON(&person)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	nationalities, err := getNationality(person.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get nationalities"})
		return
	}

	person.Nationalities = nationalities

	c.JSON(http.StatusCreated, person)
}
