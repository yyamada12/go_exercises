package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/yyamada12/go_exercises/ch07/ex15/eval"
)

func main() {
	// read formula from stdin
	text, err := scanText()
	handleErr(err)

	// parse formula
	expr, err := eval.Parse(text)
	handleErr(err)

	// check and extract vars
	vars := make(map[eval.Var]bool)
	err = expr.Check(vars)
	handleErr(err)

	// read env values from stdin
	env, err := scanVars(vars)
	handleErr(err)

	// eval formula
	val := expr.Eval(env)
	fmt.Printf("%s = %.6g\n", expr, val)
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func scanText() (string, error) {
	_, err := fmt.Printf("please input a formula.\n> ")
	if err != nil {
		return "", err
	}
	scanner := bufio.NewScanner(os.Stdin)
	ok := scanner.Scan()
	if !ok {
		return "", fmt.Errorf("cannot read from stdin")
	}
	return scanner.Text(), nil
}

func scanVars(vars map[eval.Var]bool) (eval.Env, error) {
	var env eval.Env
	env = make(map[eval.Var]float64)
	for v := range vars {
		var val float64
		_, err := fmt.Printf("please input value of %s\n> ", v)
		if err != nil {
			return env, err
		}
		_, err = fmt.Scan(&val)
		if err != nil {
			return env, err
		}
		env[v] = val
	}
	return env, nil
}
