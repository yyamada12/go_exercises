package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/yyamada12/go_exercises/ch10/ex02/archive"

	_ "github.com/yyamada12/go_exercises/ch10/ex02/archive/tar"
	_ "github.com/yyamada12/go_exercises/ch10/ex02/archive/zip"
)

func main() {
	// reader, format, err := archive.NewReader("ch01.tar")
	reader, format, err := archive.NewReader("ch01.zip")
	fmt.Println(format)
	fmt.Println(err)
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(b)
}
