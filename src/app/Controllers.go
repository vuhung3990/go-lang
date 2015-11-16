package main
import (
	"github.com/gin-gonic/gin"
	"strconv"
	"fmt"
)
func Index(c *gin.Context) {
	c.String(200, "hello world")
}

func Create(c *gin.Context) {
	//todo validation
	var student Student
	if c.Bind(&student) == nil{
		db.Create(&student)
		c.JSON(200, JsonResponse{Action:"Create", IsSuccess: true}.setJsonResponseData(student))
	}else {
		c.JSON(400, JsonResponse{Action:"Create", IsSuccess: false}.setJsonResponseData(c.Bind(&student).Error()))
	}
}

// update not working
func Edit(c *gin.Context) {
	// c.Param instead of c.PostForm
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	student := Student{}
	// find in db
	db.First(&student, id)
	student.Name = "vuhung 111"
	db.Save(student)

	fmt.Println(student)
}

func Delete(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	db.Delete(Student{Id:int(id)})
	c.JSON(200, JsonResponse{Action:"Edit", IsSuccess: true})
}