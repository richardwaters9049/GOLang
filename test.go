package main

import "fmt"

func main() {
	fmt.Println("Welcome to the quiz!!")
	fmt.Printf("Please enter your name: ")

	var name string
	fmt.Scan(&name)

	fmt.Printf("Hello %v, welcome to the game!!\n", name)

	fmt.Printf("Enter your age: ")
	var age uint
	fmt.Scan(&age)

	fmt.Println(age >= 10)

	if age >= 10 {
			fmt.Println("Yay you can play!")
	}else{
		fmt.Println("Go away!")
		return
	}

	// Allow GO to decide data type
	score := 0
	num_questions := 2

	fmt.Printf("Who are better LFC 2009 or MUFC 2010? ")
	var answer string
	var answer2 string
	fmt.Scan(&answer, &answer2)

	if answer + " " + answer2 == "LFC 2009" || answer + " " + answer2 == "lfc 2009" {
		fmt.Println("Correct!")
		score ++ //Increase score by 1
		fmt.Println("Your current score is ", score)
	}else{
		fmt.Println("Incorrect!")
		fmt.Println("Your current score is ", score)
	}

	fmt.Printf("How many feet do I have? ")
	var feet uint
	fmt.Scan(&feet)

	if feet == 2 {
		fmt.Println("Correct!")
		score ++ //Increase score by 1
		fmt.Println("Your current score is ", score)
	}else{
		fmt.Println("Incorrect!")
		fmt.Println("Your current score is ", score)
	}

	fmt.Printf("You scored %v out of %v.\n", score, num_questions)
	percent := (float64(score) / float64(num_questions)) * 100
	fmt.Printf("You scored: %v%% ", percent)
}