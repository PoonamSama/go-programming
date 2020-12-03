package main
import "fmt"
func main(){
a:=incrementor()//calling the function
b:=incrementor()
fmt.Println(a())
fmt.Println(a())
fmt.Println(a())
fmt.Println(a())
fmt.Println(b())
fmt.Println(b())
}
func incrementor() func()int{//func incr return a func tht returns an int
	var x int  //scope of x is ltd here. x is assigned 0
return func() int{
	 x++
	 return x
}
}