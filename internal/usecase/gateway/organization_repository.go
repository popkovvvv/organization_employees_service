package gateway

import (
	"practice/internal/usecase/employee"
	"practice/internal/usecase/organization"
)

type OrganizationRepository interface {
	organization.CreateOrganizationUseCaseRequestRepository
	employee.CreateEmployeeUseCaseRequestRepository
}
