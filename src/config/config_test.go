package config_test

import (
	"flag"
	"os"
	"testing"

	"github.com/elbuki/ctrl-api/src/config"
	"github.com/elbuki/ctrl-api/src/control"
)

type flagTestScenario struct {
	args        []string
	shouldThrow bool
}

var conf *config.Config

func TestMain(m *testing.M) {
	conf = &config.Config{
		Controller: &control.Controller{},
	}

	os.Exit(m.Run())
}

func TestController(t *testing.T) {
	if err := conf.SetController(); err != nil {
		t.Fatalf("could not set the controller: %v", err)
	}
}

func TestFlagSetting(t *testing.T) {
	scenarios := []flagTestScenario{
		flagTestScenario{
			args:        []string{"-port=1234", "-cost=foo", "-P"},
			shouldThrow: true,
		},
		flagTestScenario{
			args:        []string{"-port=1234", "-cost=10", ""},
			shouldThrow: false,
		},
	}

	f := flag.NewFlagSet("args", flag.ContinueOnError)
	for _, s := range scenarios {
		err := conf.SetFlags(f, s.args...)
		if !s.shouldThrow && err != nil {
			t.Errorf("unexpected error received: %v", err)
		}

		if s.shouldThrow && err == nil {
			t.Error("expecting error but it went fine")
		}
	}
}
