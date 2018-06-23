package main

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/budgets", handleBudgets)
	r.HandleFunc("/budgets/{year}/{month}", handleBudget)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func handleBudget(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	year := vars["year"]
	month := vars["month"]
	data := readFile("c:\\budgets\\" + year + "\\" + month + ".json")
	budgetTemplate := template.Must(template.ParseFiles("view/budget.html"))
	budgetTemplate.Execute(w, data)
}

func handleBudgets(w http.ResponseWriter, r *http.Request) {
	budget := read(r.Body)

	//TODO save to file

	w.WriteHeader(http.StatusOK)
	byteSlice, _ := json.Marshal(budget)
	w.Write([]byte(byteSlice))
}
