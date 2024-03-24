package employee

type CreateEmployeeRequest struct {
	Name   string   `json:"name"`
	OrgId  string   `json:"orgId"`
	Role   string   `json:"role"`
	Scopes []string `json:"scopes"`
}

type CreateEmployeeResponse struct {
	Id        string   `json:"id"`
	Name      string   `json:"name"`
	OrgId     string   `json:"orgId"`
	Role      string   `json:"role"`
	Scopes    []string `json:"scopes"`
	CreatedAt string   `json:"createdAt"`
}
