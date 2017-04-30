package ynab

import (
	"time"
	"io"
	"encoding/csv"
	"strconv"
	"strings"
	"log"
)

const header = "Date,Payee,Category,Memo,Outflow,Inflow"
const dateLayout = "2006-01-02"

type BudgetEntry struct {
	Date    time.Time
	Payee   string
	Memo    string
	Outflow float32
	Inflow  float32
}

func WriteCsv(entries []BudgetEntry, out io.Writer) {
	w := csv.NewWriter(out)
	if len(entries) < 1 {
		return
	}
	out.Write([]byte(header + "\n"))
	w.WriteAll(renderRecords(entries))
}

func renderRecords(entries []BudgetEntry) (result [][]string) {
	deleteComas := func(s string) string {
		return strings.Replace(s, ",", "", -1)
	}
	for _, e := range entries {
		var in, out string
		if e.Outflow != 0 {
			out = strconv.FormatFloat(float64(e.Outflow), 'f', 2, 32)
		} else if e.Inflow != 0 {
			in = strconv.FormatFloat(float64(e.Inflow), 'f', 2, 32)
		}

		result = append(result, []string{
			e.Date.Format(dateLayout),
			deleteComas(e.Payee),
			"",
			deleteComas(e.Memo),
			out,
			in,
		})
	}

	log.Printf("Saved %d YNAB entries", len(entries))
	return
}
