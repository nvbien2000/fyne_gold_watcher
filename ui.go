package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func (app *Config) makeUI() {
	// get the current price of gold
	openPrice, currPrice, changePrice := app.getPriceText()

	// put the price information into a container
	priceContent := container.NewGridWithColumns(3, openPrice, currPrice, changePrice)
	app.PriceContainer = priceContent

	// get toolbar
	toolbar := app.getToolBar()
	app.Toolbar = toolbar

	priceTabContent := app.pricesTab()
	holdingsTabContent := app.holdingsTab()

	// get app tabs
	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Prices", theme.HomeIcon(), priceTabContent),
		container.NewTabItemWithIcon("Holdings", theme.InfoIcon(), holdingsTabContent),
	)
	tabs.SetTabLocation(container.TabLocationTop)

	// add container to window
	finalContent := container.NewVBox(priceContent, toolbar, tabs)

	app.MainWindow.SetContent(finalContent)

	go func() {
		for range time.Tick(time.Second * 10) {
			fmt.Println("10s passed: refresh...")
			fyne.Do(func() {
				app.refreshPriceContent()
			})
		}
	}()
}

func (app *Config) refreshPriceContent() {
	open, curr, change := app.getPriceText()
	app.PriceContainer.Objects = []fyne.CanvasObject{open, curr, change}
	app.PriceContainer.Refresh()

	chart := app.getChart()
	app.PriceChartContainer.Objects = []fyne.CanvasObject{chart}
	app.PriceChartContainer.Refresh()
}
