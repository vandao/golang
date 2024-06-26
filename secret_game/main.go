package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	persons := [3]string{"Max", "Alex", "Tom"}

	fmt.Println("Number of secret person 0: Max, 1: Alex, 2: Tom.")

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	secretPersonNumber := r.Intn(3)

	scanner := bufio.NewScanner(os.Stdin)
	var answer int
	var err error
	for {

		fmt.Printf("Please enter your answer with numer 0 - 2: ")
		scanner.Scan()
		input := scanner.Text()
		answer, err = strconv.Atoi(input)

		if err != nil {
			log.Panic(err)
		}

		if answer >= 0 && answer <= 2 {
			break
		}
	}

	if answer == secretPersonNumber {
		fmt.Printf("Congratulations!! You answer is correct, %s is the secret person.\n", persons[answer])
	} else {
		fmt.Printf("Sorry!! You answer is incorrect, %d: %s is the secret person.\n", secretPersonNumber, persons[secretPersonNumber])
	}
}
