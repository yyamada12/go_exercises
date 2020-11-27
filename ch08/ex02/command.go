package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

type status struct {
	user          string
	addr          string // default: client addr, in passive mode: server addr
	dtype         int    // 0(default): ASCII, 1: Image
	isPassive     bool
	isLoggedIn    bool
	isTransfering bool
	transferDone  <-chan struct{}
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
	case "QUIT":
		if st.isTransfering {
			<-st.transferDone
		}
		fmt.Fprintln(c, "221 Logout")
		c.Close()
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
	case "MODE":
		if cmd.arg == "" {
			fmt.Fprintln(c, "501 Missing argument")
			return
		}
		if strings.ToUpper(cmd.arg) != "S" {
			fmt.Fprintln(c, "504 Please use S(tream) mode")
			return
		}
		fmt.Fprintln(c, "200 S OK")
	case "STRU":
		if cmd.arg == "" {
			fmt.Fprintln(c, "501 Missing argument")
			return
		}
		if strings.ToUpper(cmd.arg) != "F" {
			fmt.Fprintln(c, "504 Only F(ile) is supported")
			return
		}
		fmt.Fprintln(c, "200 F OK")
	case "RETR":
		if !st.isLoggedIn {
			fmt.Fprintln(c, "530 You aren't logged in")
			return
		}
		if cmd.arg == "" {
			fmt.Fprintln(c, "501 No file name")
			return
		}
		go retr(c, st, cmd.arg)
	case "STOR":
		if !st.isLoggedIn {
			fmt.Fprintln(c, "530 You aren't logged in")
			return
		}
		if cmd.arg == "" {
			fmt.Fprintln(c, "501 No file name")
			return
		}
		go store(c, st, cmd.arg)
	case "NOOP":
		fmt.Fprintln(c, "200 Zzz...")
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

func retr(c net.Conn, st *status, filename string) {
	conn, err := net.Dial("tcp", st.addr)
	if err != nil {
		fmt.Fprintln(c, "425 No data connection")
		return
	}
	fmt.Fprintln(c, "150 Accepted data connection")
	transferDone := make(chan struct{})
	st.transferDone = transferDone
	st.isTransfering = true
	defer conn.Close()
	defer func() { st.isTransfering = false }()
	defer func() { close(transferDone) }()

	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintln(c, "553 Can't open that file:", err)
		return
	}

	_, err = io.Copy(conn, file)
	if err != nil {
		fmt.Fprintln(c, "451", err)
		return
	}
	if err = file.Close(); err != nil {
		fmt.Fprintln(c, "451", err)
		return
	}
	fmt.Fprintln(c, "226 File successfully transferred")
}

func store(c net.Conn, st *status, filename string) {
	conn, err := net.Dial("tcp", st.addr)
	if err != nil {
		fmt.Fprintln(c, "425 No data connection")
		return
	}
	fmt.Fprintln(c, "150 Accepted data connection")
	transferDone := make(chan struct{})
	st.transferDone = transferDone
	st.isTransfering = true
	defer conn.Close()
	defer func() { st.isTransfering = false }()
	defer func() { close(transferDone) }()

	file, err := os.Create(filename)
	if err != nil {
		fmt.Fprintln(c, "553 Can't open that file:", err)
		return
	}

	_, err = io.Copy(file, conn)
	if err != nil {
		fmt.Fprintln(c, "451", err)
		return
	}
	if err = file.Close(); err != nil {
		fmt.Fprintln(c, "451", err)
		return
	}
	fmt.Fprintln(c, "226 File successfully transferred")
}
