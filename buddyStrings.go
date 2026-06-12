func buddyStrings(s string, goal string) bool {
	if min(len(s), len(goal)) == 1 {
		return false
	}
	if len(s) != len(goal) {
		return false
	}

	var freq [26]byte 
	diff1, diff2 := -1, -1 
	hasTwo := false    
	for i := range s {
		if s[i] != goal[i] {
			if diff1 == -1 { 
				diff1 = i
			} else if diff2 == -1 { 
				diff2 = i
			} else {
				return false 
			}
		}
		freq[s[i] - 'a']++
		if freq[s[i] - 'a'] == 2 {
			hasTwo = true
		}
	}
	if diff2 != -1 { 
		return s[diff1] == goal[diff2] && s[diff2] == goal[diff1] 
	}
	if diff1 == -1 && hasTwo { 
		return true
	}
	return false
}
