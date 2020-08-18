package main

// func main() {
// 	fmt.Println(echo1(os.Args))
// }

func echo2(args []string) string {
	var s, sep string
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	return s
}
