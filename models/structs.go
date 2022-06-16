package models

type DocInpt struct {
	Name    string `json:"name" binding:"required" `
	Contact string `json:"contact_no" binding:"required" `
}

type DocInptContact struct {
	Contact string `json:"contact_no" binding:"required" `
}

type PatInpt struct {
	Name    string `json:"name" binding:"required"`
	Contact string `json:"contact_no" binding:"required"`
	Address string `json:"address" binding:"required"`
	DocId   int    `json:"doctor_id" binding:"required"`
}

type PatPatchInpt struct {
	Contact string `json:"contact_no"`
	Address string `json:"address"`
	DocId   int    `json:"doctor_id"`
}
