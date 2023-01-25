package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"golang.org/x/exp/slices"
	"log"
	"os"
	"strings"
)

func filterCSV(
	data [][]string,
	columns []string,
	isExclude bool,
) ([][]string, error) {
	var headers = data[0]

	var headerIndexes []int
	for index, header := range headers {
		var isContain = slices.Contains(columns, header)
		if isContain && !isExclude {
			headerIndexes = append(headerIndexes, index)
			continue
		}
		if !isExclude {
			continue
		}
		if isContain && isExclude {
			continue
		}
		headerIndexes = append(headerIndexes, index)
	}

	if len(headerIndexes) == 0 {
		return nil, errors.New("Empty csv")
	}

	var newCSV [][]string

	for _, row := range data {
		var rowArr []string
		for _, index := range headerIndexes {
			rowArr = append(rowArr, row[index])
		}
		newCSV = append(newCSV, rowArr)
	}

	return newCSV, nil
}

var file = ""
var comma = ``
var exclude = ""
var include = ""

func main() {

	flag.StringVar(&file, "file", "", "")
	flag.StringVar(&comma, "comma", ";", "")
	flag.StringVar(&exclude, "exclude", "", "")
	flag.StringVar(&include, "include", "", "")
	flag.Parse()

	if len(file) == 0 {
		log.Fatal("Command line argument file is empty")
	}

	var columns []string
	var isExclude = len(exclude) != 0

	if len(exclude) != 0 {
		columns = strings.Split(exclude, ",")
	} else if len(include) != 0 {
		columns = strings.Split(include, ",")
	} else {
		log.Fatal("Command line args \"exclude\" or \"include\" is empty")
	}

	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.Comma = []rune(comma)[0]
	data, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	newCSV, error := filterCSV(data, columns, isExclude)
	if error != nil {
		log.Fatal(error)
	}

	newCsvFileName := file + ".new"

	newCSVf, err := os.OpenFile(newCsvFileName, os.O_RDWR|os.O_CREATE, 0660)
	if err != nil {
		log.Fatal(err)
	}
	defer newCSVf.Close()

	csvWriter := csv.NewWriter(newCSVf)
	csvWriter.Comma = []rune(comma)[0]
	csvWriter.WriteAll(newCSV)
}
