package main

import (
	"fmt"
	"math/rand"
)

func main() {

	//Generate a psuedo-random number between 0 - 100.
	randomNumber := rand.Intn(100)

	//if statement to see if number is < or > 50.
	if randomNumber > 50 {
		fmt.Printf("my random number is %d and is greater than 50\n", randomNumber)
	} else {
		fmt.Printf("my random number is %d and is less than 50\n", randomNumber)
	}
	//Assign Holberton School string to a variable.
	school := "Holberton School"

	if school == "Holberton School" {
		fmt.Printf("I am a student of the %s\n", school)
	}

	//Assign boolean value of true to beautifulWeather.
	beautifulWeather := true

	if beautifulWeather == true {
		fmt.Println("It's a beautiful weather!")
	}

	//Assign a variable with the three founders' names.
	holbertonFounders := []string{"Rudy Rigot", "Sylvain Kalache", "Julien Barbier"}

	for _, each := range holbertonFounders {
		fmt.Printf("%s is a founder\n", each)
	}
}
