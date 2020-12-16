package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/yyamada12/go_exercises/ch10/ex02/archive"

	_ "github.com/yyamada12/go_exercises/ch10/ex02/archive/tar"
)

func main() {
	// r := NewReader("ch01.zip")

	file, err := os.Open("ch01.tar")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := bufio.NewReader(file)

	reader, format, err := archive.NewReader(r)
	fmt.Println(format)
	fmt.Println(err)
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(b)
}
