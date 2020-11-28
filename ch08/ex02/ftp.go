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
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	fmt.Fprintln(c, "220 Welcome")

	client := c.RemoteAddr()
	log.Printf("%s connected", client)

	scanner := bufio.NewScanner(c)
	st := new(status)

	for scanner.Scan() {
		input := scanner.Text()
		log.Printf("%s: %s", client, input)
		cmd := parseInput(input)
		handleCommand(cmd, c, st)
	}
	log.Printf("%s closed connection", c.RemoteAddr())
}
