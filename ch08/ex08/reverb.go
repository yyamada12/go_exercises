// Reverb is a TCP server that simulates an echo.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	var wg sync.WaitGroup

	text := make(chan string)
	go func(input *bufio.Scanner, text chan string) {
		for input.Scan() {
			text <- input.Text()
		}
	}(input, text)
loop:
	for {
		select {
		case t := <-text:
			wg.Add(1)
			go func(c net.Conn, shout string) {
				defer wg.Done()
				echo(c, shout, 1*time.Second)
			}(c, t)
		case <-time.After(10 * time.Second):
			fmt.Println("timed out")
			break loop
		}
	}

	wg.Wait()
	// NOTE: ignoring potential errors from input.Err()
	c.Close()
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
