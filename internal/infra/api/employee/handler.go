package employee

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"practice/internal/usecase/employee"
	"practice/internal/usecase/gateway"
)

func CreateEmployee(c *gin.Context, repository gateway.OrganizationRepository) {
	// Получаем данные из тела запроса
	var requestBody CreateEmployeeRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var repo employee.CreateEmployeeUseCaseRequestRepository = repository
	useCase := employee.NewCreateEmployeeUseCase(&repo)
	execute, err := useCase.Execute(employee.CreateEmployeeUseCaseRequest{Name: requestBody.Name, OrgId: requestBody.OrgId})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Отправляем ответ клиенту
	c.JSON(http.StatusOK, gin.H{"organization": execute.Organization})
}
