package repository

import (
	"practice/internal/domain"
	"practice/internal/infra/datasource/entity"
	"practice/internal/usecase/gateway"
	"sync"
)

type OrganizationRepository struct {
	gateway.OrganizationRepository
	OrganizationsMap map[string]entity.Organization
	mutex            sync.Mutex
}

func NewOrganizationRepository() *OrganizationRepository {
	return &OrganizationRepository{OrganizationsMap: make(map[string]entity.Organization)}
}

func (r *OrganizationRepository) Save(organization *domain.Organization) (*domain.Organization, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.OrganizationsMap[organization.ID] = entity.OrganizationDomainToEntity(organization)
	savedOrganization := r.OrganizationsMap[organization.ID]
	return savedOrganization.ToDomain(), nil
}

func (r *OrganizationRepository) GetOrganizationById(id string) *domain.Organization {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	organization, ok := r.OrganizationsMap[id]
	if !ok {
		return nil
	}
	return organization.ToDomain()
}

func (r *OrganizationRepository) GetOrganizationByName(name string) (organization *domain.Organization) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	for _, o := range r.OrganizationsMap {
		if o.Name == name {
			organization = o.ToDomain()
		}
	}

	return
}
