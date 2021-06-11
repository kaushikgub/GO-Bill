package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// Making A New Bill
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

// Formating Bills
func (b bill) format() string {
	fs := "Bill Breakdown \n"
	total := 0.0

	// Listing The Items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%.2f \n", k + ":", v)
		total += v
	}

	// Tip
	fs += fmt.Sprintf("%-25v ...$%.2f \n", "Tip:", b.tip)

	// Total
	fs += fmt.Sprintf("%-25v ...$%.2f \n", "Total:", total + b.tip)

	return fs
}

// Updting The Tip
func (b *bill) updateTip(tip float64){
		b.tip = tip
}

// Add Item To Bill
func (b *bill) addItem(name string, price float64){
	b.items[name] = price
}

// Save Bill
func (b *bill) save(){
	data := []byte(b.format())

	err := os.WriteFile("bills/" + b.name + ".txt", data, 0644)

	if err != nil {
		panic(err)
	}

	fmt.Println("Bill is saved to file")
}
