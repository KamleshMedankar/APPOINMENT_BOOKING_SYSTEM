package routes

import (
	"APPOINMENT_BOOKING_SYSTEM/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(router *gin.Engine) {

	// Coach availability
	router.POST("/coaches/:coach_id/availability", controllers.SetAvailability)

	// Get slots for a coach
	router.GET("/coaches/:coach_id/slots", controllers.GetAvailableSlots)

	// Booking APIs
	router.POST("users/bookings", controllers.BookAppointment)
	router.GET("/users/:user_id/bookings", controllers.GetUserBookings)

	// Cancel booking
	// router.DELETE("/bookings/:booking_id", controllers.CancelBooking)
}
