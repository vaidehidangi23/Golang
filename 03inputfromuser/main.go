package main

import (
	"strconv"
	"os"
	"bufio"
	"strings"
	"fmt"
)

func main(){
	welcome := "welcome"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Rating:")

	//comma ok(as we dont have try and catch in go)
	//if you dont care about anything in th go you just use an underscore
    input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating,",input)
	fmt.Printf("Type of rating:%T", input)

   //conversion in go
   numRating, err := strconv.ParseFloat(strings.TrimSpace(input),64)

   if err!=nil{
	fmt.Println(err)
   }else{
	fmt.Println("Added 1 to your rating",numRating+1)
   }
}