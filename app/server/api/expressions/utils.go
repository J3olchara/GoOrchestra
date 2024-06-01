package expressions

func ContainsOnly(str string, chars string) bool {
	for i := range str {
		var flag = false
		for j := range chars {
			flag = flag || str[i] == chars[j]
		}
		if !flag {
			return false
		}
	}
	return true
}
