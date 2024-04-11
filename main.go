package main

import (
	"errors"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"strings"
)

func factorial(n int) *big.Int {
	value := big.NewInt(1)
	for i := n; i > 0; i-- {
		value.Mul(value, big.NewInt(int64(i)))
	}
	return value
}

func fibRecursive(n int, counter map[int]int) (int, map[int]int) {
	counter[n]++
	if n == 0 || n == 1 {
		return 1, counter
	}
	result1, counter1 := fibRecursive(n-1, counter)
	result2, counter2 := fibRecursive(n-2, counter1)
	return result1 + result2, counter2
}

func strToAscii(value string) []byte {
	return []byte(value)
}

func isAliasValueInNFactorial(alias string, n int) bool {
	ascii := strToAscii(alias)
	factorialValue := factorial(n)
	for _, char := range ascii {
		if !strings.Contains(factorialValue.String(), strconv.FormatInt(int64(char), 10)) {
			return false
		}
	}
	return true
}

func findStrongNumber(fullName string) (int, error) {
	fullName = strings.ToLower(fullName)
	names := strings.Split(fullName, " ")
	if len(names) < 2 {
		return 0, errors.New("invalid full name")
	}
	firstName := names[0]
	lastName := names[1]
	alias := firstName[:3] + lastName[:3]
	for n := 1; !isAliasValueInNFactorial(alias, n-1); n++ {
		if isAliasValueInNFactorial(alias, n) {
			return n, nil
		}
	}
	return 0, errors.New("strong number not found")
}

func main() {
	fullName := "Michał Pomirski"
	strongNumber, err := findStrongNumber(fullName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Twoja Silna liczba to: %d\n", strongNumber)

	a, b := fibRecursive(30, make(map[int]int))
	fmt.Println("Kolejne ilości wywołań funkcji fib dla fib(30):")
	fmt.Printf("fib(%d): Wartość: %d\nIlości: %d\n", 30, a, b)
	fmt.Println("Twoja słaba liczba to argument przy wartości najbliższej do liczby wywołań danego argumentu funkcji fib(30).")
	fmt.Println("Np. dla Silnej liczby 247, najbliższa wartość to 233, więc słaba liczba to 18 [233 razy]")
}