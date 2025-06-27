package main

import (
	"bytes"
	"fyne_gold_watcher/repository"
	"io"
	"math"
	"net/http"
	"os"
	"testing"

	"fyne.io/fyne/v2/test"
)

var testApp Config

func TestMain(m *testing.M) {
	a := test.NewApp()
	testApp.App = a
	testApp.MainWindow = a.NewWindow("")
	testApp.HTTPClient = client
	testApp.DB = repository.NewTestRepository()
	os.Exit(m.Run())
}

var jsonToReturn = `
{
  "ts": 1750060473766,
  "tsj": 1750060465306,
  "date": "Jun 16th 2025, 03:54:25 am NY",
  "items": [
    {
      "curr": "USD",
      "xauPrice": 3412.025,
      "xagPrice": 36.4714,
      "chgXau": -21.44,
      "chgXag": 0.1534,
      "pcXau": -0.6244,
      "pcXag": 0.4224,
      "xauClose": 3433.465,
      "xagClose": 36.318
    }
  ]
}
`

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

var client = NewTestClient(func(req *http.Request) *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(jsonToReturn)),
		Header:     make(http.Header),
	}
})

const epsilon = 0.0001 // small value for floating point comparison

// FloatEqual compares 2 float64 values with the epsilon
func FloatEqual(a, b float64) bool {
	return math.Abs(a-b) < epsilon
}
