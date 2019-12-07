// https://www.codewars.com/kata/55fd2d567d94ac3bc9000064/train/go

package kata

func RowSumOddNumbers(n int) int {
	// your code here
	return (1 + (n-1)*n/2.0 + (n - 1)) * n
}
