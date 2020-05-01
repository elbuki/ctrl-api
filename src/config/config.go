package config

import (
	"github.com/elbuki/ctrl-api/src/bcrypt"
	"github.com/elbuki/ctrl-api/src/control"
)

type Config struct {
	APIPort       string
	HashCost      int
	UsePassphrase bool
	Encryptor     bcrypt.Encryptor
	Controller    *control.Controller
}
