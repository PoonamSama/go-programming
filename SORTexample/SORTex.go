package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}
type byage []user

func (a byage) Len() int           { return len(a) }
func (a byage) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byage) Less(i, j int) bool { return a[i].Age < a[j].Age }

type byname []user

func (bn byname) Len() int           { return len(bn) }
func (bn byname) Swap(i, j int)      { bn[i], bn[j] = bn[j], bn[i] }
func (bn byname) Less(i, j int) bool { return bn[i].Last < bn[j].Last }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	// your code goes here
	sort.Sort(byage(users))
	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
		//DONT CLOSE THE FOR RANGE LOOP HERE
		sort.Strings(u.Sayings) //WE DONT NEED TO DO IT AGAIN BCZ IT SORTS THE UNDERLYING ARRAY
		for _, say := range u.Sayings {

			fmt.Println("\t \t", say)
		}
	}

	var f1 string
	sort.Sort(byname(users))
	for _, k := range users {
		fmt.Println(k.First, k.Last, k.Age)
		//[]STRING NOT SORTED HERE BUT THE RESULT IS SORTED
		for _, said := range k.Sayings {

			fmt.Println("\t \t", said)
		}
	}
	fmt.Println(f1)

}
