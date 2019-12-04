// https://www.codewars.com/kata/583f158ea20cfcbeb400000a/train/go

package arefmetic

func arithmetic(a int, b int, operator string) int {

	switch operator {
	case "add":
		return a + b
	case "subtract":
		return a - b
	case "multiply":
		return a * b
	case "divide":
		return a / b
	}

	return 0
}

func main() {}
