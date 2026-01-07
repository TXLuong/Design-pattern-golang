func longestNiceSubstring(s string) string {
    n := len(s)
	var result string = ""
	maxLen := 0
    for i:= 0 ; i < n; i++ {
        for j:= i+1; j < n; j++ {
			if isNiceString(s[i:j+1]) {
				if maxLen < (j-i+1) {
					result = s[i:(j+1)]
				}
			}
        }
    }
	return result
}

func isNiceString(s string) bool {
    table := make(map[rune][int]])
    for _, val := range s {
        table[val]++
    }
    for _, val := range s {
        if unicode.IsUpper(val){
			// check if lower case exist 
			if _, ok := table[val+32]; !ok {
				return false
			}
		} else {
			// check if upper case exist 
			if _, ok := table[val-32]; !ok {
				return false
			}
		}
    }
	return true
}