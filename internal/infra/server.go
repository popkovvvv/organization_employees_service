package infra

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"practice/internal/infra/api/employee"
	"practice/internal/infra/api/organization"
)

type Server struct {
	serviceProvider *serviceProvider
}

func StartServer() {
	server := &Server{
		serviceProvider: newServiceProvider(),
	}

	r := gin.Default()

	r.POST("/organization", func(c *gin.Context) {
		organization.CreateOrganization(c, server.serviceProvider.OrganizationRepository())
	})
	r.POST("/organization/employee", func(c *gin.Context) {
		employee.CreateEmployee(c, server.serviceProvider.OrganizationRepository())
	})

	err := r.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}
