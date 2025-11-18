package main

import (
	"flag"
	"log"

	"github.com/blitzdb/blitz/config"
	"github.com/blitzdb/blitz/server"
)

func setUpFlags(){
	flag.StringVar(&config.Host,"host","0.0.0.0","host for blitzdb server")
	flag.IntVar(&config.Port,"port",9999,"port for blitzdb server")
	flag.Parse()
}

func main(){
	setUpFlags()
	log.Println("Starting BlitzDB 🎇")
	server.RunAsyncTCPServer()
}