package n26

import (
	"time"
	"encoding/csv"
	"io"
	"strconv"
	"strings"
)

type Transaction struct {
	Date  time.Time
	Payee string
	Cents int
}

const dateLayout = "2006-01-02"

func ParseCSV(in io.Reader) (result []Transaction, err error) {
	r := csv.NewReader(in)
	r.LazyQuotes = true

	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	for i, fields := range records {
		if i == 0 {
			continue //skip header
		}
		result = append(result, parseFields(fields))
	}
	return
}

func parseFields(fields []string) Transaction {
	d, _ := time.Parse(dateLayout, strings.Trim(fields[0], " \""))
	p := strings.Trim(fields[1], " \"")
	c, _ := strconv.ParseFloat(strings.Trim(fields[6], " \""), 32)
	return Transaction{d, p, int(c * 100)}
}
