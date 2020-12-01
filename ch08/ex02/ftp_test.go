// Clock is a TCP server that periodically writes the time.
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"testing"
)

func Test_Login(t *testing.T) {
	tests := []struct {
		name string
		cmd  string
		arg  string
		want string
	}{
		{"no user name", "USER", "", "530 This is a private system - No anonymous login"},
		{"test user login", "USER", "test", "230 OK. Current directory is /"},
		{"PORT before login", "PORT", "127,0,0,1,4,20", "530 You aren't logged in"},
		{"RETR before login", "RETR", "memo.md", "530 You aren't logged in"},
		{"STOR before login", "STOR", "memo.md", "530 You aren't logged in"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server, client := net.Pipe()
			go func() {
				handleConn(server)
				server.Close()
			}()
			input := bufio.NewScanner(client)
			input.Scan() // 220 Welcome

			// Login
			fmt.Fprintln(client, tt.cmd, tt.arg)
			input.Scan()
			got := input.Text()
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
			client.Close()
		})
	}
}

func Test_Port(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{"no args", "", "501 Syntax error in IP address"},
		{"invalid host ", "256,255,255,1,10,1", "501 Syntax error in IP address"},
		{"invalid port ", "127,0,0,1,256,1", "501 Syntax error in IP address"},
		{"port < 1024", "127,0,0,1,3,1", "501 Sorry, but I won't connect to ports < 1024"},
		{"127.0.0.1:1034", "127,0,0,1,4,10", "200 PORT command successful"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server, client := net.Pipe()
			go func() {
				handleConn(server)
				server.Close()
			}()
			input := bufio.NewScanner(client)
			input.Scan() // 220 Welcome

			// Login
			fmt.Fprintln(client, "USER test")
			input.Scan()

			// PORT
			fmt.Fprintln(client, "PORT", tt.arg)
			input.Scan()
			got := input.Text()
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
			client.Close()
		})
	}
}
func Test_Type(t *testing.T) {
	tests := []struct {
		name  string
		arg   string
		want1 string
		want2 string
	}{
		{"no arg", "", "501-Missing argument", "501 TYPE is now ASCII"},
		{"ascii", "A", "200-TYPE command successful", "200 TYPE is now ASCII"},
		{"image", "i", "200-TYPE command successful", "200 TYPE is now 8-bit binary"},
		{"unknown type", "xx", "504-Unknown Type: xx", "504 TYPE is now ASCII"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server, client := net.Pipe()
			go func() {
				handleConn(server)
				server.Close()
			}()
			input := bufio.NewScanner(client)
			input.Scan() // 220 Welcome

			// TYPE
			fmt.Fprintln(client, "TYPE", tt.arg)
			input.Scan()
			got1 := input.Text()
			if got1 != tt.want1 {
				t.Errorf("got %q, want %q", got1, tt.want1)
			}
			input.Scan()
			got2 := input.Text()
			if got2 != tt.want2 {
				t.Errorf("got %q, want %q", got2, tt.want2)
			}
			client.Close()
		})
	}
}

func Test_Mode(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{"no arg", "", "501 Missing argument"},
		{"unknown mode", "xx", "504 Please use S(tream) mode"},
		{"stream mode", "S", "200 S OK"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server, client := net.Pipe()
			go func() {
				handleConn(server)
				server.Close()
			}()
			input := bufio.NewScanner(client)
			input.Scan() // 220 Welcome

			// MODE
			fmt.Fprintln(client, "MODE", tt.arg)
			input.Scan()
			got := input.Text()
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
			client.Close()
		})
	}
}

func Test_Stru(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{"no arg", "", "501 Missing argument"},
		{"unknown structure", "xx", "504 Only F(ile) is supported"},
		{"File structure", "F", "200 F OK"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server, client := net.Pipe()
			go func() {
				handleConn(server)
				server.Close()
			}()
			input := bufio.NewScanner(client)
			input.Scan() // 220 Welcome

			// STRU
			fmt.Fprintln(client, "STRU", tt.arg)
			input.Scan()
			got := input.Text()
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
			client.Close()
		})
	}
}

