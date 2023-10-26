package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Структура для хранения ФИО
type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var people []Person

// Обработка POST запроса для добавления ФИО
func AddPerson(c *gin.Context) {
	var person Person
	err := c.BindJSON(&person)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	person.ID = len(people) + 1
	people = append(people, person)

	c.Status(http.StatusCreated)
}

func GetPeople(c *gin.Context) {
	c.JSON(http.StatusOK, people)
}

func GetPersonByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if id <= 0 || id > len(people) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		return
	}

	person := people[id-1]
	c.JSON(http.StatusOK, person)
}

func DeletePersonByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	person := findPersonByID(id)
	if person == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		return
	}

	people = removePersonByID(id)

	c.JSON(http.StatusOK, gin.H{"message": "Person deleted"})
}

func findPersonByID(id int) *Person {
	for i := 0; i < len(people); i++ {
		if people[i].ID == id {
			return &people[i]
		}
	}
	return nil
}

// Вспомогательная функция для удаления ФИО из слайса по ID
func removePersonByID(id int) []Person {
	var updatedPeople []Person
	for i := 0; i < len(people); i++ {
		if people[i].ID != id {
			updatedPeople = append(updatedPeople, people[i])
		}
	}
	return updatedPeople
}

// Обработка GET запроса с фильтрацией и пагинацией
func SearchPeople(c *gin.Context) {

	name := c.Query("name")
	pageStr := c.Query("page")
	sizeStr := c.DefaultQuery("size", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1 // Значение по умолчанию
	}

	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		size = 10 // Значение по умолчанию
	}

	filteredPeople := make([]Person, 0)
	for _, person := range people {
		if name != "" && person.Name != name {
			continue
		}
		filteredPeople = append(filteredPeople, person)
	}

	startIndex := (page - 1) * size
	endIndex := page * size
	if startIndex >= len(filteredPeople) {
		startIndex = len(filteredPeople)
	}
	if endIndex > len(filteredPeople) {
		endIndex = len(filteredPeople)
	}

	paginatedPeople := filteredPeople[startIndex:endIndex]

	c.JSON(http.StatusOK, paginatedPeople)
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.POST("PeoplePost", AddPerson)

	r.GET("PepleGet", GetPeople)

	r.GET("PersonId", GetPersonByID)

	r.GET("PeopleSerch", SearchPeople)

	r.GET("deletePerson", DeletePersonByID)

	r.Run(":8080")
}
