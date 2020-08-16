package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	echo(os.Args)
}

func echo(args []string) {
	fmt.Println(strings.Join(args, " "))
}
