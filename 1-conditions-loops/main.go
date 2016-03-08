package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	holbertonFounders := []string{"Rudy Rigot", "Sylvain Kalache", "Julien Barbier"}

	if randomNumber := r1.Intn(100); randomNumber > 50 {
		fmt.Println("my random number is", randomNumber, "and is greater than 50")
	} else {
		fmt.Println("my random number is", randomNumber, "and is 50 or less")
	}

	if school := "Holberton School"; school == "Holberton School" {
		fmt.Println("I am a student of the", school)
	} else {
		fmt.Println("What is school?")
	}

	if beautifulWeather := true; beautifulWeather {
		fmt.Println("It's a beautiful weather!")
	} else {
		fmt.Println("Don't you wish you lived somewhere nicer?")
	}

	for i := 0; i < len(holbertonFounders); i++ {
		fmt.Println(holbertonFounders[i], "is a founder")
	}

}
