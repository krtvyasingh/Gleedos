package netutil

type CongestionController struct {
	WindowSize int
	MinWindow  int
	MaxWindow  int
}

func NewCongestionController() *CongestionController {
	return &CongestionController{WindowSize: 4, MinWindow: 1, MaxWindow: 32}
}

func (c *CongestionController) OnSuccess() {
	if c.WindowSize < c.MaxWindow {
		c.WindowSize++
	}
}

func (c *CongestionController) OnDrop() {
	c.WindowSize /= 2
	if c.WindowSize < c.MinWindow {
		c.WindowSize = c.MinWindow
	}
}
