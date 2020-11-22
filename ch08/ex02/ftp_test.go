// Clock is a TCP server that periodically writes the time.
package main

import (
	"bufio"
	"fmt"
	"net"
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server, client := net.Pipe()
			go func() {
				handleConn(server)
				server.Close()
			}()
			fmt.Fprintln(client, tt.cmd, tt.arg)
			input := bufio.NewScanner(client)
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
		cmd  string
		arg  string
		want string
	}{
		{"no args", "PORT", "", "501 Syntax error in IP address"},
		{"invalid host ", "PORT", "256,255,255,1,10,1", "501 Syntax error in IP address"},
		{"invalid port ", "PORT", "127,0,0,1,256,1", "501 Syntax error in IP address"},
		{"port < 1024", "PORT", "127,0,0,1,3,1", "501 Sorry, but I won't connect to ports < 1024"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server, client := net.Pipe()
			go func() {
				handleConn(server)
				server.Close()
			}()
			// Login
			fmt.Fprintln(client, "USER test")
			input := bufio.NewScanner(client)
			input.Scan()

			// PORT
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
func Test_Type(t *testing.T) {
	tests := []struct {
		name  string
		cmd   string
		arg   string
		want1 string
		want2 string
	}{
		{"no arg", "TYPE", "", "501-Missing argument", "501 TYPE is now ASCII"},
		{"image", "TYPE", "i", "200-TYPE command successful", "200 TYPE is now 8-bit binary"},
		{"unknown type", "TYPE", "xx", "504-Unknown Type: xx", "504 TYPE is now ASCII"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server, client := net.Pipe()
			go func() {
				handleConn(server)
				server.Close()
			}()
			fmt.Fprintln(client, tt.cmd, tt.arg)
			input := bufio.NewScanner(client)
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
		cmd  string
		arg  string
		want string
	}{
		{"no arg", "MODE", "", "501 Missing argument"},
		{"unknown mode", "MODE", "xx", "504 Please use S(tream) mode"},
		{"stream mode", "MODE", "S", "200 S OK"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server, client := net.Pipe()
			go func() {
				handleConn(server)
				server.Close()
			}()
			fmt.Fprintln(client, tt.cmd, tt.arg)
			input := bufio.NewScanner(client)
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
		cmd  string
		arg  string
		want string
	}{
		{"no arg", "STRU", "", "501 Missing argument"},
		{"unknown structure", "STRU", "xx", "504 Only F(ile) is supported"},
		{"File structure", "STRU", "F", "200 F OK"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server, client := net.Pipe()
			go func() {
				handleConn(server)
				server.Close()
			}()
			fmt.Fprintln(client, tt.cmd, tt.arg)
			input := bufio.NewScanner(client)
			input.Scan()
			got := input.Text()
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
			client.Close()
		})
	}
}