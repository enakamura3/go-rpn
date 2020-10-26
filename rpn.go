package rpn

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/enakamura3/go-lifo"
)

// Exec evaluate an expression
func Exec(exp string) (int64, error) {
	l := lifo.New()
	exp = strings.Trim(exp, " ")
	exps := strings.Split(exp, " ")
	for _, v := range exps {
		s := string(v)
		switch s {
		case "+":
			a, b, err := getValues(&l)
			if err != nil {
				return 0, err
			}
			c := a + b
			lifo.Add(&l, strconv.FormatInt(c, 10))
		case "-", "−":
			a, b, err := getValues(&l)
			if err != nil {
				return 0, err
			}
			c := a - b
			lifo.Add(&l, strconv.FormatInt(c, 10))
		case "x", "×", "*":
			a, b, err := getValues(&l)
			if err != nil {
				return 0, err
			}
			c := a * b
			lifo.Add(&l, strconv.FormatInt(c, 10))
		case "/":
			a, b, err := getValues(&l)
			if err != nil {
				return 0, err
			}
			c := a / b
			lifo.Add(&l, strconv.FormatInt(c, 10))
		default:
			lifo.Add(&l, s)
		}
	}
	out, err := strconv.ParseInt(lifo.Remove(&l), 10, 64)
	if err != nil {
		return 0, err
	}
	return out, nil
}

func getValues(l *lifo.Lifo) (int64, int64, error) {
	b, err := strconv.ParseInt(lifo.Remove(l), 10, 64)
	if err != nil {
		fmt.Printf("Error converting '%v' to int\n", b)
		return 0, 0, err
	}
	a, err := strconv.ParseInt(lifo.Remove(l), 10, 64)
	if err != nil {
		fmt.Printf("Error converting '%v' to int\n", a)
		return 0, 0, err
	}
	return a, b, nil
}
