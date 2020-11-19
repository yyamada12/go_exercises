package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

type result struct {
	tz   string
	time string
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}
	clocks := map[string]string{}
	results := make(chan result)
	tzs := []string{}
	for _, arg := range os.Args[1:] {
		a := strings.Split(arg, "=")
		if len(a) != 2 {
			log.Fatalf("invalid arg: %s", arg)
		}
		tzs = append(tzs, a[0])
		go clock(a[0], a[1], results)
	}
	for {
		r := <-results
		clocks[r.tz] = r.time
		printClocks(clocks, tzs)
	}
}

func printUsage() {
	fmt.Println(`USAGE:
	go run clockwall.go [CLOCK NAME]=[ADDR] [[CLOCK NAME]=[ADDR] ...]
	ex) go run clockwall.go Tokyo=localhost:8000 NewYork=localhost:8010

CLOCK NAME:
	give the clock name	like "Tokyo", "NewYork"
ADDR:
	the addr to connect like "localhost:8000"`)
}

func printClocks(clocks map[string]string, tzs []string) {
	var times []string
	for _, tz := range tzs {
		if time, ok := clocks[tz]; ok {
			times = append(times, tz+" "+time)
		}
	}
	fmt.Print("\r" + strings.Join(times, "    "))
}

func clock(tz, addr string, results chan result) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		results <- result{tz, scanner.Text()}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
