package unique

func unique(str []string) []string {
	last := 0
	for _, s := range str {
		if str[last] == s {
			continue
		}
		last++
		str[last] = s
	}
	return str[:last+1]
}
