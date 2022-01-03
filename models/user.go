package models

type User struct {
	EmployeeID    string `json:"employeeId"`
	Name          string `json:"name"`
	CurrentReport string `json:"currentReport"`
	NFCHex        string `json:"nfcHex"`
}
