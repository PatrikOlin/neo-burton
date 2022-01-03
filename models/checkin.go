package models

import "time"

type Checkin struct {
	ID         string    `json:"id"`
	EmployeeID string    `json:"employee_id"`
	Time       time.Time `json:"checkin_time"`
}
