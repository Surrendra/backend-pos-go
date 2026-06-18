package request

type CreateRoleRequest struct {
	Name        string   `json:"name" binding:"required"`
	Description *string  `json:"description"`
	Active      string   `json:"active" binding:"required"`
	Permissions []string `json:"permissions"`
}

type UpdateRoleRequest struct {
	Name        string   `json:"name" binding:"required"`
	Description *string  `json:"description"`
	Active      string   `json:"active" binding:"required"`
	Permissions []string `json:"permissions"`
}
