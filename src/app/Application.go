package main
import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"	// postgres SQL driver
)

// ======================= Global Variables ============== //
var db gorm.DB

func main() {
	// ==================== START ========================= //
	// init db
	dbm, err := gorm.Open("postgres", "user=golang_admin password='v0thuong' dbname=golang_database sslmode=disable")  //mysql[user:password@host/database]

	// if fail close app
	if err != nil {
		panic(err)
	}

	// save db global
	db = dbm;

	// auto migrate db
	db.AutoMigrate(&Student{})

	// ==================== ROUTES ======================== //
	route := gin.Default()
	route.GET("/", Index)

	// sample restful
	route.POST("/student", Create)
	route.PUT("/student/:id", Edit)
	route.DELETE("/student/:id", Delete)

	// ================= Listen on port 9001 =========== //
	route.Run(":9001")
}