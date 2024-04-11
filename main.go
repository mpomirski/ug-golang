package main

import (
	"errors"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"strings"
)

func factorial(n int) big.Int {
    var value big.Int = *big.NewInt(1);
    for i := n; i > 0; i-- {
        value = *value.Mul(&value, big.NewInt(int64(i)));
    }
    return value;
}

func fibRecursive(n int, counter map[int]int) (int, map[int]int) {
	if n == 0 || n == 1{
		counter[n] += 1;
		return 1, counter;
	}
	counter[n] += 1;
	result1, counter1 := fibRecursive(n-1, counter);
	result2, counter2 := fibRecursive(n-2, counter1);
	return result1 + result2, counter2;
}

func strToAscii(value string) []byte {
    return []byte(value);
}

func isAliasValueInNFactorial(alias string, n int) bool {
    ascii := strToAscii(alias);
	result := true;
	factorialValue := factorial(n);
    for _, char := range(ascii){
        if (!strings.Contains(factorialValue.String(), strconv.FormatInt(int64(char), 10))) {
			result = false;
		}
    }
	return result;
}

func findStrongNumber(fullName string) (int, error) {
	fullName = strings.ToLower(fullName);
	firstName := strings.Split(fullName, " ")[0];
	lastName := strings.Split(fullName, " ")[1];
	alias := firstName[:3] + lastName[:3];
	for n := 1; !isAliasValueInNFactorial(alias, n-1); n++{
		if isAliasValueInNFactorial(alias, n) {
			return n, nil;
		}
	}
	return 0, errors.New("Something went wrong");
}
func main() {
	fullName := "Michał Pomirski";
	strongNumber, err := findStrongNumber(fullName);
	if err != nil{
		log.Fatal(err);
	}
	fmt.Printf("Twoja Silna liczba to: %d\n", strongNumber);
	// Moja Silna liczba: 247

	a, b := fibRecursive(30, make(map[int]int));
	fmt.Println("Kolejne ilości wywołań funkcji fib dla fib(30):");
	fmt.Printf("fib(%d): Wartość: %d\nIlości: %d\n", 30, a, b);
	fmt.Println("Twoja słaba liczba to argument przy wartości najbliższej do liczby wywołań danego argumentu funkcji fib(30).");
	fmt.Println("Np. dla Silnej liczby 247, najbliższa wartość to 233, więc słaba liczba to 18 [233 razy]")
	//Najbliższa wartość do 247, to ilość wywołań fib(18): 233
	// Moja słaba liczba: 30 [233 razy]
}


