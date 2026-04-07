package controllers

import (
	mysql "APPOINMENT_BOOKING_SYSTEM/db"
	"APPOINMENT_BOOKING_SYSTEM/models"
	"APPOINMENT_BOOKING_SYSTEM/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

func SetAvailability(c *gin.Context) {

	coachIDParam := c.Param("coach_id")
	coachID, err := strconv.Atoi(coachIDParam)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid coach_id",
		})
		return
	}

	var req models.Availability
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON body",
		})
		return
	}
	
	if req.Day == "" || req.Starttime == "" || req.Endtime == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "All fields are required",
		})
		return
	}

	day, ok := utils.IsValidDay(req.Day)
	if !ok {
		c.JSON(400, gin.H{
			"status": "failed",
			"msg":    "invalid day, must be Monday to Sunday",
		})
		return
	}

	_, err = time.Parse("15:04", req.Starttime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid start_time format (use HH:MM)",
		})
		return
	}

	_, err = time.Parse("15:04", req.Endtime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid end_time format (use HH:MM)",
		})
		return
	}
	exists, err := mysql.CheckCoachExists(coachID)
	if err != nil {
		log.Println("error checking coach", err)
		return
	}

	if !exists {
		log.Println("coach does not exist")
		c.JSON(400, gin.H{
			"status": "failed",
			"msg":    "coache not found",
		})
		return
	}
	
	exists, err = mysql.CheckAvailabilityExists(coachID, day)
	if err != nil {
		log.Println("error checking availability", err)
		return
	}

	if exists {
		c.JSON(400, gin.H{
			"status": "failed",
			"msg":    "availability already exists for this day",
		})
		return
	}

	limitReached, _ := mysql.CheckMaxAvailabilityReached(req.CoachID)
	if limitReached {
		c.JSON(400, gin.H{"msg": "only 7 days allowed","status":"failed"})
		return
	}

	err = mysql.InsertAvailability(coachID, day, req.Starttime, req.Endtime)
	if err != nil {
		log.Println("unable to insert availability", err)
		c.JSON(500, gin.H{
			"msg": "internal server error",
		})
		return
	}

	req.CoachID = coachID

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Availability set successfully",
		"data":    req,
	})
}
