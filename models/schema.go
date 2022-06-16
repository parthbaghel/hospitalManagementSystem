package models

import "time"

type Doctor struct {
	ID        int64 `gorm:"primary_key;auto_increment;not_null"` //string `gorm:"primary_key" gorm:"type:char(5)" `
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `json:"name" gorm:"type:varchar"`
	ContactNo string `json:"contact_no" gorm:"type:char(10)"`
}

/*
CREATE TABLE doctors (
  id SERIAL PRIMARY KEY,
  created_at timestamp,
  updated_at  timestamp,
  name varchar(255),
  contact_no varchar(10)
);

CREATE TABLE patients(
	id SERIAL PRIMARY KEY,
	created_at timestamp,
  	updated_at  timestamp,
  	name varchar(255),
  	contact_no varchar(10),
	doctor_id int,
	address varchar(255),
	FOREIGN KEY (doctor_id) REFERENCES doctors(id)
);

*/

type Patient struct {
	ID        int64 `gorm:"primary_key;auto_increment;not_null"` //string `gorm:"primary_key"  gorm:"type:char(5)" `
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `json:"name" gorm:"type:varchar"`
	ContactNo string `json:"contact_no" gorm:"type:char(10)"`
	DoctorId  int    `json:"doctor_id"`
	Address   string
}
