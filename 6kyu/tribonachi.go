// https://www.codewars.com/kata/556deca17c58da83c00002db/train/go

package 6kyu


func sum(nums []float64) float64{
	sum_ := 0.0
	for _, num := range nums {
		sum_ += num
	}
	return sum_
}

func Tribonacci(signature [3]float64, n int) []float64 {
	// your code here
	ans := make([]float64, n)
	for i:= 0; i < n; i++ {
		if i < 3 {
			ans[i] = signature[i]
		} else {
		ans[i] = sum(ans[i-3:i])
		}
	}
	return ans
}

func main ()