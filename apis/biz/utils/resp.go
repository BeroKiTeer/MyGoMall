package utils

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

// SendErrResponse  pack error response
func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {
	// todo edit custom code
	c.JSON(code, map[string]interface{}{
		"code":    code,
		"message": err.Error(),
		"data":    nil,
	})
}

// SendSuccessResponse  pack success response
func SendSuccessResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {
	// todo edit custom code
	c.JSON(code, map[string]interface{}{
		"code":    code,
		"message": "success",
		"data":    data,
	})
}
