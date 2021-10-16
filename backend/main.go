package main

import (
	"fmt"

	"github.com/S1CKK/sa-64-example/controller"

	"github.com/S1CKK/sa-64-example/entity"

	"github.com/gin-gonic/gin"
)

func main() {

	entity.SetupDatabase()
	/*gin.SetMode(gin.ReleaseMode)*/
	r := gin.Default()
	r.Use(CORSMiddleware())

	// User Routes

	r.GET("/users", controller.ListUsers)
	r.GET("/user/:id", controller.GetUser)
	r.POST("/users", controller.CreateUser)
	//r.POST("/users", controller.CreateMedicationRacord)
	r.PATCH("/users", controller.UpdateUser)
	r.DELETE("/users/:id", controller.DeleteUser)

	r.GET("/admissions", controller.ListAdmission)
	r.GET("/admission/:id", controller.GetAdmission)
	r.POST("/admissions", controller.CreateAdmission)
	r.PATCH("/admissions", controller.UpdateAdmission)
	r.DELETE("/admission/:id", controller.DeleteAdmission)

	r.GET("/medicines", controller.ListMedicine)
	r.GET("/medicine/:id", controller.GetMedicine)
	r.POST("/medicines", controller.CreateMedicine)
	r.PATCH("/medicines", controller.UpdateMedicine)
	r.DELETE("/medicine/:id", controller.DeleteMedicine)
	// Run the server

	r.GET("/medicationrecords", controller.ListMedicationRacord)
	r.GET("/medicationrecords/:id", controller.GetMedicationRacord)
	r.POST("/medicationrecords", controller.CreateMedicationRecord)
	r.PATCH("/medicationrecords", controller.UpdateMedicationRacord)
	r.DELETE("/medicationrecords/:id", controller.DeleteMedicationRacord)

	r.GET("/treatments", controller.ListTreatmentRecord)
	r.GET("/treatments/:id", controller.GetTreatmentRecord)
	r.POST("/treatments", controller.CreateTreatmentRecord)
	r.PATCH("/treatments", controller.UpdateTreatmentRecord)
	r.DELETE("/treatments/:id", controller.DeleteTreatmentRecord)

	r.Run()
	fmt.Print("Test main.go")
}

func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {

			c.AbortWithStatus(204)

			return

		}

		c.Next()

	}

}
