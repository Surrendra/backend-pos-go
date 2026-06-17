package request

type CreateMerchantRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Address     string `json:"address" binding:"required"`
	PicName     string `json:"pic_name" binding:"required"`
	Domain      string `json:"domain" binding:"required"`
	Active      string `json:"active" binding:"required"`
	PicEmail    string `json:"pic_email" binding:"required"`
	PicAddress  string `json:"pic_address" binding:"required"`
	PicContact  string `json:"pic_contact" binding:"required"`
}

type UpdateMerchantRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Address     string `json:"address" binding:"required"`
	PicName     string `json:"pic_name" binding:"required"`
	Domain      string `json:"domain" binding:"required"`
	Active      string `json:"active" binding:"required"`
	PicEmail    string `json:"pic_email" binding:"required"`
	PicAddress  string `json:"pic_address" binding:"required"`
	PicContact  string `json:"pic_contact" binding:"required"`
}
