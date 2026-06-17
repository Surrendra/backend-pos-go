package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Success(ctx *gin.Context, message string, data interface{}) {
	ctx.JSON(200, APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func SuccessWithMeta(ctx *gin.Context, message string, data, meta interface{}) {
	ctx.JSON(200, APIResponse{
		Success: true,
		Message: message,
		Data:    data,
		Meta:    meta,
	})
}

func Error(ctx *gin.Context, status int, message string, errors interface{}) {
	ctx.JSON(status, APIResponse{
		Success: false,
		Message: message,
		Errors:  errors,
	})
}

func FailedBinding(ctx *gin.Context, message string, errors interface{}) {
	ctx.JSON(http.StatusBadRequest, APIResponse{
		Success: false,
		Message: message,
		Errors:  errors,
	})
}

func BadRequest(ctx *gin.Context, message string, errors interface{}) {
	ctx.JSON(400, APIResponse{
		Success: false,
		Message: message,
		Errors:  errors,
	})
}

func SuccessDatatable(ctx *gin.Context, message string, data interface{}, meta DataTableMeta) {
	ctx.JSON(http.StatusOK, DataTableResponse{
		Success: true,
		Message: message,
		Meta:    meta,
		Data:    data,
	})
}

func ValidationError(ctx *gin.Context, message string, err interface{}) {
	// use ErrorMessage struct to format validation errors

	// if not a validator.ValidationErrors, just return as-is
	verrs, ok := err.(validator.ValidationErrors)
	if !ok {
		ctx.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Message: message,
			Errors:  err,
		})
		return
	}

	// build []ErrorMessage list
	var formatted []ErrorMessage
	for _, fe := range verrs {
		field := jsonFieldFromError(fe)

		msg := "invalid " + field
		switch fe.Tag() {
		case "required":
			msg = field + " is required"
		case "email":
			msg = field + " must be a valid email address"
		case "min":
			msg = field + " must be at least " + fe.Param()
		case "max":
			msg = field + " must be at most " + fe.Param()
		}

		formatted = append(formatted, ErrorMessage{
			Field:   field,
			Message: msg,
		})
	}

	ctx.JSON(http.StatusBadRequest, APIResponse{
		Success: false,
		Message: message,
		Errors:  formatted,
	})
}

// jsonFieldFromError returns the json tag (snake_case) if available, otherwise the field name.
func jsonFieldFromError(fe validator.FieldError) string {
	if jsonTag := fe.StructField(); jsonTag != "" {
		return jsonTag
	}
	return fe.Field()
}

type DataTableMeta struct {
	Draw            int    `json:"draw"`
	RecordsTotal    int64  `json:"records_total"`
	RecordsFiltered int64  `json:"records_filtered"`
	Page            int    `json:"page"`
	SearchValue     string `json:"search_value"`
	Limit           int    `json:"limit"`
}

type DataTableResponse struct {
	Success bool          `json:"success"`
	Message string        `json:"message"`
	Meta    DataTableMeta `json:"meta"`
	Data    interface{}   `json:"data"`
}

type ErrorMessage struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
