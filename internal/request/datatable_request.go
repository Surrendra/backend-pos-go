package request

type DataTableRequest struct {
	Draw     int    `form:"draw"`
	Page     int    `form:"page"`
	Limit    int    `form:"limit"`
	Search   string `form:"search[term]"`
	OrderBy  string `form:"order_by"`
	OrderDir string `form:"order_dir"`
}

func (r *DataTableRequest) SetDefaults() {
	if r.Page <= 0 {
		r.Page = 1
	}
	if r.Limit <= 0 {
		r.Limit = 10
	}
	if r.Limit > 100 {
		r.Limit = 100
	}
	if r.OrderDir != "asc" && r.OrderDir != "desc" {
		r.OrderDir = "desc"
	}
}

func (r *DataTableRequest) Offset() int {
	return (r.Page - 1) * r.Limit
}
