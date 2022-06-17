package main

import (
	"github.com/gin-gonic/gin"
	"hostpitalProject/controllers"
	"hostpitalProject/models"
)

func main() {

	router := gin.Default()

	models.ConnectDatabase()

	router.POST("/doctor", controllers.DocPost)
	router.GET("/doctor/:id", controllers.DocGet)
	router.PATCH("/doctor/:id", controllers.DocPatch)

	router.POST("/patient", controllers.PatPost)
	router.GET("/patient/:id", controllers.PatGet)
	router.PATCH("/patient/:id", controllers.PatPatch)

	router.GET("/doctorPatientList/:id", controllers.GetDoctorPatientList)
	err := router.Run()
	if err != nil {
		return
	}
}
