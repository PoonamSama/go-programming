package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

func main() {
	Host := []string{
		"127.0.0.1:27017",
	}
	const (
		Username    = "YOUR_USERNAME"
		Password    = "YOUR_PASS"
		Database    = "YOUR_DB"
		Collection1 = "WITHDRAWAL_COLLECTION"
		Collection2 = "BALANCE_COLLECTION"
	)
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

	c1 := session.DB(Database).C(Collection1)
	c2 := session.DB(Database).C(Collection2)
	// Remove
	c2.RemoveAll(nil)
	id1 := bson.NewObjectId()
	if err := c1.Insert(&withdrawValue{ID: id1, Amount: xput, TimeStamp: time.Now()}); err != nil {
		fmt.Println("Error in inserting amount:", err)
	}
	id2 := bson.NewObjectId()
	if err := c2.Insert(&userBalance{ID: id2, Balance: 9563}); err != nil {
		fmt.Println("Error in inserting balance:", err)
	}
	selector := bson.M{"_id": id2}
	updator := bson.M{"$inc": bson.M{"balance": -xput}}
	if err := c2.Update(selector, updator); err != nil {
		fmt.Println("Error in updating balance:", err)
	}

	var results1 []withdrawValue
	err2 := c1.Find(nil).All(&results1)
	if err2 != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Results  All Withdrawals: ", results1)
	}
	var results2 []userBalance

	err3 := c2.Find(nil).All(&results2)
	if err3 != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Results All Balance: ", results2)
	}
}
