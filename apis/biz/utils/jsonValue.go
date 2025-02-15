package utils

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func bindJson(ctx context.Context, c *app.RequestContext, req interface{}) error {
	if err := c.Bind(&req); err != nil {
		c.JSON(400, map[string]string{"error": "Invalid request body"})
		SendErrResponse(ctx, c, consts.StatusOK, err)
		return err
	}
	return nil
}
