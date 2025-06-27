package main

import (
	"testing"

	"fyne.io/fyne/v2/test"
)

func TestApp_getToolbar(t *testing.T) {
	tb := testApp.getToolBar()
	if len(tb.Items) != 4 {
		t.Error("Wrong number of items in toolbar")
	}
}

func TestApp_addHoldingsDialog(t *testing.T) {
	testApp.addHoldingsDialog()

	test.Type(testApp.AddHoldingsAmountEntry, "1")
	test.Type(testApp.AddHoldingsPurchasePriceEntry, "1000")
	test.Type(testApp.AddHoldingsPurchaseDateEntry, "2020-01-01")

	if testApp.AddHoldingsAmountEntry.Text != "1" {
		t.Error("amount not correct")
	}
	if testApp.AddHoldingsPurchaseDateEntry.Text != "2020-01-01" {
		t.Error("date not correct")
	}
	if testApp.AddHoldingsPurchasePriceEntry.Text != "1000" {
		t.Error("price not correct")
	}
}
