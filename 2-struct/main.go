package main

import (
	"fmt"
	"time"
)

type user struct {
	Name string `json:"name"`
	DOB  string `json:"date_of_birth"`
	City string `json:"city"`
}

func age(u user) {
	value := "March 7, 1917"           //actual DOB in string.
	form := "January 2, 2006"          //form setting standard date to identify DOB.
	year, _ := time.Parse(form, value) //making DOB readable as a date and not string.
	birthyear := year.Year()           //extracting birth year from DOB.
	today := time.Now()                //setting today's date.
	Yeartoday := today.Year()          //extracting current year from today's date.
	age := Yeartoday - birthyear       //subtracting years to find out age.

	fmt.Printf("%s who was born in %s would be %d years old today.\n", u.Name, u.City, age)
}
func main() {
	u := user{Name: "Betty Holberton", DOB: "March 7, 1917", City: "Philadelphia"}
	fmt.Printf("Hello %s\n", u.Name)

	age(u)
}
