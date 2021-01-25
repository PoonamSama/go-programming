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
	database    = "ATM_DB" //ATM_db
	collection1 = "WITHDRAWAL_COLLECTION"
	collection2 = "BALANCE_COLLECTION"
)

type userBalance struct {
	ID      bson.ObjectId `bson:"_id,omitempty"`
	Balance int           `bson:"balance"`
}
type withdrawalsTable struct {
	ID                     bson.ObjectId `bson:"_id,omitempty"`
	Amount                 int           `bson:"amount"`
	TimeStamp              time.Time     `bson:"Timestamp"`
	WithdrawalNumberIn24Hr int           `bson:"WithdrawalNumber"`
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
	var transactionDetails []withdrawalsTable
	var numberOfWithdrawals int
	var timeConstraint string
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
	}
	defer session.Close()
	newbalanceID := bson.NewObjectId()
	errinBalanceID := isObjectIDValid(newbalanceID)
	if errinBalanceID != nil {
		fmt.Println("Error in New Balance ID generation:", errinBalanceID)
	}
	transactionID := bson.NewObjectId()
	errinTransactionID := isObjectIDValid(transactionID)
	if errinTransactionID != nil {
		fmt.Println("Error in transaction ID generation:", errinTransactionID)
	}
	balanceTable := session.DB(database).C(collection2)
	allWithdrawals := session.DB(database).C(collection1)
	//balanceTable.RemoveAll(nil)
	//allWithdrawals.RemoveAll(nil)
	err = balanceTable.Find(nil).All(&balanceDetails)
	if err != nil {
		fmt.Println("Error in loading all balance data:", err)
	}
	if len(balanceDetails) == 0 {
		//Inserting initial value in the balance table
		if errinBalanceID == nil {
			if errInBalance = balanceTable.Insert(&userBalance{ID: newbalanceID, Balance: 9563}); errInBalance != nil {
				fmt.Println("Error in inserting balance:", err)
			}
			err = balanceTable.FindId(newbalanceID).Select(bson.M{"balance": 1}).All(&balanceDetails)
			if err != nil {
				fmt.Println("Error in finding balance field:", err)
			}
		}
	}
	err = allWithdrawals.Find(nil).All(&transactionDetails)
	if err != nil {
		fmt.Println(err)
	}
	oldnumberOfWithdrawals := len(transactionDetails)
	if len(transactionDetails) == 0 {
		numberOfWithdrawals = 0
	}
	if len(transactionDetails) > 0 {
		numberOfWithdrawals = transactionDetails[len(transactionDetails)-1].WithdrawalNumberIn24Hr
		lastTimeofWithdraw := transactionDetails[len(transactionDetails)-1].TimeStamp
		presentTime := time.Now()
		timeLimit := time.Duration(24 * time.Hour)
		timeDifference := presentTime.Sub(lastTimeofWithdraw)
		if timeDifference >= timeLimit {
			timeConstraint = "NEW DAY"
			numberOfWithdrawals = 0
		}
	}
	newbalanceID = balanceDetails[0].ID
	err = balanceTable.FindId(newbalanceID).Select(bson.M{"balance": 1}).All(&balanceDetails)
	if err != nil {
		fmt.Println("Error in finding balance field:", err)
	}
	server := gin.Default()

	server.POST("/", func(c *gin.Context) {
		clientInput := withdrawValue{}
		c.BindJSON(&clientInput)

		clientAmount, showMessage, numOfWithdrawals := withdrawAmount(clientInput.Value, balanceDetails[0].Balance, numberOfWithdrawals)
		if errinBalanceID == nil && errinTransactionID == nil {
			if clientAmount != 0 {
				if errInWithdrawal = allWithdrawals.Insert(&withdrawalsTable{ID: transactionID, Amount: clientAmount, TimeStamp: time.Now(), WithdrawalNumberIn24Hr: numOfWithdrawals}); errInWithdrawal != nil {
					fmt.Println("Error in inserting amount:", errInWithdrawal)
				}
			}
			//Updating the balance table
			selector := bson.M{"_id": newbalanceID}
			updator := bson.M{"$inc": bson.M{"balance": -clientAmount}}
			if errInWithdrawal == nil {
				if errInBalance = balanceTable.Update(selector, updator); errInBalance != nil {
					fmt.Println("Error in updating balance:", errInBalance)
					allWithdrawals.Remove(bson.M{"_id": transactionID})
				}
			}
			err = balanceTable.FindId(newbalanceID).Select(bson.M{"balance": 1}).All(&balanceDetails)
			if err != nil {
				fmt.Println(err)
			}
			err = allWithdrawals.Find(nil).All(&transactionDetails)
			if err != nil {
				fmt.Println(err)
			}
			numberOfWithdrawals = transactionDetails[len(transactionDetails)-1].WithdrawalNumberIn24Hr
		}
		if timeConstraint == "NEW DAY" {
			numberOfWithdrawals = len(transactionDetails) - oldnumberOfWithdrawals
		}

		c.JSON(200, gin.H{"Result": showMessage})
		transactionID = bson.NewObjectId()
		errinTransactionID = isObjectIDValid(transactionID)
		if errinTransactionID != nil {
			fmt.Println(errinTransactionID)
		}
	})

	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"Your balance is": balanceDetails[0].Balance})
	})

	server.POST("/exit", func(c *gin.Context) {
		os.Exit(1)
	})

	err = server.Run(":8080") //  listen for incoming connections
	if err != nil {
		fmt.Println(err)
	}
}
func withdrawAmount(clientValue int, userbalance int, numberOfWithdrawals int) (int, string, int) {
	if numberOfWithdrawals >= 5 {
		fmt.Println("Error! you have made maximum number of transactions for today: 5.")
		message := fmt.Sprintf("Maximum transactions reached for a day:5. You can exit using POST/exit.")
		return 0, message, 0
	}
	if userbalance < 100 {
		message := fmt.Sprintf("Balance less than 100.You can't withdraw anymore.You can exit using POST/exit.")
		return 0, message, 0
	}
	checkValue, amountToWithdraw := isAmountValid(clientValue, userbalance)
	switch checkValue {
	case 0:
		message := ("Error!Please enter ONE natural number.")
		return 0, message, 0
	case 10:
		message := fmt.Sprintf("Error!Please enter a Positive value ")
		return 0, message, 0
	case 20:
		message := fmt.Sprintf("Error!Please enter amount divisible by 100 ")
		return 0, message, 0
	case 30:
		message := fmt.Sprintf("Error! Please enter amount less than or equal to 5000 ")
		return 0, message, 0
	case 40:
		message := fmt.Sprintf("Error!Please enter amount less than or equal to your account balance")
		return 0, message, 0
	default:
		message := getDenominations(amountToWithdraw) + "TRANSACTION SUCCESSFUL."
		numberOfWithdrawals++
		return amountToWithdraw, message, numberOfWithdrawals
	}
}
func isAmountValid(clientInput int, balance int) (int, int) {
	if clientInput == 0 {
		return 0, clientInput
	}
	if clientInput < 0 {
		return 10, clientInput
	}
	if clientInput%100 != 0 {
		return 20, clientInput
	}
	if clientInput > 5000 {
		return 30, clientInput
	}
	if balance-clientInput < 0 {
		return 40, clientInput
	}
	return 50, clientInput
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
	printDenominations := fmt.Sprintf("500*%d +200*%d +100*%d ", q1, q2, q3)
	return printDenominations
}
