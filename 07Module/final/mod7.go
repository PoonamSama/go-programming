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
	database    = "ATM_DB"
	collection1 = "WITHDRAWAL_COLLECTION"
	collection2 = "BALANCE_COLLECTION"
)

type userBalance struct {
	ID      bson.ObjectId `bson:"_id,omitempty"`
	Balance int           `bson:"balance"`
}
type withdrawalsTable struct {
	ID         bson.ObjectId `bson:"_id,omitempty"`
	CustomerID bson.ObjectId `bson:"customer_id,omitempty"`
	Amount     int           `bson:"amount"`
	TimeStamp  time.Time     `bson:"Timestamp"`
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
	var errInBalanceUpdate error
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
	balanceTable := session.DB(database).C(collection2)
	allWithdrawals := session.DB(database).C(collection1)

	server := gin.Default()

	server.POST("/:id", func(c *gin.Context) {
		if c.Param("id") == "exit" {
			os.Exit(1)
		}
		checkIDvalid := bson.IsObjectIdHex(c.Param("id"))
		if checkIDvalid == false {
			c.JSON(400, gin.H{"Result": "Invalid Object ID"})
		} else {
			balanceID := bson.ObjectIdHex(c.Param("id"))
			errInFindBalance := balanceTable.FindId(balanceID).Select(bson.M{"balance": 1}).All(&balanceDetails)
			if errInFindBalance != nil {
				fmt.Println("Error in finding balance field:", errInFindBalance)
				c.JSON(500, gin.H{"Result": "Internal Error in finding balance field."})
			} else if len(balanceDetails) == 0 {
				c.JSON(400, gin.H{"Result": "This ID does not exist"})
			} else {
				presentTime := time.Now()
				timeDifference := presentTime.Add(time.Hour * (-24))
				numOfWithdrawals, errInCountingWithdrawal := allWithdrawals.Find(bson.M{"customer_id": balanceID, "Timestamp": bson.M{"$gt": timeDifference}}).Count()
				clientInput := withdrawValue{}
				c.BindJSON(&clientInput)

				transactionID := bson.NewObjectId()
				errinTransactionID := isObjectIDValid(transactionID)
				if errinTransactionID != nil {
					fmt.Println(errinTransactionID)
				}
				clientAmount, showMessage := withdrawAmount(clientInput.Value, balanceDetails[0].Balance, numOfWithdrawals)
				if errinTransactionID == nil && errInCountingWithdrawal == nil && clientAmount != 0 {
					if errInWithdrawal = allWithdrawals.Insert(&withdrawalsTable{ID: transactionID, CustomerID: balanceID, Amount: clientAmount, TimeStamp: time.Now()}); errInWithdrawal != nil {
						fmt.Println("Error in inserting amount:", errInWithdrawal)
					} // err in inserting amount

					//Updating the balance table
					selector := bson.M{"_id": balanceID}
					updator := bson.M{"$inc": bson.M{"balance": -clientAmount}}
					if errInWithdrawal == nil {
						if errInBalanceUpdate = balanceTable.Update(selector, updator); errInBalanceUpdate != nil {
							fmt.Println("Error in updating balance:", errInBalanceUpdate)
							allWithdrawals.Remove(bson.M{"_id": transactionID})
						}
					}
				}
				if errinTransactionID == nil && errInCountingWithdrawal == nil && errInWithdrawal == nil && errInBalanceUpdate == nil {
					c.JSON(200, gin.H{"Result": showMessage})
				} else {
					c.JSON(500, gin.H{"Result": "Internal Error in transaction.Try Again."})
				}
			}
		}
	})

	server.GET("/:id", func(c *gin.Context) {
		checkIDvalid := bson.IsObjectIdHex(c.Param("id"))
		if checkIDvalid == false {
			c.JSON(400, gin.H{"Result": "Invalid Object ID"})
		} else {
			balanceID := bson.ObjectIdHex(c.Param("id"))
			errInFindBalance := balanceTable.FindId(balanceID).Select(bson.M{"balance": 1}).All(&balanceDetails)
			if errInFindBalance != nil {
				fmt.Println("Error in finding balance field:", errInFindBalance)
				c.JSON(500, gin.H{"Result": "Internal error in finding balance field. Please try again later."})
			} else if len(balanceDetails) == 0 {
				c.JSON(400, gin.H{"Result": "This ID does not exist"})
			} else {
				c.JSON(200, gin.H{"Your balance is": balanceDetails[0].Balance})
			}
		}
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
