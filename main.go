package main

import (
	"fmt"
	"strings"
	"math/big"
)

func factorial(n Int) Int {
    var value Int = 1;
    for i := n; i > 0; i-- {
        value *= i
    }
    return value
}

func strToAscii(value string) []byte {
    return []byte(value)
}

func isAliasValueInNFactorial(alias string, n Int) bool {
    ascii := strToAscii(alias)
	result := true
	factorialValue := factorial(n)
    for _, char := range(ascii){
        if !strings.Contains(string(factorialValue), string(char)) {
			result = false
		}
    }
	return result
}

func main() {
	fullName := "Michał Pomirski"
	fullName = strings.ToLower(fullName)
	firstName := strings.Split(fullName, " ")[0]
	lastName := strings.Split(fullName, " ")[1]
	alias := firstName[:3] + lastName[:3]
	n := int64(30)
	fmt.Printf("%d! = %d\n", n, factorial(n))
	alias = "piar"
	if isAliasValueInNFactorial(alias, n) {
		fmt.Printf("Alias %s (%d) is in %d factorial: %d\n", alias, strToAscii(alias), n, factorial(n))
	} else {
		fmt.Printf("Alias %s (%d) is not in %d factorial: %d\n", alias, strToAscii(alias), n, factorial(n))
	}
	

}
