package organization

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"practice/internal/usecase/gateway"
	"practice/internal/usecase/organization"
)

func CreateOrganization(c *gin.Context, repository gateway.OrganizationRepository) {
	// Получаем данные из тела запроса
	var requestBody CreateOrganizationRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var repo organization.CreateOrganizationUseCaseRequestRepository = repository
	useCase := organization.NewCreateOrganizationUseCase(&repo)
	useCaseResult, err := useCase.Execute(organization.CreateOrganizationUseCaseRequest{Name: requestBody.Name})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Отправляем ответ клиенту
	c.JSON(http.StatusOK, gin.H{"organization": useCaseResult})
}
