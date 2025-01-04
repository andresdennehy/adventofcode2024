package main

import (
	"adventofcode2024/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func calculateBananas(initial int, prices *map[int]int) int {
	var number = initial
	var priceChanges []int
	var processed = make(map[int]bool) // Memo to skip already processed combinations
	var price = number % 10
	for i := 0; i < 2000; i++ {
		number = ((number * 64) ^ number) % 16777216
		number = ((number / 32) ^ number) % 16777216
		number = ((number * 2048) ^ number) % 16777216
		var priceChange = number%10 - price
		if len(priceChanges) == 4 {
			priceChanges = priceChanges[1:]
		}
		priceChanges = append(priceChanges, priceChange)
		if len(priceChanges) >= 4 {
			var key = priceChanges[0]*1000000 + priceChanges[1]*10000 + priceChanges[2]*100 + priceChanges[3]
			if !processed[key] {
				(*prices)[key] = (*prices)[key] + number%10
				processed[key] = true
			}
		}
		price = number % 10
	}
	return 0
}

func main() {

	file, err := os.Open("input.txt")
	utils.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var prices = make(map[int]int) // Map of sequences (encoded in a key) -> # bananas

	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		utils.Check(err)
		calculateBananas(number, &prices)
	}

	var maxBananas int
	for _, bananas := range prices {
		if bananas > maxBananas {
			maxBananas = bananas
		}
	}

	fmt.Printf("Result is %d\n", maxBananas)
}
