package checkpoint

func Hiddenp() int {
	s1 := "DD"
	s2 := "DABC"
	str := ""
	index := 0
	for i := 0; i < len(s1); i++ {
		for j := index; j < len(s2); j++ {
			if s1[i] == s2[j] {
				str += string(s1[i])
				index++
				break
			} else {
				continue
			}
		}
	}

	if len(str) == len(s1) {
		return 1
	}

	return 0
}
