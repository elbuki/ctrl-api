package main

import (
	"log"
)

func main() {
	log.Printf("port: %s\n", apiPort)
	log.Printf("cost: %v\n", hashCost)
	log.Printf("passphrase: %s\n", passphraseHash)
}
