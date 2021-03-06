package config

import (
	"flag"
	"fmt"
	"os"
)

const (
	defaultAPIPort    = "9516"
	defaultHashCost   = 16
	setParaphraseFlag = "P"
)

func (c *Config) SetFlags(f *flag.FlagSet, args ...string) error {
	var err error

	if f == nil {
		f = flag.NewFlagSet("args", flag.ContinueOnError)
	}

	f.StringVar(
		&c.APIPort,
		"port",
		defaultAPIPort,
		"rpc serving port",
	)

	f.IntVar(
		&c.HashCost,
		"cost",
		defaultHashCost,
		"hash salt cost for the passphrase",
	)

	f.BoolVar(
		&c.UsePassphrase,
		setParaphraseFlag,
		false,
		"use a passphrase for client connections",
	)

	if len(args) > 0 {
		err = f.Parse(args)
	} else {
		err = f.Parse(os.Args[1:])
	}

	if err != nil {
		return fmt.Errorf("could not parse the flags: %v", err)
	}

	return nil
}
