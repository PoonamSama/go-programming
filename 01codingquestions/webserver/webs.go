package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var balance = 9563
var count = 0

func main() {
	handlerequests()
}
func choices(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		showbalance(w, r)
		return
	case "POST":
		testfunc(w, r)
		return
	case "PUT":
		funcput()
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}
}
func funcput() {
	fmt.Println("Initiating Exit")
	os.Exit(1)
}
func showbalance(w http.ResponseWriter, r *http.Request) {
	j := fmt.Sprintf("Your balance: %d\n", balance)
	w.Write([]byte(j))
}
func testfunc(w http.ResponseWriter, r *http.Request) {
	bodybytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	var input int
	json.Unmarshal(bodybytes, &input)
	if count >= 5 {
		fmt.Println("Error! you have made maximum number of transactions for today: 5.Exit using PUT/")
		w.Write([]byte("Maximum transactions reached for a day:5. You can exit using PUT/ and check balance using GET/ . "))
		return
	}
	if balance < 100 {
		w.Write([]byte("Balance less than 100.You can't withdraw. You can exit using PUT/ and find your balance using GET/ . "))
		return
	}
	x, amount := amt(input, balance)
	switch x {
	case 0:
		fmt.Println("Error!Please enter a natural number")
		w.Write([]byte("Error!Please enter a natural number \n"))
		break

	case 10:
		fmt.Println("Error!Please enter a Positive value")
		w.Write([]byte("Error!Please enter a Positive value \n"))
		break
	case 20:
		fmt.Println("Error!Please enter amount divisible by 100 ")
		w.Write([]byte("Error!Please enter amount divisible by 100 \n"))
		break
	case 30:
		fmt.Println("Error! Please enter amount less than or equal to 5000")
		w.Write([]byte("Error! Please enter amount less than or equal to 5000 \n"))
		break
	case 40:
		fmt.Printf("Error!Please enter amount less than or equal to your account balance that is: %v\n", balance)
		j := fmt.Sprintf("Error!Please enter amount less than or equal to your account balance \n")
		w.Write([]byte(j))
		break
	default:
		count++
		i := denoms(amount)
		balance = balance - amount
		fmt.Println("\n TRANSACTION SUCCESSFUL.YOUR BALANCE NOW IS:\n", balance)
		j := fmt.Sprintf("\n TRANSACTION SUCCESSFUL.\n ")
		w.Write([]byte(i))
		w.Write([]byte(j))
		break
	}
	w.Write([]byte("Do you want to start another transaction?Enter amount and Send to continue another transaction. You can use PUT/ to exit and GET/ to get your balance."))

}

func handlerequests() {
	http.HandleFunc("/", choices)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func amt(input int, balance int) (int, int) {
	if input == 0 {
		return 0, input
	}
	if input < 0 {
		return 10, input
	}
	if input%100 != 0 {
		return 20, input
	}
	if input > 5000 {
		return 30, input
	}
	if balance-input < 0 {
		return 40, input
	}
	return 50, input
}

func denoms(x int) string {
	var q1 int
	var q2 int
	var q3 int
	q1 = x / 500
	x = x % 500
	q2 = x / 200
	x = x % 200
	q3 = x / 100
	fmt.Printf("500*%d +200*%d +100*%d \n", q1, q2, q3)
	i := fmt.Sprintf("500*%d +200*%d +100*%d ", q1, q2, q3)
	return i
}
