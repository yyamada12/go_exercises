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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server, client := net.Pipe()
			go func() {
				handleConn(server)
				server.Close()
			}()
			fmt.Fprintf(client, "%s %s\n", tt.cmd, tt.arg)
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
