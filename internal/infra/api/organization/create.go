package organization

type CreateOrganizationRequest struct {
	Name string `json:"name"`
}

type CreateOrganizationResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}
