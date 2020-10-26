package rpn

import (
	"testing"
)

func TestExec(t *testing.T) {
	exp := "5 1 2 + 4 × + 3 −"
	val, _ := Exec(exp)
	var want int64 = 14
	if val != want {
		t.Fatalf("Want '%v', got '%v'", want, val)
	}
}
