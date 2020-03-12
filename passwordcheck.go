// passwordcheck
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func checker(a string, passLength int) (upperCase int, lowerCase int, numbers int, specialChar int) {

	for i := 0; i < passLength; i++ {
		letter := a[i : i+1]
		upper := strings.Contains("ABCDEFGHIJKLMNOPQRSTUVWXYZ", letter)
		lower := strings.Contains("abcdefghijklmnopqrstuvwxyz", letter)
		num := strings.Contains("1234567890", letter)
		if upper == true {
			upperCase += 1
		} else if lower == true {
			lowerCase += 1
		} else if num == true {
			numbers += 1
		} else {
			specialChar += 1
		}
		fmt.Printf("*")
	}
	return
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	duration := time.Second
	fmt.Println("Please fill a password in:")
	scanner.Scan()
	password := scanner.Text()
	length := len(password)
	fmt.Println("\nChecking the strength of the password...")
	time.Sleep(duration)
	fmt.Println("The password length is:", length)
	time.Sleep(duration)
	resultUpperCase, resultLowerCase, resultNumbers, resultSpecialChar := checker(password, length)
	fmt.Println("\nAmount of uppercase letters:", resultUpperCase)
	time.Sleep(duration)
	fmt.Println("Amount of numbers:", resultNumbers)
	time.Sleep(duration)
	fmt.Println("Amount of lowercase letters:", resultLowerCase)
	time.Sleep(duration)
	fmt.Println("Amount of special characters:", resultSpecialChar)
	time.Sleep(duration)
	if length < 12 || resultLowerCase < 2 || resultNumbers < 2 || resultUpperCase < 2 || resultSpecialChar < 1 {
		fmt.Println("Your password is weak!")
	} else {
		fmt.Println("Good enough I suppose")
	}

}
