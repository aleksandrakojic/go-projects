package main

import "fmt"

func main() {
    adidasFactory, _ := GetSportsFactory("adidas")
	nikeFactory, _ := GetSportsFactory("nike")

    adidasShoe := adidasFactory.makeShoe()
	adidasShirt := adidasFactory.makeShirt()
	nikeShirt := nikeFactory.makeShirt()
	nikeShoe := nikeFactory.makeShoe()

	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)

	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)

}

func printShoeDetails(shoe IShoe) {
	fmt.Printf("Logo: %s", shoe.getLogo())
    fmt.Println()
    fmt.Printf("Size: %d", shoe.getSize())
    fmt.Println()
}

func printShirtDetails(shirt IShirt) {
    fmt.Printf("Logo: %s", shirt.getLogo())
    fmt.Println()
    fmt.Printf("Size: %d", shirt.getSize())
    fmt.Println()
}
