package safe

import (
	"strconv"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type Bill uint

func StrToBill(s string) (*Bill, error) {
	num, err := strconv.Atoi(s)
	if err != nil {
		return nil, err
	}
	bill := Bill(num * 100)
	return &bill, nil
}

type BillBunch struct {
	Bill
	Count uint
}

func (bb BillBunch) Sum() uint {
	return bb.Count * uint(bb.Bill)
}

type BillBunches []BillBunch

func (bbs BillBunches) Sum() uint {
	sum := uint(0)
	for _, bb := range bbs {
		sum += bb.Sum()
	}
	return sum
}

func (bbs BillBunches) String() string {
	p := message.NewPrinter(language.English)
	return p.Sprintf("%.2f", float64(bbs.Sum())/100)
}
