package ynab

import (
	"testing"
	"bytes"
	"time"
)

func TestWriteCsv(t *testing.T) {
	input := []BudgetEntry{{
		time.Date(2017, 04, 18, 0, 0, 0, 0, time.UTC),
		"",
		"www.bvg.de",
		9.42,
		0}, {

		time.Date(2017, 04, 18, 0, 0, 0, 0, time.UTC),
		"",
		"Edeka, Kaisers",
		0,
		100},
	}

	expected := header + `
2017-04-18,,,www.bvg.de,9.42,
2017-04-18,,,Edeka Kaisers,,100.00
`
	b := bytes.NewBufferString("")
	WriteCsv(input, b)
	actual := b.String()
	if expected != actual {
		t.Errorf("Expected:\n=========\n%v\n=========\nBut got:\n=========\n%v\n=========\n", expected, actual)
	}
}
