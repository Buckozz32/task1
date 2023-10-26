package main

import (
 "encoding/json"
 "io/ioutil"
 
 "net/http"

 "github.com/gin-gonic/gin"
)


type Person struct {
 Name       string `json:"name"`
 Surname    string `json:"surname"`
 Patronymic string `json:"patronymic"`
 Gender     string `json:"gender"`
}


func getGender(name string) (string, error) {
 response, err := http.Get("https://api.genderize.io/?name=" + name)
 if err != nil {
  return "", err
 }
 defer response.Body.Close()

 body, err := ioutil.ReadAll(response.Body)
 if err != nil {
  return "", err
 }

 var result struct {
  Gender string `json:"gender"`
 }
 err = json.Unmarshal(body, &result)
 if err != nil {
  return "", err
 }

 return result.Gender, nil
}




func AddPersons (c *gin.Context) {
 var person Person
 err := c.ShouldBindJSON(&person)
 if err != nil {
  c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
  return
 }

 gender, err := getGender(person.Name)
 if err != nil {
  c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get gender"})
  return
 }

 person.Gender = gender

 c.JSON(http.StatusCreated, person)
}
