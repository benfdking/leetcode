package leetcode

import "strconv"

func evalRPN(tokens []string) int {
	switch len(tokens) {
	case 0:
		panic("0 length")
	case 1:
		return parseInt(tokens[0])
	case 2:
		panic("don't know how to do deal with two")
	case 3:
		return operators[tokens[2]]([2]string{tokens[0], tokens[1]})
	default:
		for i, s := range tokens {
			if f, ok := operators[s] {

				tokens[0:i-2] + evalRPN()

				return
			}
		}
		panic("should never reach")
	}
}

var operators = map[string]func([2]string) int{
	"/": func(s [2]string) int {
		v1, v2 := parseInt(s[0]), parseInt(s[1])
		return v1 / v2
	},
	"*": func(s [2]string) int {
		v1, v2 := parseInt(s[0]), parseInt(s[1])
		return v1 * v2
	},
	"+": func(s [2]string) int {
		v1, v2 := parseInt(s[0]), parseInt(s[1])
		return v1 + v2
	},
	"-": func(s [2]string) int {
		v1, v2 := parseInt(s[0]), parseInt(s[1])
		return v1 - v2
	},
}

func parseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(s)
	}
	return n
}
