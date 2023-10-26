package main

import (
	"encoding/json"
	"io/ioutil"

	"net/http"

	"github.com/gin-gonic/gin"
)

type PersonAge struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Age        int    `json:"age"`
}

func getAge(name string) (int, error) {
	response, err := http.Get("https://api.agify.io/?name=" + name)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, err
	}

	var result struct {
		Age int `json:"age"`
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return 0, err
	}

	return result.Age, nil
}

func AddPerson3(c *gin.Context) {
	var person Person
	err := c.ShouldBindJSON(&person)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	age, err := getAge(person.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get age"})
		return
	}

	person.PersonAge = age

	c.JSON(http.StatusCreated, person)
}
