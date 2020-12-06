package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)
func main(){

s := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost) //this is to encrypt
	if err != nil {                                                  //returns a sliceofbytes and error
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(bs)

	loginPword1 := `password1234`

	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPword1)) // func to decrypt a given pword
	if err != nil {                                             // returns error
		fmt.Println("YOU CAN'T LOGIN")
		return
	}

	fmt.Println("You're logged in")
}