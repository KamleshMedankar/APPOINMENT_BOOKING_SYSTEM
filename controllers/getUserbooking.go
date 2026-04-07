package controllers
import(
	mysql "APPOINMENT_BOOKING_SYSTEM/db"
	"strconv"
	"github.com/gin-gonic/gin"
	"log"

)

func GetUserBookings(c *gin.Context) {

	userIDParam := c.Param("user_id")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		c.JSON(400, gin.H{
			"status": "failed",
			"msg":    "invalid user_id",
		})
		return
	}
	log.Println("UserID:", userID)
	bookings, err:= mysql.GetUserBookings(userID)

	if err != nil {
		c.JSON(500, gin.H{
			"status": "failed",
			"msg":    "failed to fetch bookings",
		})
		return
	}


	c.JSON(200, gin.H{
		"status":   "success",
		"bookings": bookings,
	})
}