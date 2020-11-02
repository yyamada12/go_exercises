package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}
	for _, arg := range os.Args[1:] {
		a := strings.Split(arg, "=")
		if len(a) < 2 {
			log.Fatalf("invalid arg: %s", arg)
		}
		go clock(a[0], a[1])
	}
	for {
	}
}

func printUsage() {
	fmt.Println(`USAGE:
	go run clockwall.go [CLOCK NAME]=[ADDR], [[CLOCK NAME]=[ADDR], ...]

CLOCK NAME:
	give the clock name	like "Tokyo", "NewYork"
ADDR:
	the addr to connect like "localhost:8000"`)
}

func clock(tz, addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(tz + ": " + scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
