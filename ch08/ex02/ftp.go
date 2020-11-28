// FTP server based on RFC 959
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:10021")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		log.Printf("%s connected", conn.RemoteAddr())
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	input := bufio.NewScanner(c)
	st := new(status)

	fmt.Fprintln(c, "220 Welcome")
	for input.Scan() {
		cmd := parseInput(input.Text())
		handleCommand(cmd, c, st)
	}
	log.Printf("%s closed connection", c.RemoteAddr())
}
