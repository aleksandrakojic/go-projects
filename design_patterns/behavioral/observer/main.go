/* Observer is a behavioral design pattern that allows some objects to notify other objects about changes in their state. */

package main

func main() {

    shirtItem := newItem("Nike Shirt")

    observerFirst := &Customer{id: "abc@gmail.com"}
    observerSecond := &Customer{id: "xyz@gmail.com"}

    shirtItem.register(observerFirst)
    shirtItem.register(observerSecond)

    shirtItem.updateAvailability()
}