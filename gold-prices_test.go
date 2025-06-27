package main

import (
	"testing"
)

func TestGold_GetPrices(t *testing.T) {
	g := Gold{
		Items:  nil,
		Client: client,
	}

	p, err := g.GetPrices()
	if err != nil {
		t.Error(err)
	}

	if !FloatEqual(p.GoldPrice, 3412.025) {
		t.Error("wrong gold price returned: ", p.GoldPrice)
	}
}
