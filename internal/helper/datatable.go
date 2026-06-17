package helper

import (
	"fmt"
	"strings"

	"BackendPOS/internal/request"
	"BackendPOS/internal/response"

	"gorm.io/gorm"
)

// Paginate is a reusable generic function for DataTable-style pagination with search and sorting.
// Usage: Paginate[model.Merchant](db.Model(&model.Merchant{}), req, []string{"name", "code", "domain"})
func Paginate[T any](baseDB *gorm.DB, req request.DataTableRequest, searchColumns []string) ([]T, response.DataTableMeta, error) {
	var totalRecords int64
	var filteredRecords int64

	// Count total records (before any search filter)
	if err := baseDB.Count(&totalRecords).Error; err != nil {
		return nil, response.DataTableMeta{}, fmt.Errorf("failed to count total records: %w", err)
	}

	// Build search query
	searchDB := baseDB
	if req.Search != "" && len(searchColumns) > 0 {
		var conditions []string
		var args []interface{}
		for _, col := range searchColumns {
			conditions = append(conditions, fmt.Sprintf("%s LIKE ?", col))
			args = append(args, "%"+req.Search+"%")
		}
		searchDB = baseDB.Where(strings.Join(conditions, " OR "), args...)
	}

	// Count filtered records
	if err := searchDB.Count(&filteredRecords).Error; err != nil {
		return nil, response.DataTableMeta{}, fmt.Errorf("failed to count filtered records: %w", err)
	}

	// Validate OrderBy against whitelist to prevent SQL injection
	orderByColumn := "created_at"
	if req.OrderBy != "" {
		for _, col := range searchColumns {
			if req.OrderBy == col {
				orderByColumn = req.OrderBy
				break
			}
		}
		// also allow ordering by "created_at", "updated_at", "id"
		switch req.OrderBy {
		case "created_at", "updated_at", "id":
			orderByColumn = req.OrderBy
		}
	}

	orderDir := req.OrderDir
	if orderDir != "asc" {
		orderDir = "desc"
	}

	var items []T
	err := searchDB.
		Order(fmt.Sprintf("%s %s", orderByColumn, orderDir)).
		Offset(req.Offset()).
		Limit(req.Limit).
		Find(&items).Error
	if err != nil {
		return nil, response.DataTableMeta{}, fmt.Errorf("failed to fetch records: %w", err)
	}

	meta := response.DataTableMeta{
		Draw:            req.Draw,
		RecordsTotal:    totalRecords,
		RecordsFiltered: filteredRecords,
		Page:            req.Page,
		SearchValue:     req.Search,
		Limit:           req.Limit,
	}

	return items, meta, nil
}
