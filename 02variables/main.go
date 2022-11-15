package main

import(
	"fmt"
)
//const Logintoken string ="bdhfbsadfj"//this is a public variable(by putting the first letter capital you can make it a public variable )
//abc := 300000(this will give an error because inside the method you are allowed to use this walrus operator)
func main(){
	var username string = "vaidehi"
	fmt.Println(username)
	fmt.Printf("variable is of type:%T \n", username)
	
	var isRegistered bool = true
	fmt.Println(isRegistered)
	fmt.Printf("variable is of type:%T \n", isRegistered)

	var float float64 = 233.45678934578
	fmt.Println(float)
	fmt.Printf("variable is of type:%T \n", float)

   //default values
	var variable int
	fmt.Println(variable)
	fmt.Printf("variable is of type:%T \n", variable)

	//implicit type
	var a="vaidehi"
	fmt.Println(a)
	fmt.Printf("variable is of type:%T \n", a)

	//no var style
	users := 20000
	fmt.Println(users)

}
