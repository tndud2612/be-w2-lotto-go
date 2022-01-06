package domain

import (
	"bufio"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Lotto struct {
	numList []int
}

func NewLotto() *Lotto {
	l := &Lotto{
		numList: initNumList(),
	}
	return l
}

func initNumList() []int {
	res := make([]int, 0)
	for i := 1; i <= 45; i++ {
		res = append(res, i)
	}
	return res
}

func (l *Lotto) AutoPick() []int {
	shuffle(l.numList)
	res := l.numList[:6]
	sort.Ints(res)
	return res
}

func (l *Lotto) ManuallyPick() []int{
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	res := make([]int, 0)
	for _ , input := range strings.Split(sc.Text(), ","){
		tmp , _ := strconv.Atoi(input)
		res = append(res, tmp)
	}
	return res
}

func shuffle(list []int) {
	rand.Seed(time.Now().UnixNano())
	for i:=0; i< len(list); i++{
		j:=rand.Intn(i+1)
		list[i], list[j] = list[j], list[i]
	}
}