package goal_parser_interpretation

func interpret(command string) string {
	var prev, result string
	commandLen := len(command)
	for i := 0; i < commandLen; i++ {
		c := string(command[i])
		if prev == "(" {
			if c == ")" {
				result += "o"
				prev = ""
			} else {
				prev = c
			}
		} else if prev == ")" {
			prev = c
		} else {
			result += prev
			prev = c
		}
	}
	if !(prev == "(" || prev == ")") {
		result += prev
	}
	return result
}
