package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"strings"
)

var out = flag.String("out", "png", "output image format")

func main() {
	flag.Parse()

	var err error
	switch strings.ToLower(*out) {
	case "jpeg", "jpg":
		err = toJPEG(os.Stdin, os.Stdout)
	case "png":
		err = toPNG(os.Stdin, os.Stdout)
	case "gif":
		err = toGIF(os.Stdin, os.Stdout)
	default:
		err = fmt.Errorf("unsupported format %q", *out)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "imgconvert: %v\n", err)
		os.Exit(1)
	}
}

func toJPEG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}

func toPNG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return png.Encode(out, img)
}

func toGIF(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return gif.Encode(out, img, &gif.Options{})
}
