package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var bit = flag.Int("bit", 256, "SHA hash length")

func main() {
	flag.Parse()

	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	switch *bit {
	case 256:
		fmt.Printf("%x\n", sha256.Sum256(input.Bytes()))
	case 384:
		fmt.Printf("%x\n", sha512.Sum384(input.Bytes()))
	case 512:
		fmt.Printf("%x\n", sha512.Sum512(input.Bytes()))
	default:
		fmt.Println("bit must be 256, 384 or 512")
	}
}
