package main

import "fmt"

func main() {
	err := panicAndRecover()
	if err != nil {
		fmt.Println(err)
	}
}

func panicAndRecover() (err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("return by panic and recover")
		}
	}()
	panic(struct{}{})
}
