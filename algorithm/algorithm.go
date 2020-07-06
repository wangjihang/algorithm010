package algorithm

func lemonadeChange(bills []int) bool {
	var five, ten int
	for _, bill := range bills {
		if bill == 5 {
			five++
		} else if bill == 10 {
			if five < 1 {
				return false
			}
			five--
			ten++
		} else {
			if ten > 0 && five > 0 {
				ten--
				five--
			} else if five > 3 {
				five -= 3
			} else {
				return false
			}
		}
	}
	return true
}
