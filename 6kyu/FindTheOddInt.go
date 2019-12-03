// Given an array, find the integer that appears an odd number of times.
// There will always be only one integer that appears an odd number of times.

// https://www.codewars.com/kata/54da5a58ea159efa38000836/train/go

package findtheoddint

func myfindOdd(seq []int) int {
	occurrence := make(map[int]int)
	for _, element := range seq {
		occurrence[element]++
	}
	var returned int
	for key, value := range occurrence {
		if value%2 == 1 {
			returned = key
			break
		}
	}
	return returned
}

func bestFindOdd(seq []int) int {
	res := 0
	for _, x := range seq {
		res ^= x
	}
	return res
}

func main() {
}
