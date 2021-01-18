package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	database    = "YOUR_DB"
	collection1 = "WITHDRAWAL_COLLECTION"
	collection2 = "BALANCE_COLLECTION"
)

type withdrawValue struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Amount    int           `bson:"amount"`
	TimeStamp time.Time     `bson:"Timestamp"`
}
type userBalance struct {
	ID      bson.ObjectId `bson:"_id,omitempty"`
	Balance int           `bson:"balance"`
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
	Host := []string{
		"127.0.0.1:27017",
	}

	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs: Host,
	})
	if err != nil {
		fmt.Println("Error in session", err)
	}
	defer session.Close()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the amount you want to withdraw:")
	scanner.Scan()
	input, err := strconv.ParseInt(scanner.Text(), 10, 64)
	xput := int(input)
	if err != nil {
		fmt.Println(err)
	}

	allWithdrawals := session.DB(database).C(collection1)
	balanceTable := session.DB(database).C(collection2)
	// Remove
	balanceTable.RemoveAll(nil)

	id1 := bson.NewObjectId()
	errinID := isObjectIDValid(id1)
	fmt.Println("error in id generation:", errinID)
	if errinID != nil {
		log.Fatal(errinID)
	}
	if err := allWithdrawals.Insert(&withdrawValue{ID: id1, Amount: xput, TimeStamp: time.Now()}); err != nil {
		fmt.Println("Error in inserting amount:", err)
	}
	id2 := bson.NewObjectId()
	errinID = isObjectIDValid(id2)
	fmt.Println("error in id generation:", errinID)
	if errinID != nil {
		log.Fatal(errinID)
	}
	if err := balanceTable.Insert(&userBalance{ID: id2, Balance: 9563}); err != nil {
		fmt.Println("Error in inserting balance:", err)
	}
	selector := bson.M{"_id": id2}
	updator := bson.M{"$inc": bson.M{"balance": -xput}}
	if err := balanceTable.Update(selector, updator); err != nil {
		fmt.Println("Error in updating balance:", err)
	}

	var results1 []withdrawValue
	err2 := allWithdrawals.Find(nil).All(&results1)
	if err2 != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Results  All Withdrawals: ", results1)
	}
	var results2 []userBalance

	err3 := balanceTable.Find(nil).All(&results2)
	if err3 != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Results All Balance: ", results2)
	}
}
