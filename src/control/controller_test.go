package control_test

import (
	"testing"

	"github.com/micmonay/keybd_event"

	"github.com/elbuki/ctrl-api/src/control"
)

func TestController(t *testing.T) {
	c := &control.Controller{}
	keys := []int{
		keybd_event.VK_A,
		keybd_event.VK_S,
		keybd_event.VK_D,
		keybd_event.VK_F,
	}

	if err := c.Init(); err != nil {
		t.Fatalf("could not init the controller: %v", err)
	}

	if err := c.SendKeys(keys...); err != nil {
		t.Fatalf("unable to send keys in the controller: %v", err)
	}
}
