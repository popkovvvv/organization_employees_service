package domain

import (
	"log"
	"time"
)

type Organization struct {
	ID        string
	Name      string
	Employees []Employee
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (o *Organization) AddEmployee(e Employee) (err error) {
	for _, employee := range o.Employees {
		if employee.Name == e.Name {
			log.Println("Employee already exists with name:", e.Name)
			err = EmployeeAlreadyExists
		}
	}

	o.Employees = append(o.Employees, e)
	log.Println("Added employee:", e.Name)
	return
}

func (o *Organization) removeEmployee(name string) error {
	for i, employee := range o.Employees {
		if employee.Name == name {
			o.Employees = append(o.Employees[:i], o.Employees[i+1:]...)
			log.Println("Removed employee:", name)
			return nil
		}
	}

	log.Println("Employee not found:", name)
	return EmployeeNotFound
}

func (o *Organization) updateEmployee(e *Employee) error {
	for i, employee := range o.Employees {
		if employee.Name == e.Name {
			o.Employees[i] = *e
			log.Println("Updated employee:", e.Name)
			return nil
		}
	}

	log.Println("Employee not found:", e.Name)
	return EmployeeNotFound
}

type Employee struct {
	ID             string
	Name           string
	OrganizationID int64
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
