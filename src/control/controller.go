package control

import (
	"fmt"
	"runtime"
	"time"

	"github.com/micmonay/keybd_event"
)

type Controller struct {
	kb keybd_event.KeyBonding
}

func (c *Controller) Init() error {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		return fmt.Errorf("could not initialize key bonding: %v", err)
	}

	// For Linux is very important to wait 2 seconds
	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}

	c.kb = kb

	return nil
}

func (c *Controller) SendKeys(keys ...int) error {
	c.kb.SetKeys(keys...)

	return c.launch()
}

func (c *Controller) launch() (err error) {
	if err = c.kb.Launching(); err != nil {
		return fmt.Errorf("could not launch sequence of keys: %v", err)
	}

	return
}
