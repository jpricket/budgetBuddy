package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

// BudgetItem is the smallest part of the budget.
type BudgetItem struct {
	Name         string
	BudgetAmount float64
	ActualAmount float64
}

// BudgetGroup groups budget items.
type BudgetGroup struct {
	Name        string
	Incoming    bool
	BudgetItems []BudgetItem
}

// The Budget type represents a monthly budget document.
type Budget struct {
	Month        string
	Year         int
	URL          string
	BudgetGroups []BudgetGroup
}

// The BudgetYear type represents all the budget documents for a year.
type BudgetYear struct {
	Year    int
	Budgets []Budget
}

func (budget Budget) getFilePath() string {
	return getFilePathForMonth(strconv.Itoa(budget.Year), budget.Month, "")
}

func (budget Budget) writeFile() {
	filePath := budget.getFilePath()
	dat, err := json.Marshal(budget)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile(filePath, dat, 0644)
}

func getBudgetYear(year int) BudgetYear {
	files, err := ioutil.ReadDir(getFilePathForYear(strconv.Itoa(year)))
	if err != nil {
		panic(err)
	}

	var budgets []Budget
	for _, f := range files {
		filename := filepath.Base(f.Name())
		ext := filepath.Ext(filename)
		filenameNoExt := strings.TrimSuffix(filename, ext)
		if !strings.HasSuffix(filenameNoExt, "_actuals") {
			url := strconv.Itoa(year) + "/" + filenameNoExt
			b := Budget{Year: year, Month: filenameNoExt, URL: url}
			budgets = append(budgets, b)
		}
	}

	budgetYear := BudgetYear{Year: year, Budgets: budgets}
	return budgetYear
}

func readFile(filePath string) Budget {
	dat, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return readBytes(dat)
}

func read(reader io.Reader) Budget {
	decoder := json.NewDecoder(reader)
	var budget Budget
	err := decoder.Decode(&budget)
	if err != nil {
		panic(err)
	}
	return budget
}

func readBytes(jsonBudget []byte) Budget {
	var budget Budget
	err := json.Unmarshal(jsonBudget, &budget)
	if err != nil {
		panic(err)
	}
	return budget
}
