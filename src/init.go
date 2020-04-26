package main

import (
	"flag"
	"os"
)

var (
	API_PORT = "9516"
)

func init() {
	flag.StringVar(
		&API_PORT,
		"port",
		EnvString("PORT", API_PORT),
		"api's rpc serving port",
	)
}

func EnvString(key, defaultValue string) string {
	v, ok := os.LookupEnv(key)

	if !ok {
		return defaultValue
	}

	return v
}
