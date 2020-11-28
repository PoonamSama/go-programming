package main

import (
	"fmt"
)

func main() {
x:= []int{1,2,3}
y:= []int{4,5,6}
x=append(x,7,8,9,10,11,12,13,14)
x=append(x,y...)

	fmt.Println(x)
	fmt.Println(y)
	x=append(x[:2],x[4:]...)// 2 se 4 deleted
	fmt.Println(x)
	fmt.Println(x[4:9]) //4se pehle,including 4 removed;9 k baad removed 


}
