// Tempflag prints the value of its -temp (temperature) flag.
package main

import (
	"flag"
	"fmt"

	"github.com/yyamada12/go_exercises/ch07/ex06/tempconv"
)

var celsius = tempconv.CelsiusFlag("celsius", 20.0, "the temperature in Celsius")
var kelvin = tempconv.KelvinFlag("kelvin", 20.0, "the temperature in Kelvin")

func main() {
	flag.Parse()
	fmt.Println(*celsius, *kelvin)
}
