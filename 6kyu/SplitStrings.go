// https://www.codewars.com/kata/515de9ae9dcfc28eb6000001/train/go

package splitstrings

func splitStrings(str string) []string {
	answer := make([]string, 0)
	for i := 0; i <= len(str); i += 2 {
		if str[i+1] != '' {
			append(answer, str[i:i+2])
		} else {
			append(answer, str[i]+'_')
		}
	}
	return answer
}

// not mine
func Solution(str string) []string {
	var res []string
	if len(str)%2 != 0 {
		str += "_"
	}
	for i := 0; i < len(str); i += 2 {
		res = append(res, str[i:i+2])
	}
	return res
}

func main() {}
