package budgeter

import (
	"testing"
	"strings"
	"bytes"
)

func TestConvertData(t *testing.T) {
	input := `"Date","Payee","Account number","Transaction type","Payment reference","Category","Amount (EUR)","Amount (Foreign Currency)","Type Foreign Currency","Exchange Rate"
"2017-04-18","www.bvg.de","","MasterCard Payment","","Transport & Car","-9.0","-9.0","EUR","1.0"
"2017-04-17","Dmitry Erokhin","DE87100110012629884849","Income","Hey, I ve just sent you a MoneyBeam from my N26 Bank account. Learn more on n26.com and try it yourself.","Income","100.0","","",""
`
	expected := `Date,Payee,Category,Memo,Outflow,Inflow
2017-04-18,,,www.bvg.de,9.00,
2017-04-17,,,Dmitry Erokhin,,100.00
`

	r := strings.NewReader(input)
	w := bytes.NewBufferString("")

	err := ConvertData(r, w)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	actual := w.String()
	if expected != actual {
		t.Errorf("Expected:\n=========\n%v\n=========\nBut got:\n=========\n%v\n=========\n", expected, actual)

	}
}
