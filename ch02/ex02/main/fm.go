// Fm converts its numeric argument to Feet and Meter.
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/yyamada12/go_exercises/ch02/ex02/lengthconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := lengthconv.Feet(t)
		m := lengthconv.Meter(t)
		fmt.Printf("%s = %s, %s = %s\n",
			f, lengthconv.FToM(f), m, lengthconv.MToF(m))
	}
}
