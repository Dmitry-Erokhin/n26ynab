package n26

import (
	"testing"
	"time"
	"strings"
)

func TestParseCSV(t *testing.T) {
	input := `"Date","Payee","Account number","Transaction type","Payment reference","Category","Amount (EUR)","Amount (Foreign Currency)","Type Foreign Currency","Exchange Rate"
"2017-04-18","www.bvg.de","","MasterCard Payment","","Transport & Car","-9.0","-9.0","EUR","1.0"
"2017-04-30","DM-Drogerie, Markt","","MasterCard Payment","","Healthcare & Drug Stores","7.4","7.4","EUR","1.0"`
	expected := []Transaction{
		{
			time.Date(2017, 04, 18, 0, 0, 0, 0, time.UTC),
			"www.bvg.de",
			-900,
		},
		{
			time.Date(2017, 04, 30, 0, 0, 0, 0, time.UTC),
			"DM-Drogerie, Markt",
			740,
		},
	}

	r := strings.NewReader(input)

	actual, err := ParseCSV(r)

	if err != nil {
		t.Error("Did not processed reading")
		t.FailNow()
	}

	if len(actual) != 2 {
		t.Error("Processed lines conut does not match expectation")
		t.FailNow()
	}

	for i := range actual {
		if actual[i] != expected[i] {
			t.Errorf("Expected: %v, bot got : %v", expected[i], actual[i])
		}
	}
}
