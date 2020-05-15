package config_test

import (
	"errors"
	"flag"
	"os"
	"testing"

	"github.com/elbuki/ctrl-api/src/bcrypt"
	"github.com/elbuki/ctrl-api/src/config"
	"github.com/elbuki/ctrl-api/src/control"
)

type flagTestScenario struct {
	args        []string
	shouldThrow bool
}

type mockPasswordReader struct{}
type mockPasswordReaderNoBytes struct{}
type mockPasswordReaderError struct{}

func (mockPasswordReader) ReadPassword() ([]byte, error) {
	return []byte{0}, nil
}

func (mockPasswordReaderNoBytes) ReadPassword() ([]byte, error) {
	return []byte{}, nil
}

func (mockPasswordReaderError) ReadPassword() ([]byte, error) {
	return nil, errors.New("test error")
}

var conf *config.Config

func TestMain(m *testing.M) {
	conf = &config.Config{
		Controller: &control.Controller{},
		Encryptor:  bcrypt.Encryptor{},
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

	for _, s := range scenarios {
		f := flag.NewFlagSet("args", flag.ContinueOnError)
		err := conf.SetFlags(f, s.args...)
		if !s.shouldThrow && err != nil {
			t.Errorf("unexpected error received: %v", err)
		}

		if s.shouldThrow && err == nil {
			t.Error("expecting error but it went fine")
		}
	}
}

func TestPassphrase(t *testing.T) {
	if _, err := conf.GetPassphrase(&mockPasswordReader{}); err != nil {
		t.Errorf("unexpected error when getting passphrase: %v", err)
	}

	if _, err := conf.GetPassphrase(&mockPasswordReaderNoBytes{}); err != nil {
		t.Errorf("unexpected error when getting passphrase: %v", err)
	}

	if _, err := conf.GetPassphrase(&mockPasswordReaderError{}); err == nil {
		t.Error("expected error but got none when getting passphrase")
	}

	conf.Encryptor = bcrypt.NewEncryptor(50)

	if _, err := conf.GetPassphrase(&mockPasswordReader{}); err == nil {
		t.Error("expected error from changing the salt level got nothing")
	}
}

func TestOriginalPasswordReader(t *testing.T) {
	var reader config.StdinPasswordReader

	if _, err := reader.ReadPassword(); err == nil {
		t.Error("expected error but it went fine")
	}
}
