package main

import (
	"fmt"
	"time"
)

type user struct {
	Name string
	DOB  string
	City string
}

func main() {
	u := user{Name: "Betty Holberton", DOB: "March 7, 1917", City: "Philadelphia"}
	fmt.Printf("Hello %s\n", u.Name)

	value := "March 7, 1917"
	form := "January 2, 2006"
	year, _ := time.Parse(form, value)
	birthyear := year.Year()
	today := time.Now()
	Yeartoday := today.Year()
	age := Yeartoday - birthyear

	fmt.Printf("%s who was born in %s would be %d years old today.\n", u.Name, u.City, age)

}
