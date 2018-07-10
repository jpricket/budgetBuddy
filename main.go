package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// WritableData is data that is written to a file
type WritableData interface {
	getFilePath() string
	writeFile()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/budgets", handleBudgets)
	r.HandleFunc("/budgets/{year}", handleBudgetYear)
	r.HandleFunc("/budgets/{year}/{month}", handleBudget)
	r.HandleFunc("/actuals", handleActuals)
	r.HandleFunc("/actuals/{year}", handleActualsYear)
	r.HandleFunc("/actuals/{year}/{month}", handleActualMonth)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func handleBudget(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	year := vars["year"]
	month := vars["month"]
	data := readFile(getFilePathForMonth(year, month, ""))
	budgetTemplate := template.Must(template.ParseFiles("view/budget.html"))
	budgetTemplate.Execute(w, data)
}

func handleBudgetYear(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	year, _ := strconv.Atoi(vars["year"])
	data := getBudgetYear(year)
	budgetTemplate := template.Must(template.ParseFiles("view/budgetYear.html"))
	budgetTemplate.Execute(w, data)
}

func handleBudgets(w http.ResponseWriter, r *http.Request) {
	budget := read(r.Body)

	budget.writeFile()

	w.WriteHeader(http.StatusOK)
	byteSlice, _ := json.Marshal(budget)
	w.Write([]byte(byteSlice))
}

func handleActualMonth(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	year := vars["year"]
	month := vars["month"]
	data := readActualMonthFromFile(getFilePathForMonth(year, month, "_actuals"))
	actualsTemplate := template.Must(template.ParseFiles("view/actualMonth.html"))
	actualsTemplate.Execute(w, data)
}

func handleActualsYear(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	year, _ := strconv.Atoi(vars["year"])
	data := getActualYear(year)
	yearTemplate := template.Must(template.ParseFiles("view/actualsYear.html"))
	yearTemplate.Execute(w, data)
}

func handleActuals(w http.ResponseWriter, r *http.Request) {
	actualMonth := readActualMonth(r.Body)

	actualMonth.writeFile()

	w.WriteHeader(http.StatusOK)
	byteSlice, _ := json.Marshal(actualMonth)
	w.Write([]byte(byteSlice))
}

func getFilePathForMonth(year string, month string, suffix string) string {
	return getFilePathForYear(year) + "\\" + month + suffix + ".json"
}

func getFilePathForYear(year string) string {
	return "c:\\budgets\\" + year
}
