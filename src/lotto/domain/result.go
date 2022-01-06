package domain

import "lotto/enum"

type Result struct {
	lr                  *LottoRepository
	winningNums         []int
	bonusBall, quantity int
	prizeList           [6]int
}

func NewResult(lr *LottoRepository, winningNums []int, bonusBall int) *Result {
	result := &Result{
		lr:          lr,
		winningNums: winningNums,
		quantity:    len(lr.List()),
		bonusBall:   bonusBall,
	}
	result.setPrizeList()
	return result
}

func (result *Result)PrizeList() [6]int {
	return result.prizeList
}

func (result *Result) setPrizeList() {
	for _, lotto := range result.lr.List() {
		rank := enum.WinningMoneyMatch(result.numberMatch(lotto), result.bonusMatch(lotto, result.bonusBall))
		result.prizeList[rank]++
	}
}

func (result *Result) numberMatch(lotto []int) int {
	res := 0
	for _, num := range lotto {
		if contains(result.winningNums, num) {
			res++
		}
	}
	return res
}

func (result *Result) bonusMatch(lotto []int, bonusBall int) bool {
	return contains(lotto, bonusBall)
}

func contains(list []int, i int) bool {
	for _, j := range list {
		if j == i {
			return true
		}
	}
	return false
}

func (result *Result) CalculateYield() float64 {
	var totalPrize int
	for k, v := range result.prizeList {
		totalPrize += v * enum.WinningMoney[k]
	}
	totalSpent := result.quantity * 1000
	return float64(totalPrize)/float64(totalSpent)*100 - 100
}
