package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golangs")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Guava"

	fmt.Println("Fruit list : ", fruitList)
	fmt.Println("Fruit list : ", len(fruitList))

	var vegList = [5]string{"potato", "ladyfinger", "cabbage"}
	fmt.Println("Vegy list is: ", len(vegList))

}