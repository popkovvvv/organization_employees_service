package infra

import (
	"practice/internal/infra/datasource/repository"
	"practice/internal/usecase/gateway"
)

type serviceProvider struct {
	organizationRepository gateway.OrganizationRepository
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (sp *serviceProvider) OrganizationRepository() gateway.OrganizationRepository {
	if sp.organizationRepository == nil {
		sp.organizationRepository = repository.NewOrganizationRepository()
	}
	return sp.organizationRepository
}
