package domain

import "github.com/pkg/errors"

var (
	OrganizationAlreadyExists = errors.New("organization already exists")
	OrganizationNotFound      = errors.New("organization not found")
	EmployeeNotFound          = errors.New("employee not found")
	EmployeeAlreadyExists     = errors.New("employee already exists")
)
