# #3 Wire Transfer

Wire Transfer is a Golang project where customers can:

1. Wire Transactions can be issued via a Browser or terminal command-line interface

2. User will get a status whether the funds were transferred or insufficient funds to transfer

3. NO business logic in the transfer server.

4. No Database is needed. You can hard code the accounts in a file

5. Record successful and unsuccessful transactions to a file

6. You can add account holders [Name, Acct #, Opening Balance]

## Tools

[golang](https://go.dev/dl/go1.19.darwin-amd64.pkg) to install golang and check for go version.

[gorilla mux](https://github.com/gorilla/mux) to install gorilla mux.

[postman](https://www.postman.com/downloads/) to install postman.

```bash
go version
```

## Code 

```golang
import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

# Get All Accounts
func getAccounts(w http.ResponseWriter, r *http.Request)

# Delete Account
func deleteAccount(w http.ResponseWriter, r *http.Request)

# Get An Account
func getAccount(w http.ResponseWriter, r *http.Request)

# Create An Account
func createAccount(w http.ResponseWriter, r *http.Request)

# Update An Account
func updateAccount(w http.ResponseWriter, r *http.Request)


# main router
r := mux.NewRouter()

	accounts = append(accounts, Account{CustomerID: "2021194v", SWIFT: "JD21-0676", 
	Balance: 8889, Account_Holder: &Account_Holder{Firstname: "Chealyfey", Lastname: "Vutha", Address: "Singapore"}})
	accounts = append(accounts, Account{CustomerID: "2021074v", SWIFT: "JD21-0677", 
	Balance: 2300, Account_Holder: &Account_Holder{Firstname: "David", Lastname: "Vicheth", Address: "Malaysia"}})
	accounts = append(accounts, Account{CustomerID: "2021093t", SWIFT: "JDXX-0000", 
	Balance: 1110, Account_Holder: &Account_Holder{Firstname: "Vesal", Lastname: "Thong", Address: "Cambodia"}})
	
	r.HandleFunc("/accounts", getAccounts).Methods("GET")
	r.HandleFunc("/accounts/{customerid}", getAccount).Methods("GET")
	r.HandleFunc("/accounts", createAccount).Methods("POST")
	r.HandleFunc("/accounts/{customerid}", updateAccount).Methods("PUT")
	r.HandleFunc("/accounts/{customerid}", deleteAccount).Methods("DELETE")
```

## Contribution
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[LYFEYTECH](https://github.com/lyfeytech)
