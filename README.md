## HospitalManagementSystem

This is a hospital management system

How to use:
```
$ go run main.go
```


### Database schema
- Doctor
```
id 		SERIAL PRIMARY KEY  
created_at 	timestamp  
updated_at  	timestamp
name 		varchar(255)
contact_no 	varchar(10)
```
- Patient
```
id 		SERIAL PRIMARY KEY  
created_at 	timestamp 
updated_at  	timestamp 
name 		varchar(255) 
contact_no 	varchar(10)
doctor_id 	int  
address 	varchar(255)  
FOREIGN KEY (doctor_id) REFERENCES doctors(id)
```

### APIs

The following are the API supported by the service.

#### POST /doctor/
Create a new entry of doctor in database
**Request** 
```
{		  
	"name": 	”abcd”,  
	"contact_no": 	“1234567890”  
}
```
**Response**
```
{  
	"id": 		1,
	“created_at”: 	"2022-06-16T13:15:57.485725Z",
	“updated_at”: 	"2022-06-16T13:15:57.485725Z",  
	“name”: 	“abcd”,
	“contact_no” : 	“1234567890”  
}
```
#### GET /doctor/:id
Get a doctor entry from database
**Response**
```
{  
	"id": 		1,
	“created_at”: 	"2022-06-16T13:15:57.485725Z",
	“updated_at”: 	"2022-06-16T13:15:57.485725Z",  
	“name”: 	“abcd”,
	“contact_no” : 	“1234567890”  
}
```
#### PATCH /doctor/:id
Only, “contact_no” is allowed to be updated for doctor.
**Request**  
```
{
	“contact_no”: “1234567899”  
}
```
  

**Response**  
```
{  
	"id": 		1,
	“created_at”: 	"2022-06-16T13:15:57.485725Z",
	“updated_at”: 	"2022-06-16T13:15:57.485725Z",  
	“name”: 	“abcd”,
	“contact_no” : 	“1234567899”
}
```



#### POST /patient/  
Create a new entry of patient in database
**Request**
```  
{  
	"name": 	”xyz”,  
	"contact_no": 	“1234567890”,
	“address”: 	“H No.- 45/W, Street 5, Moti Bagh, West Bengal, India”
	“doctor_id”: 	3 
}  
```
  
**Response**
```  
{  
	"id": 		"20001",
	“created_at”: 	"2022-06-16T13:15:57.485725Z",
	“updated_at”: 	"2022-06-16T13:15:57.485725Z",  
	“name”: 	“xyz”,
	“contact_no” : 	“1234567890”,
	“address”: 	“H No.- 45/W, Street 5, Moti Bagh, West Bengal, India”
	“doctor_id”: 	3
}
```
#### GET /patient/:id  
**Response**
```
{  
	"id": 		1,
	“created_at”: 	"2022-06-16T13:15:57.485725Z",
	“updated_at”: 	"2022-06-16T13:15:57.485725Z",  
	“name”: 	“xyz”,
	“contact_no” : 	“1234567890”,
	“address”: 	“H No.- 45/W, Street 5, Moti Bagh, West Bengal, India”
	“doctor_id”: 	1 
}
```
#### PATCH /patient/:id  
Only, “contact_no”, “address” and “doctor_id” is allowed to be updated for a patient.
**Request**  
```
{
	“doctor_id”:  4  
}
```

  

**Response**

```  
{  
	"id": 1,
	“created_at”: "2022-06-16T13:15:57.485725Z",
	“updated_at”: "2022-06-16T13:15:57.485725Z",  
	“name”: “pqrs”,
	“contact_no” : “1234567890”,
	“address”: “H No.- 45/W, Street 5, Moti Bagh, West Bengal, India”
	“doctor_id”: 4
}
```
  
#### GET /doctorPatientList/:id
id: here id is the doctor_id.
This should return the list of patients being monitored by the same doctor.
**Response**  
```
[
{  
	"id": 		1,
	“created_at”: 	"2022-06-16T13:15:57.485725Z",
	“updated_at”: 	"2022-06-16T13:15:57.485725Z",  
	“name”: 	“xyz”,
	“contact_no” : 	“1234567890”,
	“address”: 	“H No.- 45/W, Street 5, Moti Bagh, West Bengal, India”
	“doctor_id”: 	3  
},

{  
	"id": 		2,
	“created_at”: 	"2022-06-16T13:15:57.485725Z",
	“updated_at”: 	"2022-06-16T13:15:57.485725Z",  
	“name”: 	“abc”,
	“contact_no” : 	“1234567891”,
	“address”: 	“H No.- 46/W, Street 5, Moti Bagh, West Bengal, India”
	“doctor_id”: 	3  
}
]
```
