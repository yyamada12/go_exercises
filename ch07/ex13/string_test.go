package eval

import (
	"math"
	"testing"
)

func TestString(t *testing.T) {
	tests := []struct {
		expr string
		env  Env
		want string
	}{
		{"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, "sqrt(A/pi)"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 12, "y": 1}, "pow(x,3)+pow(y,3)"},
		{"5 / 9 * (F - 32)", Env{"F": -40}, "5/9*(F-32)"},
		{"-1 + -x", Env{"x": 1}, "-1+(-x)"},
		{"-1 - x", Env{"x": 1}, "-1-x"},
	}

	for _, test := range tests {
		expr, err := Parse(test.expr)
		if err != nil {
			t.Error(err) // parse error
			continue
		}
		gotString := expr.String()

		if gotString != test.want {
			t.Errorf("Parse(%q).String() got %q, want %q", test.expr, gotString, test.want)
		}

		gotExpr, err := Parse(gotString)
		if err != nil {
			t.Errorf("Parse(Expr.String()) got err: %s", err.Error())
		}

		gotVal := gotExpr.Eval(test.env)
		val := expr.Eval(test.env)

		if gotVal != val {
			t.Errorf("Parse(expr.String()).Eval() in %v = %.6g, want %.6g", test.env, gotVal, val)
		}
	}
}
