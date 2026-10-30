package httputil

// From https://github.com/swaggo/swag/blob/master/example/celler/httputil/error.go

import "github.com/gin-gonic/gin"

// NewError example
func NewError(ctx *gin.Context, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	ctx.JSON(status, er)
}

// HTTPError example
type HTTPError struct {
	Code    int    `json:"code" example:"500"`
	Message string `json:"message" example:"Error processing data due to db issue"`
}
