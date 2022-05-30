package checkpoint

func Romannumbers(n int) string {
	num := n
	numbers := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	roman := ""
	mpRoman := make(map[int]string)
	i := 0
	z := 0

	for num > 0 {
		rRoot := num / numbers[i]
		for j := 0; j < rRoot; j++ {
			roman += symbols[i]
			mpRoman[z] = symbols[i]
			num -= numbers[i]
			z++

		}
		i++

	}

	calc := ""
	for j := 0; j < z; j++ {
		if len(mpRoman[j]) == 1 {
			calc += mpRoman[j]
		} else {
			calc += "(" + string(mpRoman[j][1]) + "-" + string(mpRoman[j][0]) + ")"
		}

		if j != z-1 {
			calc += "+"
		}
	}
	return calc + "\n" + roman
}
