package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

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
		Collection1 = "BALANCE_COLLECTION"
	)
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs: Host,
	})
	if err != nil {
		fmt.Println("Error in session", err)
	}
	defer session.Close()

	c1 := session.DB(Database).C(Collection1)
	c1.RemoveAll(nil)
	id1 := bson.NewObjectId()
	if err := c1.Insert(&userBalance{ID: id1, Balance: 9563}); err != nil {
		fmt.Println("Error in inserting balance:", err)
	}
	var results2 []userBalance

	err3 := c1.Find(nil).All(&results2)
	if err3 != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Results All Balance: ", results2)
	}
}
