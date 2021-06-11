package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(message string) (string, error) {
	fmt.Print(message)

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill {
	name, _ := getInput("Bill Name : ")

	b:= newBill(name)

	fmt.Println("CreatedThe Bill ~ ", b.name)

	return b
}

func userOptions(b bill){
	opt, _ := getInput("Chose a option: a - Add Item, t - Add Tip, s - Save Bill : ")
	switch opt{
	case "a":
		name, _ := getInput("Enter the item name : ")
		itemPrice, _ := getInput("Enter the item price : ")

		price, err := strconv.ParseFloat(itemPrice, 64)
		if err !=nil {
			fmt.Println("The price must be a number !")
			userOptions(b)
		}

		b.addItem(name, price)
		fmt.Println("Item Added : ", name, price)
		userOptions(b)

	case "t":
		billTip, _ := getInput("Enter tip amount : ")
		tip, err := strconv.ParseFloat(billTip, 64)
		if err !=nil {
			fmt.Println("The tip must be a number !")
			userOptions(b)
		}

		b.updateTip(tip)
		fmt.Println("Tips Added", tip)
		userOptions(b)
	case "s":
		b.save()
		fmt.Println("Saved :)")
	default:
		fmt.Println("That wad not a valid option")
		userOptions(b)
	}
}


func main() {
	myBill := createBill()
	userOptions(myBill)
}
