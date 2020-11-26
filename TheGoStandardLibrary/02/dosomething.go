package main

import (
	"fmt"
	"os"
	"strconv"
)
func main() {
	args := os.Args[1:]
	{
		if len(args) == 1 && args[0] == "/help"{
			fmt.Println("Usage: dinnertotal <Total Meal Amount> <Tip Percentage>")
			fmt.Println("Example: dinnertotal 20 10")
		}else {
			if len(args) != 2 {
				fmt.Println("You must enter arguments!")
			} else {
				mealTotal, _ := strconv.ParseFloat(args[0], 32)
				tipAmount, _ := strconv.ParseFloat(args[1], 32)
				fmt.Printf("Your meal total will be %.2f", calculateValue(float32(mealTotal), float32(tipAmount)))
			}
		}
	}
}

func calculateValue(mealTotal float32, tipAmount float32) float32{
	totalPrice := mealTotal + (mealTotal * (tipAmount / 100))
	return totalPrice
}
