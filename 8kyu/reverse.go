//  https://www.codewars.com/kata/5168bb5dfe9a00b126000018/train/go

package reverse

func solution(word string) string {
	var answer string
	for i := len(word) - 1; i > -1; i-- {
		answer = answer + string(word[i])
	}
	return answer
}

func main() {}
