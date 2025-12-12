package format

func DashToSpace(s string) string {
	temp := []byte(s)
	for i := range s {
		if s[i] == '-' {
			temp[i] = ' '
		} else {
			temp[i] = s[i]
		}
	}
	return string(temp)
}
