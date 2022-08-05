package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Account struct {
	CustomerID     string          `json:"customerid"`
	SWIFT          string          `json:"swift"`
	Balance        float64         `json:"balance"`
	Account_Holder *Account_Holder `json:"account_holder"`
}

type Account_Holder struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Address   string `json:"address"`
}

var accounts []Account

func getAccounts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accounts)
}

func deleteAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range accounts {

		if item.CustomerID == params["customerid"] {
			accounts = append(accounts[:index], accounts[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(accounts)
}

func getAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range accounts {
		if item.CustomerID == params["customerid"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var account Account
	_ = json.NewDecoder(r.Body).Decode(&account)
	account.CustomerID = strconv.Itoa(rand.Intn(1000000000))
	accounts = append(accounts, account)
	json.NewEncoder(w).Encode(account)
}

func updateAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range accounts {
		if item.CustomerID == params["customerid"] {
			accounts = append(accounts[:index], accounts[index+1:]...)
			var account Account
			_ = json.NewDecoder(r.Body).Decode(&account)
			account.CustomerID = params["customerid"]
			accounts = append(accounts, account)
			json.NewEncoder(w).Encode(account)
			return
		}
	}
}
func main() {
	r := mux.NewRouter()

	accounts = append(accounts, Account{CustomerID: "2021194v", SWIFT: "JD21-0676", Balance: 8889, 
		Account_Holder: &Account_Holder{Firstname: "Chealyfey", Lastname: "Vutha", Address: "Singapore"}})
	accounts = append(accounts, Account{CustomerID: "2021074v", SWIFT: "JD21-0677", Balance: 2300, 
		Account_Holder: &Account_Holder{Firstname: "David", Lastname: "Vicheth", Address: "Malaysia"}})
	accounts = append(accounts, Account{CustomerID: "2021093t", SWIFT: "JDXX-0000", Balance: 6942, 
		Account_Holder: &Account_Holder{Firstname: "Vesal", Lastname: "Thong", Address: "Cambodia"}})
	
	r.HandleFunc("/accounts", getAccounts).Methods("GET")
	r.HandleFunc("/accounts/{customerid}", getAccount).Methods("GET")
	r.HandleFunc("/accounts", createAccount).Methods("POST")
	r.HandleFunc("/accounts/{customerid}", updateAccount).Methods("PUT")
	r.HandleFunc("/accounts/{customerid}", deleteAccount).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
