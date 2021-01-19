package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	database    = "YOUR_DB"
	collection1 = "WITHDRAWAL_COLLECTION"
	collection2 = "BALANCE_COLLECTION"
)

type userBalance struct {
	ID      bson.ObjectId `bson:"_id,omitempty"`
	Balance int           `bson:"balance"`
}
type withdrawalsTable struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Amount    int           `bson:"amount"`
	TimeStamp time.Time     `bson:"Timestamp"`
}
type withdrawValue struct {
	Value int `json:"value"`
}

func isObjectIDValid(i bson.ObjectId) error {
	j := i.Valid()
	if j == false {
		err := fmt.Errorf("Invalid Object ID")
		return err
	}
	return nil
}
func main() {
	transactionCount := 0
	var results1 []userBalance
	var v userBalance

	Host := []string{
		"127.0.0.1:27017",
	}

	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs: Host,
	})
	if err != nil {
		fmt.Println("Error in Mongo session", err)
	}
	defer session.Close()
	id1 := bson.NewObjectId()
	errinID := isObjectIDValid(id1)
	if errinID != nil {
		log.Fatalln(errinID)
	}
	balanceTable := session.DB(database).C(collection2)
	allWithdrawals := session.DB(database).C(collection1)

	balanceTable.RemoveAll(nil)
	//allWithdrawals.RemoveAll(nil)

	//Inserting initial value in the balance table
	if err := balanceTable.Insert(&userBalance{ID: id1, Balance: 9563}); err != nil {
		fmt.Println("Error in inserting balance:", err)
	}
	err = balanceTable.Find(nil).Select(bson.M{"balance": 1}).All(&results1)
	if err != nil {
		fmt.Println("Error in finding balance field:", err)
	}
	for _, v = range results1 {
	}

	server := gin.Default()

	server.POST("/", func(c *gin.Context) {
		cI := withdrawValue{}
		c.BindJSON(&cI)

		clientAmount, showMessage := withdrawAmount(cI.Value, v.Balance, &transactionCount)
		if clientAmount != 0 {
			if err := allWithdrawals.Insert(&withdrawalsTable{Amount: clientAmount, TimeStamp: time.Now()}); err != nil {
				fmt.Println("Error in inserting amount:", err)
			}
		}
		//Updating the balance table
		selector := bson.M{"_id": id1}
		updator := bson.M{"$inc": bson.M{"balance": -clientAmount}}
		if err := balanceTable.Update(selector, updator); err != nil {
			fmt.Println("Error in updating balance:", err)
		}

		err = balanceTable.Find(nil).All(&results1)
		if err != nil {
			fmt.Println(err)
		}
		for _, v = range results1 {
		}
		c.JSON(200, gin.H{"Result": showMessage})
	})

	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"Your balance is": v.Balance})
	})

	server.POST("/exit", func(c *gin.Context) {
		os.Exit(1)
	})

	err = server.Run(":8080") //  listen for incoming connections
	if err != nil {
		fmt.Println(err)
	}
}
func withdrawAmount(clientVal int, userbalance int, counterPtr *int) (int, string) {
	if *counterPtr >= 5 {
		fmt.Println("Error! you have made maximum number of transactions for today: 5.")
		message := fmt.Sprintf("Maximum transactions reached for a day:5. You can exit using POST/exit.")
		return 0, message
	}
	if userbalance < 100 {
		message := fmt.Sprintf("Balance less than 100.You can't withdraw anymore.You can exit using POST/exit.")
		return 0, message
	}
	checkValue, amount := isAmountValid(clientVal, userbalance)
	switch checkValue {
	case 0:
		message := ("Error!Please enter ONE natural number.")
		return 0, message
	case 10:
		message := fmt.Sprintf("Error!Please enter a Positive value ")
		return 0, message
	case 20:
		message := fmt.Sprintf("Error!Please enter amount divisible by 100 ")
		return 0, message
	case 30:
		message := fmt.Sprintf("Error! Please enter amount less than or equal to 5000 ")
		return 0, message
	case 40:
		message := fmt.Sprintf("Error!Please enter amount less than or equal to your account balance")
		return 0, message
	default:
		*counterPtr++
		i := getDenominations(amount)
		userbalance = userbalance - amount
		j := fmt.Sprintf("TRANSACTION SUCCESSFUL.")
		message := i + j
		return amount, message
	}
}
func isAmountValid(input int, balance int) (int, int) {
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
func getDenominations(total int) string {
	var q1 int
	var q2 int
	var q3 int
	q1 = total / 500
	total = total % 500
	q2 = total / 200
	total = total % 200
	q3 = total / 100
	fmt.Printf("500*%d +200*%d +100*%d \n", q1, q2, q3)
	i := fmt.Sprintf("500*%d +200*%d +100*%d ", q1, q2, q3)
	return i
}
