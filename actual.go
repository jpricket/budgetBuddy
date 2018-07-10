package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

// ActualItem is the smallest part of the actuals.
type ActualItem struct {
	Date        string
	Description string
	Amount      float64
}

// ActualGroup groups actual items.
type ActualGroup struct {
	BudgetName  string
	ActualItems []ActualItem
	TotalAmount float64
}

// The ActualMonth type represents a monthly attual document.
type ActualMonth struct {
	Month        string
	Year         int
	URL          string
	ActualGroups []ActualGroup
}

// The ActualYear type represents all the actual documents for a year.
type ActualYear struct {
	Year    int
	Actuals []ActualMonth
}

func (actuals ActualMonth) getFilePath() string {
	return getFilePathForMonth(strconv.Itoa(actuals.Year), actuals.Month, "_actuals")
}

func (actuals ActualMonth) writeFile() {
	filePath := actuals.getFilePath()
	dat, err := json.Marshal(actuals)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile(filePath, dat, 0644)
}

func getActualYear(year int) ActualYear {
	files, err := ioutil.ReadDir(getFilePathForYear(strconv.Itoa(year)))
	if err != nil {
		panic(err)
	}

	var actuals []ActualMonth
	for _, f := range files {
		filename := filepath.Base(f.Name())
		ext := filepath.Ext(filename)
		filenameNoExt := strings.TrimSuffix(filename, ext)
		if strings.HasSuffix(filenameNoExt, "_actuals") {
			filenameNoExt = strings.TrimSuffix(filenameNoExt, "_actuals")
			url := strconv.Itoa(year) + "/" + filenameNoExt
			b := ActualMonth{Year: year, Month: filenameNoExt, URL: url}
			actuals = append(actuals, b)
		}
	}

	actualYear := ActualYear{Year: year, Actuals: actuals}
	return actualYear
}

func readActualMonthFromFile(filePath string) ActualMonth {
	dat, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return readActualMonthFromBytes(dat)
}

func readActualMonth(reader io.Reader) ActualMonth {
	decoder := json.NewDecoder(reader)
	var actuals ActualMonth
	err := decoder.Decode(&actuals)
	if err != nil {
		panic(err)
	}
	actuals.resetTotals()
	return actuals
}

func readActualMonthFromBytes(jsonActuals []byte) ActualMonth {
	var actuals ActualMonth
	err := json.Unmarshal(jsonActuals, &actuals)
	if err != nil {
		panic(err)
	}
	actuals.resetTotals()
	return actuals
}

func (actuals ActualMonth) resetTotals() {
	for _, g := range actuals.ActualGroups {
		g.TotalAmount = 0
		for _, item := range g.ActualItems {
			g.TotalAmount += item.Amount
		}
	}
}
