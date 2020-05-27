package control_test

import (
	"testing"

	"github.com/elbuki/ctrl-api/src/control"

	pb "github.com/elbuki/ctrl-protobuf/src/golang"
)

type keyTestScenario struct {
	key    pb.Key
	throws bool
}

func TestKeys(t *testing.T) {
	table := []keyTestScenario{
		keyTestScenario{-1, true},
		keyTestScenario{pb.Key_ENTER, false},
	}

	for i, s := range table {
		_, err := control.TranslateProtoKey(s.key)
		if !s.throws && err != nil {
			t.Errorf("%d: could not translate: %v", i, err)
		}

		if s.throws && err == nil {
			t.Errorf("%d: expected error but it went fine", i)
		}
	}
}
