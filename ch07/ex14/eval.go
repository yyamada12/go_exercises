// Package eval provides an expression evaluator.
package eval

import (
	"fmt"
	"math"
)

type Env map[Var]float64

func (v Var) Eval(env Env) float64 {
	return env[v]
}

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}

func (a aggregate) Eval(env Env) float64 {
	var res float64
	switch a.fn {
	case "min":
		res = a.args[0].Eval(env)
		for _, arg := range a.args[1:] { // arg length >= 2 was validated by Check()
			v := arg.Eval(env)
			if res > v {
				res = v
			}
		}
		return res
	case "max":
		res = a.args[0].Eval(env)
		for _, arg := range a.args[1:] { // arg length >= 2 was validated by Check()
			v := arg.Eval(env)
			if res < v {
				res = v
			}
		}
		return res
	}
	panic(fmt.Sprintf("unsupported aggregate operation: %s", a.fn))
}
