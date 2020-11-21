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
	user       string
	addr       string // default: client addr, in passive mode: server addr
	dtype      int    // 0(default): ASCII, 1: Image
	isPassive  bool
	isLoggedIn bool
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
		st.isLoggedIn = true
		fmt.Fprintln(c, "230 OK. Current directory is /")
	case "PORT":
		if !st.isLoggedIn {
			fmt.Fprintln(c, "530 You aren't logged in")
			return
		}
		if st.isPassive {
			fmt.Fprintln(c, "501 Active mode is disabled")
			return
		}
		addr, err := parsePort(cmd.arg)
		if err != nil {
			fmt.Fprintln(c, "501", err)
			return
		}
		st.addr = addr
		fmt.Fprintln(c, "200 PORT command successful")
	case "TYPE":
		if cmd.arg == "" {
			fmt.Fprintln(c, "501-Missing argument")
			fmt.Fprintln(c, "501 TYPE is now", typeStringify(st.dtype))
			return
		}
		dtype, err := parseType(cmd.arg)
		if err != nil {
			fmt.Fprintln(c, "504-"+err.Error())
			fmt.Fprintln(c, "504 TYPE is now", typeStringify(st.dtype))
			return
		}
		st.dtype = dtype
		fmt.Fprintln(c, "200-TYPE command successful")
		fmt.Fprintln(c, "200 TYPE is now", typeStringify(st.dtype))
	}
}

func parsePort(arg string) (string, error) {
	var h1, h2, h3, h4, p1, p2 uint
	n, _ := fmt.Sscanf(arg, "%d,%d,%d,%d,%d,%d", &h1, &h2, &h3, &h4, &p1, &p2)
	if n != 6 || h1 > 255 || h2 > 255 || h3 > 255 || h4 > 255 || p1 > 255 || p2 > 255 || (h1|h2|h3|h4) == 0 || (p1|p2) == 0 {
		return "", fmt.Errorf("Syntax error in IP address")
	}
	p := p1<<8 | p2
	if p < 1024 {
		return "", fmt.Errorf("Sorry, but I won't connect to ports < 1024")
	}
	return fmt.Sprintf("%d.%d.%d.%d:%d", h1, h2, h3, h4, p), nil
}

func parseType(arg string) (int, error) {
	switch strings.ToUpper(arg) {
	case "A":
		return 0, nil
	case "I":
		return 1, nil
	default:
		return -1, fmt.Errorf("Unknown Type: %s", arg)
	}
}

func typeStringify(dtype int) string {
	switch dtype {
	case 1:
		return "8-bit binary"
	default:
		return "ASCII"
	}
}
