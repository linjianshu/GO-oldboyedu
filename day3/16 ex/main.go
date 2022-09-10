package main

import "fmt"

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func main() {
	left := dispatchCoin()
	fmt.Println(distribution)
	fmt.Println("剩下：", left)
}
func dispatchCoin() int {
	for _, v := range users {
		coin := 0
		for i, _ := range v {
			if string(v[i]) == "e" || string(v[i]) == "E" {
				coin++
			} else if string(v[i]) == "i" || string(v[i]) == "I" {
				coin += 2
			} else if string(v[i]) == "o" || string(v[i]) == "O" {
				coin += 3
			} else if string(v[i]) == "u" || string(v[i]) == "U" {
				coin += 4
			}
		}
		distribution[v] = coin
		coins -= coin
	}
	return coins
}
