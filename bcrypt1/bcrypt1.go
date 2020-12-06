package main
import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)
func main(){
	password:="123go"
	bslice, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(password)  //the password we gave
	fmt.Println(bslice)  //the encrypted password
}