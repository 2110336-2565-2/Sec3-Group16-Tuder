package schemas

import "time"

type SchemaRegister struct {
	TutorName		string   `json:"TutorName,omitempty"`
	CourseName		string	`json:"CourseName,omitempty"`
	Subject			string  `json:"Subject,omitempty"`
	DayAvailable	[]bool  `json:"DayAvailable,omitempty"`

	
	//TimeAvailableStart		time.Time   `json:"TimeAvailableStart,omitempty"`
	//TimeAvailableEnd		time.Time   `json:"TimeAvailableEnd,omitempty"`
}
