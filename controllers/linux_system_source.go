package controllers

import (
	"encoding/json"
	"errors"

	"github.com/mengjunwei/go_api_project/models"
	"github.com/mengjunwei/go_api_project/services/linux_system_source"
)

type LinuxSystemSourceSetController struct {
	ApiJsonBaseController
}

func (c *LinuxSystemSourceSetController) Post() {
	body := c.Ctx.Input.RequestBody
	if len(body) > MaxRequestBodyLength {
		c.SetError(errors.New(ErrBodyLimit))
	}
	params := &models.SystemSetDTO{}
	if err := json.Unmarshal(body, params); err != nil {
		c.SetError(err)
	}
	service := &linux_system_source.Service{}
	switch params.Type {
	case models.SetTypeMem:
		res, err := service.SetMemory(params)
		if err != nil {
			c.SetError(err)
		}
		c.SetData(res)
	}
	c.SetError(errors.New("参数错误"))
}
