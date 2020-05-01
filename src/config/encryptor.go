package config

import "github.com/elbuki/ctrl-api/src/bcrypt"

func (c *Config) SetEncryptor() (encrypt bcrypt.Encryptor) {
	encrypt.SetCost(c.HashCost)

	return
}
