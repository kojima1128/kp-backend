package model

type User struct {
	ID        string `json:"id"`
	CognitoID string `json:"cognitoId"`
	Name      string `json:"name"`
	TenantID  string `json:"tenantId"`
	SiteID    string `json:"siteId"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
}

type CreateUserInput struct {
	CognitoID string
	Name      string
	TenantID  string
	SiteID    string
	Role      string
	Email     string
}
