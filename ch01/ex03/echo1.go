package main

// func main() {
// 	fmt.Println(echo1(os.Args))
// }

func echo1(args []string) string {
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	return s
}
