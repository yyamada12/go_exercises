package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"

	"github.com/yyamada12/go_exercises/ch10/ex02/archive"

	_ "github.com/yyamada12/go_exercises/ch10/ex02/archive/tar"
	_ "github.com/yyamada12/go_exercises/ch10/ex02/archive/zip"
)

func main() {
	// reader, format, err := archive.NewReader("ch01.tar")
	reader, format, err := archive.NewReader("ch01.zip")
	if err != nil {
		log.Fatal(err)
	}

	defer reader.Close()

	fmt.Println(format)

	for {
		header, err := reader.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Println(err)
		}
		fmt.Println("--- " + header.Name + " ---")
		if !header.IsDir {
			b, err := ioutil.ReadAll(reader)
			if err != nil {
				log.Println(err)
			}
			fmt.Println(string(b[:int(math.Min(float64(len(b)), 30))]))
		}
	}
}
