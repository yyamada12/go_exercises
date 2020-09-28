package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/big"
	"math/cmplx"
	"os"

	"github.com/yyamada12/go_exercises/ch03/ex08/bigcomplex"
)

var iterations = uint8(4)

func main() {

	if len(os.Args) < 2 {
		printUsage()
		return
	}
	var newton func(x, y float64) color.Color
	switch os.Args[1] {
	case "complex64":
		newton = newtonWithComplex64
	case "complex128":
		newton = newtonWithComplex128
	case "complexFloat":
		newton = newtonWithComplexFloat
	case "complexRat":
	default:
		printUsage()
		return
	}

	const (
		// for big.Rat
		xminNum, yminNum, xmaxNum, ymaxNum = 97, 0, 100, 3
		denom                              = 100

		xmin, ymin, xmax, ymax = 0.97, 0, 1, 0.03
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	switch os.Args[1] {
	case "complex64", "complex128", "complexFloat":
		for py := 0; py < height; py++ {
			y := float64(py)/height*(ymax-ymin) + ymin
			for px := 0; px < width; px++ {
				x := float64(px)/width*(xmax-xmin) + xmin
				img.Set(px, py, newton(x, y))
			}
		}
	case "complexRat":
		for py := 0; py < height; py++ {
			y := big.NewRat(int64(py)*(ymaxNum-yminNum)+yminNum*height, height*denom)
			for px := 0; px < width; px++ {
				x := big.NewRat(int64(px)*(xmaxNum-xminNum)+xminNum*width, width*denom)
				img.Set(px, py, newtonWithComplexRat(x, y))
			}
		}
	}

	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func printUsage() {
	fmt.Fprintln(os.Stderr,
		`USAGE:
	go run fractal.go OPTION > out.png
OPTION: 
	complex64		use complex64 
	complex128		use complex128
	complexFloat	use complexFloat based on big.Float
	complexRat		use	complexRat based on big.Rat`)
}

func newtonWithComplex64(x, y float64) color.Color {
	z := complex64(complex(x, y))
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(complex128(z-1)) < 1e-6 || cmplx.Abs(complex128(z-1i)) < 1e-6 || cmplx.Abs(complex128(z+1)) < 1e-6 || cmplx.Abs(complex128(z+1i)) < 1e-6 {
			// if cmplx.Abs(complex128(z*z*z*z-1)) < 1e-6 {
			return color.RGBA{242 - i, 213 - 7*i, 15 + 3*i, 255}
		}
	}
	return color.Black
}

func newtonWithComplex128(x, y float64) color.Color {
	z := complex(x, y)
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z-1) < 1e-6 || cmplx.Abs(z-1i) < 1e-6 || cmplx.Abs(z+1) < 1e-6 || cmplx.Abs(z+1i) < 1e-6 {
			// if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.RGBA{242 - i, 213 - 7*i, 15 + 3*i, 255}
		}
	}
	return color.Black
}

func newtonWithComplexFloat(x, y float64) color.Color {
	if x == 0 && y == 0 {
		return color.Black
	}
	z := bigcomplex.NewCmplFloat(big.NewFloat(x), big.NewFloat(y))
	for i := uint8(0); i < iterations; i++ {
		one := bigcomplex.NewCmplFloat(big.NewFloat(1), big.NewFloat(0))
		onei := bigcomplex.NewCmplFloat(big.NewFloat(0), big.NewFloat(1))
		four := bigcomplex.NewCmplFloat(big.NewFloat(4), big.NewFloat(0))
		// z -= (z - 1/(z*z*z)) / 4
		z = z.Minus(z.Minus(one.Divides(z.Times(z).Times(z))).Divides(four))
		// if z.Times(z).Times(z).Times(z).Minus(one).SquaredAbs().Cmp(big.NewFloat(1e-12)) < 0 {
		if z.Minus(one).SquaredAbs().Cmp(big.NewFloat(1e-12)) < 0 || z.Minus(onei).SquaredAbs().Cmp(big.NewFloat(1e-12)) < 0 || z.Plus(one).SquaredAbs().Cmp(big.NewFloat(1e-12)) < 0 || z.Plus(onei).SquaredAbs().Cmp(big.NewFloat(1e-12)) < 0 {
			return color.RGBA{242 - i, 213 - 7*i, 15 + 3*i, 255}
		}
	}
	return color.Black
}

func newtonWithComplexRat(x, y *big.Rat) color.Color {
	if x.Sign() == 0 && y.Sign() == 0 {
		return color.Black
	}
	z := bigcomplex.NewCmplRat(x, y)
	for i := uint8(0); i < iterations; i++ {
		one := bigcomplex.NewCmplRat(big.NewRat(1, 1), big.NewRat(0, 1))
		onei := bigcomplex.NewCmplRat(big.NewRat(0, 1), big.NewRat(1, 1))
		four := bigcomplex.NewCmplRat(big.NewRat(4, 1), big.NewRat(0, 1))
		// z -= (z - 1/(z*z*z)) / 4
		z = z.Minus(z.Minus(one.Divides(z.Times(z).Times(z))).Divides(four))
		// if z.Times(z).Times(z).Times(z).Minus(one).SquaredAbs().Cmp(big.NewRat(1, 1e12)) < 0 {
		if z.Minus(one).SquaredAbs().Cmp(big.NewRat(1, 1e12)) < 0 || z.Minus(onei).SquaredAbs().Cmp(big.NewRat(1, 1e12)) < 0 || z.Plus(one).SquaredAbs().Cmp(big.NewRat(1, 1e12)) < 0 || z.Plus(onei).SquaredAbs().Cmp(big.NewRat(1, 1e12)) < 0 {
			return color.RGBA{242 - i, 213 - 7*i, 15 + 3*i, 255}
		}
	}
	return color.Black
}
