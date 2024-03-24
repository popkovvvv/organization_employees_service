package organization

import (
	"github.com/google/uuid"
	"log"
	"practice/internal/domain"
)

type CreateOrganizationUseCaseRequest struct {
	Name string
}

type CreateOrganizationUseCaseRequestRepository interface {
	GetOrganizationByName(name string) *domain.Organization
	Save(organization *domain.Organization) (*domain.Organization, error)
}

type CreateOrganizationUseCase struct {
	Repository CreateOrganizationUseCaseRequestRepository
}

type CreateOrganizationUseCaseResponse struct {
	Organization *domain.Organization
}

func NewCreateOrganizationUseCase(repository *CreateOrganizationUseCaseRequestRepository) *CreateOrganizationUseCase {
	return &CreateOrganizationUseCase{
		Repository: *repository,
	}
}

func (uc *CreateOrganizationUseCase) Execute(request CreateOrganizationUseCaseRequest) (res *CreateOrganizationUseCaseResponse, err error) {
	organization := uc.Repository.GetOrganizationByName(request.Name)
	if organization != nil {
		return nil, domain.OrganizationAlreadyExists
	}

	organizationUUID, err := uuid.NewUUID()
	id := organizationUUID.String()

	if err != nil {
		log.Printf("ошибка генерации user UUID: %v\n", err)
		return nil, err
	}

	save, err := uc.Repository.Save(&domain.Organization{ID: id, Name: request.Name})
	if err != nil {
		return nil, err
	}
	res = &CreateOrganizationUseCaseResponse{Organization: save}

	return
}
