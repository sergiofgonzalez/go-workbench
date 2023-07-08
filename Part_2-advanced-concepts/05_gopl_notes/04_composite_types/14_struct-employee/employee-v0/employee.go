package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID 				int
	Name 			string
	Address 	string
	DoB 			time.Time
	Position 	string
	Salary 		int
	ManagerID int
}

func EmployeeByID(id int) Employee {
	return Employee{
		ID: 1,
		Name: "Employee Name",
		Address: "Employee Address",
		DoB: time.Date(1974, 05, 27, 0, 0, 0, 0, time.Local),
		Position: "Employee Position",
		Salary: 100,
		ManagerID: 0,
	}
}

func main() {
	employee := EmployeeByID(1)
	employee.Salary += 5000
	fmt.Println(employee)

	EmployeeByID(1).Salary = 0
}