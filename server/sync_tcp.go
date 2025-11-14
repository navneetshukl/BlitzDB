package server

import (
	"io"
	"log"
	"net"
	"strconv"

	"github.com/blitzdb/blitz/config"
)

func respond(cmd string, c net.Conn) error {
	if _, err := c.Write([]byte(cmd)); err != nil {
		return err
	}
	return nil
}

func readCommand(c net.Conn)(string,error){
	// TODO: max read in 1 shot is 512 bytes
	// To allow more than 512 bytes than do repeated read until error or EOF
	var buf []byte=make([]byte, 512)
	n,err:=c.Read(buf[:])
	if err!=nil{
		return "",err
	}
	return string(buf[:n]),nil
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
			log.Println("command", cmd)
			if err = respond(cmd, conn); err != nil {
				log.Print("err write:", err)
			}

		}
	}
}
