package server

import (
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"

	"github.com/blitzdb/blitz/config"
	"github.com/blitzdb/blitz/core"
)

func respondError(err error, c net.Conn) {
	c.Write([]byte(fmt.Sprintf("-%s\r\n", err)))
}

func respond(cmd *core.RedisCmd, c net.Conn) {
	err := core.EvalAndRespond(cmd, c)
	if err != nil {
		respondError(err, c)
	}
}

func readCommand(c net.Conn) (*core.RedisCmd, error) {
	// TODO: max read in 1 shot is 512 bytes
	// To allow more than 512 bytes than do repeated read until error or EOF
	var buf []byte = make([]byte, 512)
	n, err := c.Read(buf[:])
	if err != nil {
		return nil, err
	}
	tokens, err := core.DecodeArrayString(buf[:n])
	if err != nil {
		return nil, err
	}

	return &core.RedisCmd{
		Cmd:  strings.ToUpper(tokens[0]),
		Args: tokens[1:],
	}, nil
}

func RunSyncTCPServer() {
	log.Println("starting a synchronous TCP server on ", config.Host, config.Port)

	var concurrentClient int = 0

	//listening to host:port

	listner, err := net.Listen("tcp", config.Host+":"+strconv.Itoa(config.Port))
	if err != nil {
		panic(err)
	}

	for {
		// blocking call waiting for the new client to connect
		conn, err := listner.Accept()
		if err != nil {
			panic(err)
		}

		concurrentClient++
		log.Println("Client connected with address:", conn.RemoteAddr(), "concurrent_clients", concurrentClient)

		for {
			cmd, err := readCommand(conn)
			if err != nil {
				conn.Close()
				concurrentClient--
				log.Println("client disconnected", conn.RemoteAddr(), "concurrent_clients", concurrentClient)
				if err == io.EOF {
					break
				}
				log.Println("error", err)
			}
			respond(cmd, conn)

		}
	}
}
