package main
import(
	"fmt"
	"encoding/json"
)
type person struct {  //DONT USE TYPE NAME []STRUCT
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}
func main() {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println(s) //got it from json to go
	fmt.Printf("%T\n",s)
	var people []person
	err:=json.Unmarshal([]byte(s),&people) //convert s:type string into type[]byte and &people to store unmarshal there
	if err!=nil{
		fmt.Println(err)
	}
	for i,person:=range people{
		fmt.Println("person number:",i)
		fmt.Println(person.First,person.Last,person.Age)
		for _,saidthis:= range person.Sayings{
			fmt.Println("\t \t",saidthis)
		}
	}
}