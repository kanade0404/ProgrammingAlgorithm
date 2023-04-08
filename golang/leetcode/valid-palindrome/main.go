package valid_palindrome

func isPalindrome(s string) bool {
	var target []byte
	for i := range s {
		if s[i] < 48 || (s[i] > 57 && s[i] < 65) || (s[i] > 90 && s[i] < 97) || s[i] > 122 {
			continue
		}
		if s[i] > 64 && s[i] < 91 {
			c := s[i] + 32
			target = append(target, c)
			continue
		}
		target = append(target, s[i])
	}
	if len(target) == 0 {
		return true
	}
	l := len(target)
	halfLen := len(target) / 2
	for i := 0; i < halfLen; i++ {
		if target[i] != target[l-1-i] {
			return false
		}
	}
	return true
}
