package valid_parentheses

import "strings"

type Stack struct {
	Arr []string
	N   int
}

func isValid(s string) bool {
	strs := strings.Split(s, "")
	stack := Stack{
		Arr: []string{},
		N:   0,
	}
	for i := range strs {
		if stack.N == 0 {
			if strs[i] == ")" || strs[i] == "}" || strs[i] == "]" {
				return false
			} else {
				stack.Arr = append(stack.Arr, strs[i])
				stack.N += 1
			}
		} else {
			if strs[i] == "(" || strs[i] == "{" || strs[i] == "[" {
				stack.Arr = append(stack.Arr, strs[i])
				stack.N += 1
			} else {
				if (stack.Arr[stack.N-1] == "(" && strs[i] == ")") ||
					(stack.Arr[stack.N-1] == "{" && strs[i] == "}") ||
					(stack.Arr[stack.N-1] == "[" && strs[i] == "]") {
					stack.Arr = stack.Arr[:stack.N-1]
					stack.N -= 1
				} else {
					stack.Arr = append(stack.Arr, strs[i])
					stack.N += 1
				}
			}
		}
	}
	return stack.N == 0
}
