package main

import "strings"

// func main() {
// 	fmt.Println(echo3(os.Args))
// }

func echo3(args []string) {
	strings.Join(args[1:], " ")
}
