package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	echo(os.Args[1:])
}

func echo(args []string) {
	for i, arg := range args {
		fmt.Fprintln(out, i, arg)
	}
}
