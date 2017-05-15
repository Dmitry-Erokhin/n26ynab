package ynab

import (
	"os"
	"log"
	"regexp"
	"encoding/csv"
	"fmt"
	"os/user"
	"strings"
)

const rcFileName = ".n26ynab.csv"

var data = loadPayeeData(getPayeeConfigPath())

type payeeData struct {
	regexp   *regexp.Regexp
	payee    string
	category string
}

func MatchPayee(payee string) (string, string) {
	for _, entry := range data {
		if entry.regexp.MatchString(strings.ToLower(payee)) {
			return entry.payee, entry.category
		}
	}
	return "", ""
}

func loadPayeeData(path string)  (result []payeeData) {
	in, err := os.Open(path)
	defer in.Close()

	if err != nil {
		log.Printf("Can't open resource file %s. Payee matrching would not be avaliable.", rcFileName)
		return nil
	}
	r := csv.NewReader(in)
	r.LazyQuotes = true

	records, err := r.ReadAll()
	if err != nil {
		log.Printf("Error while loading resource filed. Payee matrching would not be avaliable. %s", err)
		return nil
	}

	for i, fields := range records {
		if i == 0 {
			continue //skip header
		}

		r, err := regexp.Compile(strings.Trim(fields[0]," "))

		if err != nil {
			log.Printf("Could not compile regexp %s on line %d. Will skip definition.", fields[0], i+1)
			continue
		}

		result = append(result, payeeData{
			regexp:   r,
			payee:    strings.Trim(fields[1]," "),
			category: strings.Trim(fields[2]," "),
		})
	}

	return
}

func getPayeeConfigPath() string {
	usr, err := user.Current()
	if err != nil {
		log.Println("Can't determinate current user. Payee matching would not be avaliable.")
		return ""
	}
	fmt.Println(usr.HomeDir)

	return  usr.HomeDir + string(os.PathSeparator) + rcFileName
}