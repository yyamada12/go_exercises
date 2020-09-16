// Delimiter prints its argument numbers with a comma or a space at each power of 1000.
//
// Example:
// 	$ go run delimiter.go 1 12 123 1234 1234567890 1234.567890
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//  1,234.566 890
//
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", delimiter(os.Args[i]))
	}
}

func delimiter(s string) string {
	if strings.ContainsRune(s, '.') {
		splitted := strings.Split(s, ".")
		return comma(splitted[0]) + "." + space(splitted[1])
	}
	return comma(s)
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	buf := bytes.Buffer{}
	for i, c := range s {
		buf.WriteRune(c)
		if (len(s)-i)%3 == 1 && i != len(s)-1 {
			buf.WriteRune(',')
		}
	}
	return buf.String()
}

func space(s string) string {
	buf := bytes.Buffer{}
	for i, c := range s {
		buf.WriteRune(c)
		if (i)%3 == 2 && i != len(s)-1 {
			buf.WriteRune(' ')
		}
	}
	return buf.String()
}
