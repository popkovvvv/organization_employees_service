package entity

import (
	"practice/internal/domain"
	"time"
)

type Organization struct {
	ID        string
	Name      string
	Employees []Employee
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Employee struct {
	ID             string
	Name           string
	OrganizationID int64
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (e Employee) toDomain() domain.Employee {
	return domain.Employee{
		ID:             e.ID,
		Name:           e.Name,
		OrganizationID: e.OrganizationID,
		CreatedAt:      e.CreatedAt,
		UpdatedAt:      e.UpdatedAt,
	}
}

func (o Organization) ToDomain() *domain.Organization {
	return &domain.Organization{
		ID:        o.ID,
		Name:      o.Name,
		Employees: make([]domain.Employee, len(o.Employees)),
		CreatedAt: o.CreatedAt,
		UpdatedAt: o.UpdatedAt,
	}
}

func EmployeeDomainToEntity(e domain.Employee) Employee {
	return Employee{
		ID:             e.ID,
		Name:           e.Name,
		OrganizationID: e.OrganizationID,
		CreatedAt:      e.CreatedAt,
		UpdatedAt:      e.UpdatedAt,
	}
}

func OrganizationDomainToEntity(o *domain.Organization) Organization {
	employees := make([]Employee, len(o.Employees))
	for i, e := range o.Employees {
		employees[i] = EmployeeDomainToEntity(e)
	}

	return Organization{
		ID:        o.ID,
		Name:      o.Name,
		Employees: employees,
		CreatedAt: o.CreatedAt,
		UpdatedAt: o.UpdatedAt,
	}
}
