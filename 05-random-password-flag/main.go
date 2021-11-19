package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// define all the global variables
const (
	lowerCharSet          = "abcdefghijklmnopqrstuvwxyz"
	upperCharSet          = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet string = "!@#$%^&*()-"
	numberCharSet         = "123567890"
)

func main() {

	var (
		passwordLength    int
		numberOfPasswords int
		minSpecialChar    int
		minUpperChar      int
		minNumberChar     int
	)

	// Set the flags. The first value is the flag, the second is the default value, the third is the help message. Try go run main.go -h
	pwl := flag.Int("length", 10, "The length of the desired password")
	npw := flag.Int("count", 1, "The number of desired passwords")
	s := flag.Int("min-special", 2, "The minimum number of special characters")
	u := flag.Int("min-upper", 2, "The minimum number of uppercase characters")
	n := flag.Int("min-number", 2, "The minimum number of numbers")

	// Tell the application to read the values
	flag.Parse()

	// Set the variables to be equal to what was provided via flags. Note: flags are pointers
	passwordLength = *pwl
	numberOfPasswords = *npw
	minSpecialChar = *s
	minUpperChar = *u
	minNumberChar = *n

	//check if the password length matches the criteria

	totalCharLenWithoutLowerChar := minUpperChar + minSpecialChar + minNumberChar

	if totalCharLenWithoutLowerChar >= passwordLength {
		fmt.Println("Please provide valid password length")
		os.Exit(1)
	}

	// it generate random number e
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < numberOfPasswords; i++ {
		password := generatePassword(passwordLength, minSpecialChar, minUpperChar, minNumberChar)
		fmt.Printf("Password %v is %v \n", i+1, password)
	}

}

func generatePassword(passwordLength int, minSpecialChar int, minUpperChar int, minNumberChar int) string {

	// declare empty password variable
	password := ""

	// generate random special character based on minSpecialChar

	for i := 0; i < minSpecialChar; i++ {
		random := rand.Intn(len(specialCharSet))
		//fmt.Println(specialCharSet[random])
		//fmt.Printf("%v and %T \n", random, specialCharSet[random])
		password = password + string(specialCharSet[random])
	}

	// generate random upper character based on minUpperChar
	for i := 0; i < minUpperChar; i++ {
		random := rand.Intn(len(upperCharSet))
		password = password + string(upperCharSet[random])
	}

	// generate random upper character based on minNumberChar
	for i := 0; i < minNumberChar; i++ {
		random := rand.Intn(len(numberCharSet))
		password = password + string(numberCharSet[random])
	}

	// find remaining lowerChar
	totalCharLenWithoutLowerChar := minUpperChar + minSpecialChar + minNumberChar

	remainingCharLen := passwordLength - totalCharLenWithoutLowerChar

	// generate random lower character based on remainingCharLen
	for i := 0; i < remainingCharLen; i++ {
		random := rand.Intn(len(lowerCharSet))
		password = password + string(lowerCharSet[random])
	}

	// shuffle the password string

	passwordRune := []rune(password)
	rand.Shuffle(len(passwordRune), func(i, j int) {
		passwordRune[i], passwordRune[j] = passwordRune[j], passwordRune[i]
	})

	password = string(passwordRune)
	return password
}
