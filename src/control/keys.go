package control

import (
	"fmt"

	"github.com/micmonay/keybd_event"

	pb "github.com/elbuki/ctrl-protobuf/src/golang"
)

var keysDict = map[pb.Key]int{
	pb.Key_ENTER:       keybd_event.VK_ENTER,
	pb.Key_ESC:         keybd_event.VK_ESC,
	pb.Key_F11:         keybd_event.VK_F11,
	pb.Key_LEFT_ARROW:  keybd_event.VK_LEFT,
	pb.Key_RIGHT_ARROW: keybd_event.VK_RIGHT,
	pb.Key_ALT_TAB:     keybd_event.VK_TAB,
}

func TranslateProtoKey(key pb.Key) (int, error) {
	k, ok := keysDict[key]
	if !ok {
		return 0, fmt.Errorf("could not find the specified key: %v", key)
	}

	return k, nil
}
