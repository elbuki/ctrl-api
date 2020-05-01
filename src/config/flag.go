package config

import "flag"

const (
	defaultAPIPort    = "9516"
	defaultHashCost   = 16
	setParaphraseFlag = "P"
)

func (c *Config) SetFlags() {
	flag.StringVar(
		&c.APIPort,
		"port",
		defaultAPIPort,
		"rpc serving port",
	)

	flag.IntVar(
		&c.HashCost,
		"cost",
		defaultHashCost,
		"hash salt cost for the passphrase",
	)

	flag.BoolVar(
		&c.UsePassphrase,
		setParaphraseFlag,
		false,
		"use a passphrase for client connections",
	)

	flag.Parse()
}
