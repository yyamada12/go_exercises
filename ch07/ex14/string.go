package eval

import (
	"fmt"
	"strings"
)

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return fmt.Sprintf("%g", l)
}

func (u unary) String() string {
	switch e := u.x.(type) {
	case Var, literal, call:
		return fmt.Sprintf("%c%s", u.op, e.String())
	default:
		return fmt.Sprintf("%c(%s)", u.op, e.String())
	}
}

func (b binary) String() string {
	var s strings.Builder

	switch e := b.x.(type) {
	case Var, literal, unary, call:
		s.WriteString(e.String())
	case binary:
		if (b.op == '*' || b.op == '/') && (e.op == '+' || e.op == '-') {
			// (x+y)*?, (x-y)*?, (x+y)/?, (x-y)/?
			s.WriteString("(" + e.String() + ")")
		} else {
			s.WriteString(e.String())
		}
	default:
		s.WriteString("(" + e.String() + ")")
	}

	s.WriteRune(b.op)

	switch e := b.y.(type) {
	case Var, literal, call:
		s.WriteString(e.String())
	case binary:
		if b.op == '/' {
			// ?/(x+y), ?/(x-y), ?/(x*y), ?/(x/y)
			s.WriteString("(" + e.String() + ")")
		} else if (b.op == '-' || b.op == '*') && (e.op == '+' || e.op == '-') {
			// ?-(x+y), ?-(x-y), ?*(x+y), ?*(x-y)
			s.WriteString("(" + e.String() + ")")
		} else {
			s.WriteString(e.String())
		}
	default:
		s.WriteString("(" + e.String() + ")")
	}
	return s.String()
}

func (c call) String() string {
	var s strings.Builder
	s.WriteString(c.fn)
	s.WriteByte('(')
	for i, arg := range c.args {
		if i != 0 {
			s.WriteString(",")
		}
		s.WriteString(arg.String())
	}
	s.WriteByte(')')
	return s.String()
}

func (a aggregate) String() string {
	var s strings.Builder
	s.WriteString(a.fn)
	s.WriteByte('[')
	for i, arg := range a.args {
		if i != 0 {
			s.WriteString(",")
		}
		s.WriteString(arg.String())
	}
	s.WriteByte(']')
	return s.String()
}
