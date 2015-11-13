package main
import (
	"github.com/gin-gonic/gin"
)
func main() {
	// ==================== START ========================= //
	// init db (migrate)
	

	// ==================== ROUTES ======================== //
	route := gin.Default()
	route.GET("/", Index)
	route.Run(":9001")
}