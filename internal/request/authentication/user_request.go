package request

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
}

type CreateUserRequest struct {
	Name        string   `json:"name" binding:"required"`
	Email       string   `json:"email" binding:"required,email"`
	Username    string   `json:"username" binding:"required"`
	Password    string   `json:"password" binding:"required"`
	Address     string   `json:"address" binding:"required"`
	Phone       string   `json:"phone" binding:"required"`
	RoleCodes   []string `json:"role_codes" binding:"required"`
	MerchantIDs []uint64 `json:"merchant_id" binding:"required"`
}

type UpdateUserRequest struct {
	Name      string   `json:"name" binding:"required"`
	Email     string   `json:"email" binding:"required,email"`
	Username  string   `json:"username" binding:"required"`
	Password  string   `json:"password"`
	Address   string   `json:"address"`
	Phone     string   `json:"phone"`
	Active    string   `json:"active" binding:"required"`
	Status    string   `json:"status" binding:"required"`
	RoleCodes []string `json:"role_codes"`
}
