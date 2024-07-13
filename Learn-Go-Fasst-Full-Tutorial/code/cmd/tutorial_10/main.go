package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type contactInfo struct {
	Name  string
	Email string
}

type purchaseInfo struct {
	Name   string
	Price  float32
	Amount int
}

func main() {
	var contacts []contactInfo = loadJSON[contactInfo]("./contacts.json")
	fmt.Printf("%+v\n", contacts)

	var purchases []purchaseInfo = loadJSON[purchaseInfo]("./purchases.json")
	fmt.Printf("%+v\n", purchases)
}

func loadJSON[T contactInfo | purchaseInfo](filePath string) []T {
	data, _ := os.ReadFile(filePath)

	var loaded []T
	err := json.Unmarshal(data, &loaded)
	if err != nil {
		return nil
	}

	return loaded
}
