package main

import (
	"database/sql"
	"fyne_gold_watcher/repository"
	"log"
	"net/http"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"

	_ "github.com/glebarez/go-sqlite"
)

// Config is configuration for fyne app
type Config struct {
	App                           fyne.App
	InfoLog                       *log.Logger
	ErrorLog                      *log.Logger
	DB                            repository.Repository
	MainWindow                    fyne.Window
	PriceContainer                *fyne.Container
	Toolbar                       *widget.Toolbar
	PriceChartContainer           *fyne.Container
	Holdings                      [][]interface{}
	HoldingsTable                 *widget.Table
	HTTPClient                    *http.Client
	AddHoldingsAmountEntry        *widget.Entry
	AddHoldingsPurchaseDateEntry  *widget.Entry
	AddHoldingsPurchasePriceEntry *widget.Entry
}

var myApp Config

func main() {
	// create a fyne app
	fyneApp := app.NewWithID("com.nvbien.goldWatcher")
	myApp.App = fyneApp
	myApp.HTTPClient = &http.Client{}

	// create our loggers
	myApp.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	myApp.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// open a connection to database
	sqlDB, err := myApp.connectSQL()
	if err != nil {
		log.Panic(err)
	}

	// create a database repository
	myApp.setupDB(sqlDB)

	// create & size a fyne window
	myApp.MainWindow = fyneApp.NewWindow("Gold Watcher")
	myApp.MainWindow.Resize(fyne.NewSize(770, 410))
	myApp.MainWindow.SetFixedSize(true)
	myApp.MainWindow.SetMaster()

	myApp.makeUI()

	// show & run
	myApp.MainWindow.ShowAndRun()
}

func (app *Config) connectSQL() (*sql.DB, error) {
	path := ""
	if os.Getenv("DB_PATH") != "" {
		path = os.Getenv("DB_PATH")
	} else {
		path = app.App.Storage().RootURI().Path() + "/sql.db"
		app.InfoLog.Println("db in:", path)
	}

	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (app *Config) setupDB(sqlDB *sql.DB) {
	app.DB = repository.NewSQLiteRepository(sqlDB)

	err := app.DB.Migrate()
	if err != nil {
		app.ErrorLog.Println(err)
		log.Panic()
	}
}
