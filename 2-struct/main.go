package main

import (
	"fmt"
	"time"
)

func getAge(bDay string) (int, error) {
	then, err := time.Parse("January 2, 2006", bDay)
	age := int(time.Since(then).Hours() / 24 / 365.25)
	return age, err
}

func main() {
	type user struct {
		Name string `json:"name"`
		DOB  string `json:"date_of_birth"`
		City string `json:"city"`
	}

	u := user{"Betty Holberton", "March 7, 1917", "Philadelphia"}
	age, err := getAge(u.DOB)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	fmt.Printf("Hello %s\n", u.Name)
	fmt.Printf("%s who was born in %s would be %d years old today\n", u.Name, u.City, age)
}
