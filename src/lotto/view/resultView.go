package view

import (
	"fmt"
	"lotto/domain"
	"lotto/enum"
	"strconv"
)

type ResultView struct {
	result *domain.Result
}

func NewResultView(result *domain.Result) *ResultView {
	return &ResultView{result: result}
}

func (rv *ResultView) ShowResult()  {
	fmt.Println("당첨 통계");
	prizeList := rv.result.PrizeList()
	for k, v := range prizeList{
		rv.ResultMessage(k,v)
	}
	rv.showYield()
}

func (rv *ResultView) ResultMessage(rank, quantity int)  {
	if rank == 5 {
		return
	}
	fmt.Println(strconv.Itoa(rank+1) + "등(" + strconv.Itoa(enum.WinningMoney[rank]) + "원)- " + strconv.Itoa(quantity) +"개")
}

func (rv *ResultView) showYield() {
	fmt.Printf("총 수익률은 %.2f%%입니다",rv.result.CalculateYield())
}