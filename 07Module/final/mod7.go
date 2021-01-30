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
	database                 = "ATM_DB"
	allWithdrawalsCollection = "WITHDRAWAL_COLLECTION"
	allUserBalanceCollection = "BALANCE_COLLECTION"
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
	idValidityStatus := id.Valid()
	if idValidityStatus == false {
		err := fmt.Errorf("Invalid Object ID")
		return err
	}
	return nil
}
func main() {
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
	balanceTable := session.DB(database).C(allUserBalanceCollection)
	allWithdrawalsTable := session.DB(database).C(allWithdrawalsCollection)

	server := gin.Default()

	server.POST("/:id", func(c *gin.Context) {
		if c.Param("id") == "exit" {
			os.Exit(1)
		}
		isClientIDvalid := bson.IsObjectIdHex(c.Param("id"))
		if isClientIDvalid == false {
			c.JSON(400, gin.H{"Result": "You entered Invalid Object ID"})
		} else {
			balanceID := bson.ObjectIdHex(c.Param("id"))
			var balanceDetails userBalance
			errInFindBalance := balanceTable.FindId(balanceID).Select(bson.M{"balance": 1}).One(&balanceDetails)
			if errInFindBalance != nil {
				c.JSON(400, gin.H{"Result": "Error in finding balance field:" + errInFindBalance.Error() + "Make sure you entered the correct ID that exists."})
			} else {
				presentTime := time.Now()
				timeDifference := presentTime.Add(time.Hour * (-24))
				numOfWithdrawals, errInCountingWithdrawal := allWithdrawalsTable.Find(bson.M{"customer_id": balanceID, "Timestamp": bson.M{"$gt": timeDifference}}).Count()
				if errInCountingWithdrawal != nil {
					c.JSON(500, gin.H{"Result": "Internal error in transaction:" + errInCountingWithdrawal.Error()})
				} else {
					if numOfWithdrawals >= 5 {
						c.JSON(400, gin.H{"Result": "You have made Maximum number of withdrawals allowed in 24 hours:5. You can exit using POST/exit."})
					} else {
						clientInput := withdrawValue{}
						errInClientInput := c.BindJSON(&clientInput)
						if errInClientInput == nil {
							transactionID := bson.NewObjectId()
							errinTransactionID := isObjectIDValid(transactionID)
							transactionStatus, showMessage := generateStatusAndMessage(clientInput.Value, balanceDetails.Balance)
							if errinTransactionID == nil {
								if transactionStatus == "Transaction Allowed" {
									if errInWithdrawal := allWithdrawalsTable.Insert(&withdrawalsTable{ID: transactionID, CustomerID: balanceID, Amount: clientInput.Value, TimeStamp: time.Now()}); errInWithdrawal != nil {
										fmt.Println("Error in inserting amount:", errInWithdrawal)
										c.JSON(500, gin.H{"Result": "Internal Error in transaction:" + errInWithdrawal.Error()})
									} else {
										//Updating the balance table
										selector := bson.M{"_id": balanceID}
										updator := bson.M{"$inc": bson.M{"balance": -clientInput.Value}}
										if errInBalanceUpdate := balanceTable.Update(selector, updator); errInBalanceUpdate != nil {
											fmt.Println("Error in updating balance:", errInBalanceUpdate)
											allWithdrawalsTable.Remove(bson.M{"_id": transactionID})
											c.JSON(500, gin.H{"Result": "Internal Error in transaction:" + errInBalanceUpdate.Error() + "Try Again."})
										} else {
											c.JSON(200, gin.H{"Result": showMessage})
										}
									}
								} else {
									c.JSON(400, gin.H{"Result": showMessage})
								}
							} else {
								c.JSON(500, gin.H{"Result": "Internal Error in transaction.Error occured in transactionID:" + errinTransactionID.Error() + "Try Again."})
							}
						} else {
							c.JSON(400, gin.H{"Result": "You entered Invalid Amount." + errInClientInput.Error()})
						}
					}
				}
			}
		}
	})

	server.GET("/:id", func(c *gin.Context) {
		isClientIDvalid := bson.IsObjectIdHex(c.Param("id"))
		if isClientIDvalid == false {
			c.JSON(400, gin.H{"Result": "Invalid Object ID"})
		} else {
			balanceID := bson.ObjectIdHex(c.Param("id"))
			var balanceDetails userBalance
			errInFindBalance := balanceTable.FindId(balanceID).Select(bson.M{"balance": 1}).One(&balanceDetails)
			if errInFindBalance != nil {
				c.JSON(400, gin.H{"Result": "Error in finding balance field:" + errInFindBalance.Error() + ". Make sure you entered the correct ID that exists."})
			} else {
				c.JSON(200, gin.H{"Your balance is": balanceDetails.Balance})
			}
		}
	})
	err = server.Run(":8080") //  listen for incoming connections
	if err != nil {
		fmt.Println(err)
	}
}
func generateStatusAndMessage(clientValue int, userbalance int) (string, string) {
	if userbalance < 100 {
		message := fmt.Sprintf("Balance less than 100.You can't withdraw anymore.You can exit using POST/exit.")
		return "Low Balance.Transaction not Allowed", message
	}
	statusCode := isAmountValid(clientValue, userbalance)
	switch statusCode {
	case 0:
		message := fmt.Sprintf("Error!Please enter ONE natural number, in the correct format,JSON field name is: value.")
		return "Invalid Amount.Transaction not Allowed", message
	case 10:
		message := fmt.Sprintf("Error!Please enter a Positive value ")
		return "Invalid Amount.Transaction not Allowed", message
	case 20:
		message := fmt.Sprintf("Error!Please enter amount divisible by 100 ")
		return "Invalid Amount.Transaction not Allowed", message
	case 30:
		message := fmt.Sprintf("Error! Please enter amount less than or equal to 5000 ")
		return "Invalid Amount.Transaction not Allowed", message
	case 40:
		message := fmt.Sprintf("Error!Please enter amount less than or equal to your account balance")
		return "Invalid Amount.Transaction not Allowed", message
	default:
		message := getDenominations(clientValue) + "TRANSACTION SUCCESSFUL."
		return "Transaction Allowed", message
	}
}
func isAmountValid(clientAmount int, balance int) int {
	if clientAmount == 0 {
		return 0
	}
	if clientAmount < 0 {
		return 10
	}
	if clientAmount%100 != 0 {
		return 20
	}
	if clientAmount > 5000 {
		return 30
	}
	if balance-clientAmount < 0 {
		return 40
	}
	return 50
}
func getDenominations(totalAmount int) string {
	var q1 int
	var q2 int
	var q3 int
	q1 = totalAmount / 500
	totalAmount = totalAmount % 500
	q2 = totalAmount / 200
	totalAmount = totalAmount % 200
	q3 = totalAmount / 100
	fmt.Printf("500*%d +200*%d +100*%d \n", q1, q2, q3)
	showDenominations := fmt.Sprintf("500*%d +200*%d +100*%d ", q1, q2, q3)
	return showDenominations
}
