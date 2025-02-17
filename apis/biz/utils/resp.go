package utils

import (
	"context"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	"github.com/cloudwego/hertz/pkg/app"
)

// SendErrResponse  pack error response
func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {
	// todo edit custom code
	data := map[string]interface{}{
		"code":    code,
		"message": consts.StatusMessage(code),
		"data":    err,
	}
	c.JSON(code, data)
}

// SendSuccessResponse  pack success response
func SendSuccessResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {
	// todo edit custom code
	data = map[string]interface{}{
		"code":    code,
		"message": consts.StatusMessage(code),
		"data":    data,
	}
	c.JSON(code, data)
}
