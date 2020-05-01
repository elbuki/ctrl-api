package config

import (
	"github.com/elbuki/ctrl-api/src/control"
)

func (c *Config) SetController() error {
	controller := new(control.Controller)

	if err := controller.Init(); err != nil {
		return err
	}

	c.Controller = controller

	return nil
}
