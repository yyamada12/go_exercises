package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var out io.Writer = os.Stdout

func main() {
	echo(os.Args)
}

func echo(args []string) {
	fmt.Fprintln(out, strings.Join(args, " "))
}
