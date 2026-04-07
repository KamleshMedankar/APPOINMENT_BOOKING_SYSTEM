package controllers

import (
	mysql "APPOINMENT_BOOKING_SYSTEM/db"
	"net/http"
	"strconv"
	"strings"
	"github.com/gin-gonic/gin"
)

func GetAvailableSlots(c *gin.Context) {
	coachIDParam := c.Param("coach_id")
	date := strings.TrimSpace(c.Query("date"))

	coachID, err := strconv.Atoi(coachIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid coach_id"})
		return
	}

	if date == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "date is required"})
		return
	}

	slots, err := mysql.FetchAvailableSlots(coachID, date)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to fetch slots"})
		return
	}

	c.JSON(200, gin.H{
		"status":"success",
		"slots": slots,
	})
}
