package piscine

func ConcatParams(args []string) string {
	result := ""

	for i, arg := range args {
		result = result + arg

		if i < len(args) {
			result = result + "\n"
		}
	}
	return result
}
