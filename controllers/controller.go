package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hostpitalProject/models"
	"hostpitalProject/service"
	"strconv"
)

func DocPost(c *gin.Context) {
	var inp models.DocInpt
	if err := c.ShouldBindJSON(&inp); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input!"})
		return
	}
	doc := service.InsertDoc(&inp)
	c.JSON(200, gin.H{
		"response": doc,
	})
	return
}

func PatPost(c *gin.Context) {
	var inp models.PatInpt
	if err := c.ShouldBindJSON(&inp); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input!"})
		return
	}
	pat, bol := service.InsertPat(&inp)
	if bol == false {
		c.JSON(400, gin.H{"error": "Doctor Id not present!"})
		return
	}
	c.JSON(200, gin.H{
		"response": pat,
	})
	return
}

func DocGet(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var doc models.Doctor
	q := models.DB.First(&doc, id)
	fmt.Println(q.Error)
	if q.Error != nil {
		fmt.Println(q.Error)
		c.JSON(400, gin.H{"error": "ID not present"})
		return
	}
	c.JSON(200, gin.H{"response": doc})
}

func PatGet(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	pat := models.Patient{}
	q := models.DB.First(&pat, id)
	if q.Error != nil {
		fmt.Println(q.Error)
		c.JSON(400, gin.H{"error": "ID not present"})
		return
	}
	c.JSON(200, gin.H{"response": pat})
}

func DocPatch(c *gin.Context) {
	id := c.Param("id")
	var inp models.DocInptContact
	if err := c.ShouldBindJSON(&inp); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input!"})
		return
	}
	doc, bol := service.PatchDoc(&inp, id)
	if bol == false {
		c.JSON(400, gin.H{"error": "Id not found!"})
		return
	}
	c.JSON(200, gin.H{
		"response": doc,
	})
	return
}

func PatPatch(c *gin.Context) {
	id := c.Param("id")
	var inp models.PatPatchInpt
	if err := c.ShouldBindJSON(&inp); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input!"})
		return
	}
	pat, err := service.PatchPat(&inp, id)
	if err == nil {
		c.JSON(200, gin.H{
			"response": pat,
		})
		return
	}
	fmt.Println(err)
	if err.Error() == "id not found" {
		c.JSON(400, gin.H{"error": "Id not found!"})
		return
	}

	if err.Error() == "doctor not found" {
		c.JSON(400, gin.H{"error": "Doctor not found!"})
		return
	}
}

func GetDoctorPatientList(c *gin.Context) {
	/*
		id, _ := strconv.Atoi(c.Param("id"))
		var patients []models.Patient
		q := models.DB.Where(map[string]interface{}{"doctor_id": id}).Find(patients)
		if q.Error != nil {
			fmt.Println(q.Error)
		}
	*/

	id := c.Param("id")
	patients, bol := service.GetDocPatList(id)
	if bol == false {
		c.JSON(400, gin.H{"error": "Doctor not found!"})
		return
	}
	c.JSON(200, gin.H{"response": patients})
}
