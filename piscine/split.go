package piscine

func Split(s, sep string) []string {
	var result []string
	start := 0
	sepLen := len(sep)
	for i := 0; i < len(s); i++ {
		if i+sepLen <= len(s) && s[i:i+sepLen] == sep {
			if i != start {
				result = append(result, s[start:i])
			}
			i += sepLen - 1
			start = i + 1
		}
	}
	if start < len(s) {
		result = append(result, s[start:])
	}
	return result
}
