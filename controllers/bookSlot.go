package controllers

import (
	mysql "APPOINMENT_BOOKING_SYSTEM/db"
	"APPOINMENT_BOOKING_SYSTEM/models"
	"github.com/gin-gonic/gin"
	mysqldriver "github.com/go-sql-driver/mysql"
	//"strings"
)

func BookAppointment(c *gin.Context) {
	var req models.Booking

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"msg": "invalid request"})
		return
	}

	// 1. Validate datetime
	slotTime := req.DateTime

	valid, err := mysql.IsValidSlot(req.CoachID, slotTime)
	if err != nil {
		c.JSON(500, gin.H{"msg": "error validating slot"})
		return
	}
	if !valid {
		c.JSON(400, gin.H{"msg": "invalid time slot"})
		return
	}

	// 3. Insert booking
	err = mysql.InsertBooking(req.UserID, req.CoachID, slotTime)
	if err != nil {

		if mysqlErr, ok := err.(*mysqldriver.MySQLError); ok && mysqlErr.Number == 1062 {
			c.JSON(409, gin.H{ // 409 = Conflict
				"status": "failed",
				"msg":    "slot already booked",
			})
			return
		}

		// other DB errors
		c.JSON(500, gin.H{
			"status": "failed",
			"msg":    "internal server error",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": "success",
		"msg":    "booking confirmed",
	})
}
