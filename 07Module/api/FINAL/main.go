package main

import (
	"fmt"
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

func isObjectIDValid(id bson.ObjectId) error {
	j := id.Valid()
	if j == false {
		err := fmt.Errorf("Invalid Object ID")
		return err
	}
	return nil
}
func main() {
	var balanceDetails []userBalance
	var errInBalance error
	var errInWithdrawal error

	Host := []string{
		"127.0.0.1:27017",
	}

	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs: Host,
	})
	if err != nil {
		fmt.Println("Error in Mongo session", err)
		os.Exit(1)
	}
	defer session.Close()
	balanceID := bson.NewObjectId()
	errinBalanceID := isObjectIDValid(balanceID)
	if errinBalanceID != nil {
		fmt.Println(errinBalanceID)
	}
	transactionID := bson.NewObjectId()
	errinTransactionID := isObjectIDValid(transactionID)
	if errinTransactionID != nil {
		fmt.Println(errinTransactionID)
	}
	balanceTable := session.DB(database).C(collection2)
	allWithdrawals := session.DB(database).C(collection1)

	//balanceTable.RemoveAll(nil)
	//allWithdrawals.RemoveAll(nil)
	presentTime := time.Now()
	timeDifference := presentTime.Add(time.Hour * (-24))
	numOfWithdrawals, err := allWithdrawals.Find(bson.M{
		"Timestamp": bson.M{
			"$gt": timeDifference,
		},
	}).Count()
	if err != nil {
		fmt.Println(err)
	}
	//Inserting initial value in the balance table
	if errinBalanceID == nil {
		if errInBalance = balanceTable.Insert(&userBalance{ID: balanceID, Balance: 9563}); errInBalance != nil {
			fmt.Println("Error in inserting balance:", err)
		}
	}
	errFindingBalance := balanceTable.FindId(balanceID).Select(bson.M{"balance": 1}).All(&balanceDetails)
	if errFindingBalance != nil {
		fmt.Println("Error in finding balance field:", errFindingBalance)
	}

	server := gin.Default()

	server.POST("/", func(c *gin.Context) {
		clientInput := withdrawValue{}
		c.BindJSON(&clientInput)
		clientAmount, showMessage := withdrawAmount(clientInput.Value, balanceDetails[0].Balance, numOfWithdrawals)
		if errinTransactionID == nil && clientAmount != 0 {
			if errInWithdrawal = allWithdrawals.Insert(&withdrawalsTable{ID: transactionID, Amount: clientAmount, TimeStamp: time.Now()}); errInWithdrawal != nil {
				fmt.Println("Error in inserting amount:", errInWithdrawal)
			}
		}
		//Updating the balance table
		selector := bson.M{"_id": balanceID}
		updator := bson.M{"$inc": bson.M{"balance": -clientAmount}}
		if errInWithdrawal == nil && errinBalanceID == nil {
			if errInBalance = balanceTable.Update(selector, updator); errInBalance != nil {
				fmt.Println("Error in updating balance:", errInBalance)
				allWithdrawals.Remove(bson.M{"_id": transactionID})
			}
		}
		errFindingBalance = balanceTable.FindId(balanceID).Select(bson.M{"balance": 1}).All(&balanceDetails)
		if errFindingBalance != nil {
			fmt.Println(errFindingBalance)
		}
		c.JSON(200, gin.H{"Result": showMessage})

		transactionID = bson.NewObjectId()
		errinTransactionID = isObjectIDValid(transactionID)
		if errinTransactionID != nil {
			fmt.Println(errinTransactionID)
		}
		presentTime := time.Now()
		timeDifference := presentTime.Add(time.Hour * (-24))
		numOfWithdrawals, err = allWithdrawals.Find(bson.M{
			"Timestamp": bson.M{
				"$gt": timeDifference,
			},
		}).Count()
		if err != nil {
			fmt.Println(err)
		}
	})

	server.GET("/", func(c *gin.Context) {
		if errFindingBalance == nil {
			c.JSON(200, gin.H{"Your balance is": balanceDetails[0].Balance})
		}
	})

	server.POST("/exit", func(c *gin.Context) {
		os.Exit(1)
	})

	err = server.Run(":8080") //  listen for incoming connections
	if err != nil {
		fmt.Println(err)
	}
}
func withdrawAmount(clientValue int, userbalance int, withdrawalCount int) (int, string) {
	if withdrawalCount >= 5 {
		fmt.Println("Error! you have made maximum number of transactions for today: 5.")
		message := fmt.Sprintf("Maximum transactions reached for a day:5. You can exit using POST/exit.")
		return 0, message
	}
	if userbalance < 100 {
		message := fmt.Sprintf("Balance less than 100.You can't withdraw anymore.You can exit using POST/exit.")
		return 0, message
	}
	checkValue, amount := isAmountValid(clientValue, userbalance)
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
		message := getDenominations(amount) + "TRANSACTION SUCCESSFUL."
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
	showDenominations := fmt.Sprintf("500*%d +200*%d +100*%d ", q1, q2, q3)
	return showDenominations
}
