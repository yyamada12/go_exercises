package memo_test

import (
	"testing"

	memo "github.com/yyamada12/go_exercises/ch09/ex03"
	"github.com/yyamada12/go_exercises/ch09/ex03/memotest"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := memo.New(httpGetBody)
	defer m.Close()
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	defer m.Close()
	memotest.Concurrent(t, m)
}
