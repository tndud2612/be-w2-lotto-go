package enum

var WinningMoney = map[int]int{
	0:2000000000,
	1:30000000,
	2:1500000,
	3:50000,
	4:5000,
	5:0,
}

func WinningMoneyMatch(countOfMatch int, bonus bool) int {
	switch countOfMatch {
	case 6:
		return 0
	case 5:
		if bonus {
			return 1
		}
		return 2
	case 4:
		return 3
	case 3:
		return 4
	default:
		return 5
	}
}