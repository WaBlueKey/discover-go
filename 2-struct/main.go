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
	form := "January 2, 2006"             //form setting standard date to identify DOB.
	DOB, err := time.Parse(form, u.DOB)   //making DOB readable as a date and not string.
	age := time.Now().Year() - DOB.Year() //subtracting years to find out age.

	if err != nil {
		fmt.Println("Please try again. Error:", err)
	} else {
		fmt.Printf("%s who was born in %s would be %d years old today.\n", u.Name, u.City, age)
	}
}
func main() {
	u := user{Name: "Betty Holberton", DOB: "March 7, 1917", City: "Philadelphia"}
	fmt.Printf("Hello %s\n", u.Name)
	age(u)
}
