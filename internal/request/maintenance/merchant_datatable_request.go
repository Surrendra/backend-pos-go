package request

import "BackendPOS/internal/request"

// MerchantDataTableRequest extends the generic DataTableRequest with
// merchant-specific exact-match filters.
type MerchantDataTableRequest struct {
	request.DataTableRequest

	// Active filters by merchant status. Accepted values: "Y" | "N". Empty = no filter.
	Active string `form:"search[active]"`
}
