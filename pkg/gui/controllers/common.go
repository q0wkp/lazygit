package controllers

import (
	"time"

	"github.com/jesseduffield/lazygit/pkg/gui/controllers/helpers"
)

// How long an operation may run before its progress indicator switches to a
// spinner; quick operations stay free of flicker.
const spinnerGracePeriod = 200 * time.Millisecond

type ControllerCommon struct {
	*helpers.HelperCommon
	IGetHelpers
}

type IGetHelpers interface {
	Helpers() *helpers.Helpers
}

func NewControllerCommon(
	c *helpers.HelperCommon,
	IGetHelpers IGetHelpers,
) *ControllerCommon {
	return &ControllerCommon{
		HelperCommon: c,
		IGetHelpers:  IGetHelpers,
	}
}
