package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
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
	BudgetGroups []BudgetGroup
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
