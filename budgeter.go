package budgeter

import (
	"io"
	"github.com/Dmitry-Erokhin/n26ynab/n26"
	"github.com/Dmitry-Erokhin/n26ynab/ynab"
	"log"
)

func ConvertData(r io.Reader, w io.Writer) error {
	txns, err := n26.ParseCSV(r)

	if err != nil {
		return err
	}

	log.Printf("Loaded %d N26 entries", len(txns))

	var entries []ynab.BudgetEntry
	for _, txn := range txns {
		var in, out float32
		if txn.Cents < 0 {
			out = - float32(txn.Cents) / 100
		} else {
			in = float32(txn.Cents) / 100
		}
		payee, _ := ynab.MatchPayee(txn.Payee)
		entries = append(entries, ynab.BudgetEntry{txn.Date, payee, txn.Payee, out, in})
	}

	ynab.WriteCsv(entries, w)

	return nil
}
