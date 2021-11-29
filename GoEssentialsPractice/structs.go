package main

import (
	"fmt"
)

func main() {
	t1 := trade{Symbol: "MSFT", Volume: 100, Price: 5, buy: false}
	fmt.Println(t1)

	t2 := trade{}
	t2.Symbol = "TSLA"
	t2.Price = 1000
	fmt.Println(t2)

	t3 := trade{"AMZN", 1000, 0.5, false}
	fmt.Println(t3)

	t3.BuyStock()
	fmt.Println(t3)
	fmt.Println(t3.Value())

	t4, err := NewTrade("TSLA", 30000, 100, true)
	if err == nil {
		fmt.Println(*t4)
	}

}

type trade struct {
	Symbol string
	Volume int
	Price  float64
	buy    bool
}

func (t *trade) BuyStock() {
	t.buy = true
}

func (t trade) Value() float64 {
	value := float64(t.Volume) * t.Price
	if t.buy {
		value = -value
	}
	return value
}

func NewTrade(Symbol string, Volume int, Price float64, buy bool) (*trade, error) {
	if Symbol == "" {
		return nil, fmt.Errorf("Symbol cannot be empty")
	}

	if Volume < 0 || Price < 0 {
		return nil, fmt.Errorf("Volumen or price cannot be negative")
	}

	t := trade{Symbol, Volume, Price, buy}

	return &t, nil
}
