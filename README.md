## HospitalManagementSystem

This is a hospital management system

How to use:
```
$ go run main.go
```


### Database schema
- Doctor
```
id 			SERIAL PRIMARY KEY  
created_at 	timestamp  
updated_at  timestamp
name 		varchar(255)
contact_no 	varchar(10)
```
- Patient
```
id 			SERIAL PRIMARY KEY  
created_at 	timestamp 
updated_at  timestamp 
name 		varchar(255) 
contact_no 	varchar(10)
doctor_id 	int  
address 	varchar(255)  
FOREIGN KEY (doctor_id) REFERENCES doctors(id)
```

### APIs

The following are the API supported by the service.

#### POST /doctor/
Create a doctor a new entry of doctor in database
**Request** 
```
{		  
	"name": ”abcd”,  
	"contact_no": “1234567890”  
}
```
**Response**
```
{  
"id": 1,
“created_at”: "2022-06-16T13:15:57.485725Z",
“updated_at”: "2022-06-16T13:15:57.485725Z",  
“name”: “abcd”,
“contact_no” : “1234567890”  
}
```
#### GET /doctor/:id
Get a doctor entry from database
**Response**
```
{  
"id": 1,
“created_at”: "2022-06-16T13:15:57.485725Z",
“updated_at”: "2022-06-16T13:15:57.485725Z",  
“name”: “abcd”,
“contact_no” : “1234567890”  
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
"id": 1,
“created_at”: 1655146724,
“updated_at”: 1655146724,  
“name”: “abcd”,
“contact_no” : “1234567899”
}
```