func Test_RetrNoConnect(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{"no args", "", "501 No file name"},
		{"no data connection", "memo.md", "425 No data connection"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server, client := net.Pipe()
			go func() {
				handleConn(server)
				server.Close()
			}()
			input := bufio.NewScanner(client)
			input.Scan() // 220 Welcome

			// Login
			fmt.Fprintln(client, "USER test")
			input.Scan()

			// PORT
			fmt.Fprintln(client, "PORT 127,0,0,1,4,10") // 127.0.0.1:1034
			input.Scan()

			// RETR
			fmt.Fprintln(client, "RETR", tt.arg)
			input.Scan()
			got := input.Text()
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
			client.Close()
		})
	}
}

func Test_RetrConnect(t *testing.T) {

	// Data Connection
	_, err := net.Listen("tcp", "localhost:1034")
	if err != nil {
		t.Errorf("listen error: %s", err)
	}

	tests := []struct {
		name  string
		arg   string
		want1 string
		want2 string
	}{
		{"not exist file", "no.txt", "150 Accepted data connection", "553 Can't open that file: open no.txt: no such file or directory"},
		{"exist file", "memo.md", "150 Accepted data connection", "226 File successfully transferred"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server, client := net.Pipe()
			go func() {
				handleConn(server)
				server.Close()
			}()
			input := bufio.NewScanner(client)
			input.Scan() // 220 Welcome

			// Login
			fmt.Fprintln(client, "USER test")
			input.Scan()

			// PORT
			fmt.Fprintln(client, "PORT 127,0,0,1,4,10") // 127.0.0.1:1034
			input.Scan()

			// RETR
			fmt.Fprintln(client, "RETR", tt.arg)
			input.Scan()
			got1 := input.Text()
			if got1 != tt.want1 {
				t.Errorf("got %q, want %q", got1, tt.want1)
			}
			input.Scan()
			got2 := input.Text()
			if got2 != tt.want2 {
				t.Errorf("got %q, want %q", got2, tt.want2)
			}

			client.Close()
		})
	}
}

func Test_StorNoConnect(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{"no args", "", "501 No file name"},
		{"no data connection", "memo.md", "425 No data connection"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server, client := net.Pipe()
			go func() {
				handleConn(server)
				server.Close()
			}()
			input := bufio.NewScanner(client)
			input.Scan() // 220 Welcome

			// Login
			fmt.Fprintln(client, "USER test")
			input.Scan()

			// PORT
			fmt.Fprintln(client, "PORT 127,0,0,1,4,10") // 127.0.0.1:1034
			input.Scan()

			// STOR
			fmt.Fprintln(client, "STOR", tt.arg)
			input.Scan()
			got := input.Text()
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
			client.Close()
		})
	}
}

func Test_StorConnect(t *testing.T) {

	// Data Connection
	listener, err := net.Listen("tcp", "localhost:1034")
	if err != nil {
		t.Errorf("listen error: %s", err)
	}
	go func() {
		// Accept
		for {
			conn, err := listener.Accept()
			if err != nil {
				t.Errorf("listen error: %s", err)
			}
			conn.Write([]byte("test\r\n"))
			conn.Close()
		}
	}()

	tests := []struct {
		name  string
		arg   string
		want1 string
		want2 string
	}{
		{"not exist dir", "no/no.txt", "150 Accepted data connection", "553 Can't open that file: open no/no.txt: no such file or directory"},
		{"exist file", "test.txt", "150 Accepted data connection", "226 File successfully transferred"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server, client := net.Pipe()
			go func() {
				handleConn(server)
				server.Close()
			}()
			input := bufio.NewScanner(client)
			input.Scan() // 220 Welcome

			// Login
			fmt.Fprintln(client, "USER test")
			input.Scan()

			// PORT
			fmt.Fprintln(client, "PORT 127,0,0,1,4,10") // 127.0.0.1:1034
			input.Scan()

			// STOR
			fmt.Fprintln(client, "STOR", tt.arg)
			input.Scan()
			got1 := input.Text()
			if got1 != tt.want1 {
				t.Errorf("got %q, want %q", got1, tt.want1)
			}
			input.Scan()
			got2 := input.Text()
			if got2 != tt.want2 {
				t.Errorf("got %q, want %q", got2, tt.want2)
			}

			client.Close()
			os.Remove(tt.arg)
		})
	}
}
