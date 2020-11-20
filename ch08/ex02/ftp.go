// Clock is a TCP server that periodically writes the time.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:10021")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle connections concurrently
	}
}

type status struct {
	user      string
	addr      string // default: client addr, in passive mode: server addr
	isPassive bool
}

func handleConn(c net.Conn) {
	defer c.Close()
	input := bufio.NewScanner(c)
	st := new(status)

	for input.Scan() {
		cmd := parseInput(input.Text())
		handleCommand(cmd, c, st)
	}

}

type command struct {
	code string
	arg  string
}

func parseInput(input string) command {
	spaceIndex := strings.Index(input, " ")
	if spaceIndex < 0 {
		// no args
		return command{input, ""}
	}
	return command{input[:spaceIndex], input[spaceIndex+1:]}

}

func handleCommand(cmd command, c net.Conn, st *status) {

	switch strings.ToUpper(cmd.code) {
	case "USER":
		if cmd.arg == "" {
			fmt.Fprintln(c, "530 This is a private system - No anonymous login")
			return
		}
		st.user = cmd.arg
		fmt.Fprintln(c, "230 OK. Current directory is /")
	}
}
