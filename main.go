package main

import (
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/blitzdb/blitz/config"
	"github.com/blitzdb/blitz/server"
)

func setUpFlags() {
	flag.StringVar(&config.Host, "host", "0.0.0.0", "host for blitzdb server")
	flag.IntVar(&config.Port, "port", 9999, "port for blitzdb server")
	flag.Parse()
}

var MAX_ALLOWED_KEY int64

func main() {
	maximumAllowedKey := os.Getenv("MAXIMUM_KEY_ALLOWED")
	if maximumAllowedKey == "" {
		MAX_ALLOWED_KEY = 20000
	} else {
		k, err := strconv.Atoi(maximumAllowedKey)
		if err != nil {
			MAX_ALLOWED_KEY = 20000
		}
		MAX_ALLOWED_KEY = int64(k)
	}
	config.KeysLimit = int(MAX_ALLOWED_KEY)
	setUpFlags()
	log.Println("Starting BlitzDB 🎇")
	log.Println("Maximum Key Allowed is ",config.KeysLimit)
	server.RunAsyncTCPServer()
}
