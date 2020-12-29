package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	//"github.com/gorilla/mux"
)

var balance = 9563
var count = 0 //

func main() {
	handlerequests()
}

// Amount exported
var Amount int

func choices(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		showbalance(w, r)
		return
	case "POST":
		testfunc(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}
}
func showbalance(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("here is your balance"))
	//w.Write([]byte("here is your balance"))
}
func testfunc(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "using post command")
	fmt.Println("count is:", count) //
	bodybytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	var input int
	json.Unmarshal(bodybytes, &input)
	fmt.Println(input)
	fmt.Printf("\n%T\n", input)
	j := fmt.Sprintf("\nyou entered %d:\n", input)
	w.Write([]byte(j))
	if count >= 5 {
		fmt.Println("Error! you have made maximum number of transactions for today: 5. Initiating Exit")
		w.Write([]byte("Maximum transactions reached for a day:5 "))
		return //
	}
	if balance < 100 {
		w.Write([]byte("Balance less than 100.You can't withdraw "))
		return //
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
		j := fmt.Sprintf("Error!Please enter amount less than or equal to your account balance that is: %v\n", balance)
		w.Write([]byte(j))
		break
	default:
		count++
		i := denoms(amount)
		fmt.Println(i)
		balance = balance - amount
		fmt.Println("TRANSACTION SUCCESSFUL.YOUR BALANCE NOW IS:", balance)
		j := fmt.Sprintf(" \t TRANSACTION SUCCESSFUL.YOUR BALANCE NOW IS: %d\t ", balance)
		w.Write([]byte(i))
		w.Write([]byte(j))
		break
	}
	w.Write([]byte("Do you want to start another transaction?Enter y to continue and n to exit"))

	foo(w, r)

} //

func handlerequests() {
	//myrouter := mux.NewRouter().StrictSlash(true)

	http.HandleFunc("/", choices)
	//http.HandleFunc("/article", allArticles)
	//http.HandleFunc("/article", testfunc)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func amt(input int, balance int) (int, int) {

	//input, err := strconv.ParseInt(y, 10, 64)
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

func foo(w http.ResponseWriter, r *http.Request) {
	bodyfoo, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	var input string
	json.Unmarshal(bodyfoo, &input)
	fmt.Println(input)
	fmt.Printf("\n  type of input in foo:%T\n", input)
	j := fmt.Sprintf("\nyou entered: %s\n", input) //
	w.Write([]byte(j))

	switch input {
	default:
		{
			w.Write([]byte("Server says:Error!Wrong input.You can only Enter y or n"))
			foo(w, r)
			return
		}

	case "y":
		{
			w.Write([]byte("You chose:yes, start a new transaction.Please enter the amount"))
			return
		}
	case "n":
		{
			w.Write([]byte("exitnow"))
			return //user can still check balance, so dont exit here
		}
	}
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
