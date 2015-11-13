package main
import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

// ======================= Global Variables ============== //
var db gorm.DB

func main() {
	// ==================== START ========================= //
	// init db
	dbm, err := gorm.Open("mysql", "root:@/test?charset=utf8&parseTime=True")  // user:password@host/database

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
	route.Run(":9001")
}