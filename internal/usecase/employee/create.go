package employee

import (
	"github.com/google/uuid"
	"practice/internal/domain"
)

type CreateEmployeeUseCaseRequest struct {
	Name  string
	OrgId string
}

type CreateEmployeeUseCaseRequestRepository interface {
	GetOrganizationById(id string) *domain.Organization
	Save(organization *domain.Organization) (*domain.Organization, error)
}

type CreateEmployeeUseCase struct {
	Repository CreateEmployeeUseCaseRequestRepository
}

type CreateEmployeeUseCaseResponse struct {
	Organization *domain.Organization
}

func NewCreateEmployeeUseCase(repository *CreateEmployeeUseCaseRequestRepository) *CreateEmployeeUseCase {
	return &CreateEmployeeUseCase{
		Repository: *repository,
	}
}

func (uc *CreateEmployeeUseCase) Execute(request CreateEmployeeUseCaseRequest) (*CreateEmployeeUseCaseResponse, error) {
	organization := uc.Repository.GetOrganizationById(request.OrgId)
	if organization == nil {
		return nil, domain.OrganizationNotFound
	}

	employeeUUID, err := uuid.NewUUID()
	id := employeeUUID.String()

	err = organization.AddEmployee(domain.Employee{ID: id, Name: request.Name})
	res := &CreateEmployeeUseCaseResponse{Organization: organization}

	return res, err
}
