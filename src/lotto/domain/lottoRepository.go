package domain

import (
	"fmt"
)

type LottoRepository struct {
	list                         [][]int
	autoQuantity, manualQuantity int
}

func NewLottoRepository(autoQuantity, manualQuantity int) *LottoRepository {
	lr := &LottoRepository{
		autoQuantity:   autoQuantity,
		manualQuantity: manualQuantity,
	}
	lr.checkEmptyList()
	lr.auto()
	lr.manual()
	return lr
}

func (lr *LottoRepository) auto() {
	for i := 0; i < lr.autoQuantity; i++ {
		l := NewLotto()
		lr.list = append(lr.list, l.AutoPick())
	}
}

func (lr *LottoRepository) manual() {
	for i := 0; i < lr.manualQuantity; i++ {
		l := NewLotto()
		lr.list = append(lr.list, l.ManuallyPick())
	}
}

func (lr *LottoRepository) List() [][]int {
	return lr.list
}

func (lr *LottoRepository) checkEmptyList() {
	if lr.manualQuantity+lr.autoQuantity == 0 {
		panic("로또를 구매하지 않았습니다. 수익률 계산이 불가합니다.")
	}
}

func (lr *LottoRepository) String() string {
	var res string
	for _ , lotto := range lr.list {
		res += fmt.Sprintln(lotto)
	}
	return res
}


