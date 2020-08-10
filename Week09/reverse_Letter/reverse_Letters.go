package reverse_Letter

func reverseOnlyLetters(S string) string {
	if len(S) <= 1 {
		return S
	}

	buf := []byte(S)
	for i, j := 0, len(buf)-1; i < j; {
		for i < j && !isLetter(buf[i]) {
			i++
		}
		for i < j && !isLetter(buf[j]) {
			j--
		}
		buf[i], buf[j] = buf[j], buf[i]
		i++
		j--
	}
	return string(buf)
}

func isLetter(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}
