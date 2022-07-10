package leetcode

func isValid(s string) bool {
	return isValidSub(s, []uint8{})
}

func isValidSub(s string, ongoing []uint8) bool {
	if len(s) == 0 {
		return len(ongoing) == 0
	}
	// open
	o := s[0]
	if o == '(' || o == '{' || o == '[' {
		return isValidSub(s[1:], append(ongoing, s[0]))
	}
	// close
	if len(ongoing) == 0 {
		return false
	}
	closeCharacter := ongoing[len(ongoing)-1]
	if o == ')' && closeCharacter == '(' {
		return isValidSub(s[1:], ongoing[0:len(ongoing)-1])
	} else if o == ']' && closeCharacter == '[' {
		return isValidSub(s[1:], ongoing[0:len(ongoing)-1])
	} else if o == '}' && closeCharacter == '{' {
		return isValidSub(s[1:], ongoing[0:len(ongoing)-1])
	}
	return false
}
