package main

import (
  "net/http"
  
  "github.com/gin-gonic/gin"
)

type student struct {
  ID string `json:"id"`
  Name string `json:"name"`
  Age int64 `json:"age"`
  Course string `json:"course"`
}

var students = []student{
  {ID:"1", Name:"Pedro", Age:23, Course:"ADS"},
  {ID:"2", Name:"Ana", Age:18, Course:"ADS"},
  {ID:"3", Name:"Diogo", Age:20, Course:"ADS"},
}

func main(){
  router := gin.Default()
  router.GET("/students", getStudents)
  router.POST("/students", postStudents)
  router.GET("/students/:id", getAlbumByID)
  router.Run("localhost:8080")
}

//getStudents respons with the list of all Students
func getStudents(c *gin.Context){
  c.IndentedJSON(http.StatusOK, students)
}

func postStudents(c *gin.Context){
  var newStudent student
  if err := c.BindJSON(&newStudent); err != nil{
    return
  }
  students = append(students, newStudent)
  c.IndentedJSON(http.StatusCreated, newStudent)
}

func getAlbumByID(c *gin.Context){
  id := c.Param("id")
  
  for _, s := range students{
    if s.ID == id {
      c.IndentedJSON(http.StatusOK, s)
      return
    }
  }
  c.IndentedJSON(http.StatusNotFound, gin.H{"message":"student not found"})
}