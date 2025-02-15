package utils

import (
	"github.com/cloudwego/hertz/pkg/app"
)

func BindJson(c *app.RequestContext, req interface{}) error {
	if err := c.Bind(&req); err != nil {
		c.JSON(400, map[string]string{"error": "Invalid request body"})
		return err
	}
	return nil
}
