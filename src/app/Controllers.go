package main
import (
	"github.com/gin-gonic/gin"
	"strconv"
)
func Index(c *gin.Context) {
	c.String(200, "hello world")
}

func Create(c *gin.Context) {
	age, _ := strconv.ParseInt(c.PostForm("age"), 10, 64)
	gender, _ := strconv.ParseBool(c.PostForm("gender"))
	student := Student{
		Name: c.PostForm("name"),
		Address: c.PostForm("address"),
		Age: int(age),
		Class: c.PostForm("class"),
		Gender: gender,
	}

	//todo validation

	db.Create(&student)
	c.JSON(200, JsonResponse{Action:"Create", IsSuccess: true}.setJsonResponseData(student))
}

// update not working
func Edit(c *gin.Context) {
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	db.First(Student{}, id).Update(Student{Name:"van hung", Address:"tuyen quang"})
	c.JSON(200, JsonResponse{Action:"Edit", IsSuccess: true})
}

func Delete(c *gin.Context) {
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	db.Delete(Student{Id:int(id)})
	c.JSON(200, JsonResponse{Action:"Edit", IsSuccess: true})
}