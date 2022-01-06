package controller

import (
	"bufio"
	"fmt"
	"lotto/domain"
	"lotto/view"
	"os"
	"strconv"
	"strings"
)



type LottoController struct {
	money, autoQuantity, manualQuantity, bonusBall int
	lr *domain.LottoRepository
	winningNums []int
}

func NewLottoController () *LottoController {
	lc := &LottoController{}
	lc.lottoPurchase()
	lc.lottoResult()
	return lc
}

func (lc *LottoController) lottoPurchase() {
	sc := bufio.NewScanner(os.Stdin)
	fmt.Println("구입금액을 입력해 주세요")
	sc.Scan()
	lc.money, _ = strconv.Atoi(sc.Text())
	lc.autoQuantity = lc.money / 1000

	fmt.Println("수동으로 구매할 로또 수를 입력해주세요")
	sc.Scan()
	lc.manualQuantity, _ = strconv.Atoi(sc.Text())

	fmt.Println("수동으로 구매할 번호를 입력해주세요")
	lc.lr = domain.NewLottoRepository(lc.autoQuantity, lc.manualQuantity)
	fmt.Println("수동으로" , lc.manualQuantity,"장, 자동으로 ",lc.autoQuantity, "장 구매했습니다")
	fmt.Println(lc.lr)
}

func (lc *LottoController) lottoResult()  {
	sc := bufio.NewScanner(os.Stdin)
	fmt.Println("지난 주 당첨번호를 입력해주세요")
	sc.Scan()
	for _ , input := range strings.Split(sc.Text(),","){
		tmp , _ := strconv.Atoi(input)
		lc.winningNums = append(lc.winningNums, tmp)
	}
	fmt.Println("보너스 볼을 입력해 주세요")
	sc.Scan()
	lc.bonusBall, _ = strconv.Atoi(sc.Text())

	result := domain.NewResult(lc.lr,lc.winningNums,lc.bonusBall)
	rv := view.NewResultView(result)
	rv.ShowResult()

}