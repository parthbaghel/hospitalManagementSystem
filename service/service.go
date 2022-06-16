package service

import (
	"fmt"
	"hostpitalProject/models"
	"strconv"
	"time"
)

func InsertDoc(item *models.DocInpt) models.Doctor {
	doc := models.Doctor{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      item.Name,
		ContactNo: item.Contact,
	}
	models.DB.Create(&doc)
	return doc
}

func InsertPat(item *models.PatInpt) (models.Patient, bool) {
	pat := models.Patient{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      item.Name,
		ContactNo: item.Contact,
		DoctorId:  item.DocId,
		Address:   item.Address,
	}
	q := models.DB.Create(&pat)
	if q.Error != nil {
		fmt.Println(q.Error)
		return pat, false
	}
	return pat, true
}

func PatchDoc(item *models.DocInptContact, DocId string) (models.Doctor, bool) {

	doc := models.Doctor{}
	id, _ := strconv.Atoi(DocId)
	q := models.DB.First(&doc, id)
	if q.Error != nil {
		fmt.Println(q.Error)
		return doc, false
	}
	if item.Contact != "" {
		doc.ContactNo = item.Contact
	}
	models.DB.Save(&doc)
	return doc, true
}

func PatchPat(item *models.PatPatchInpt, PatId string) (models.Patient, error) {
	pat := models.Patient{}
	id, _ := strconv.Atoi(PatId)
	q := models.DB.First(&pat, id)
	if q.Error != nil {
		fmt.Println(q.Error)
		return pat, fmt.Errorf("id not found")
	}
	if item.Contact != "" {
		pat.ContactNo = item.Contact
	}
	if item.DocId != 0 {
		pat.DoctorId = item.DocId
	}
	if item.Address != "" {
		pat.Address = item.Address
	}
	q = models.DB.Save(&pat)
	if q.Error != nil {
		fmt.Println(q.Error)
		return pat, fmt.Errorf("doctor not found")
	}
	return pat, nil
}
