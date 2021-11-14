package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// define all the global variables
var (
	lowerCharSet          = "abcdefghijklmnopqrstuvwxyz"
	upperCharSet          = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet string = "!@#$%^&*()-"
	numberCharSet         = "123567890"
	minSpecialChar        = 2
	minUpperChar          = 2
	minNumberChar         = 2
	passwordLength        = 10
)

func main() {

	//check if the password length matches the criteria

	totalCharLenWithoutLowerChar := minUpperChar + minSpecialChar + minNumberChar

	if totalCharLenWithoutLowerChar >= passwordLength {
		fmt.Println("Please provide valid password length")
		os.Exit(1)
	}

	// Get the user input - target folder needs to be organized
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("How many passwords you want to generate? - ")
	scanner.Scan()

	numberOfPasswords, err := strconv.Atoi(scanner.Text())

	if err != nil {
		fmt.Println("Please provide correct value for number of passwords")
		os.Exit(1)
	}

	// it generate random number e
	rand.Seed(time.Now().Unix())

	for i := 0; i < numberOfPasswords; i++ {
		password := generatePassword()
		fmt.Printf("Password %v is %v \n", i+1, password)
	}

}

func generatePassword() string {

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
