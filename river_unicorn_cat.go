package main

import "fmt"

// The Mentoring for Success Project
// This project will provide educational resources and mentorship opportunities for those seeking to gain the skills and knowledge necessary for success in the world of technology.

// ------------
// MAIN PROGRAM
// ------------

func main() {
	// Welcome message
	fmt.Println("Welcome to The Mentoring for Success Project!")
	
	// Create a list of available courses
	var courses = []string{"Intro to Coding","JavaScript Fundamentals","Python for Data Science","Intro to Machine Learning","Web Development with HTML/CSS"}
	
	// Prompt the user to select a course
	fmt.Println("Please enter the number of the course you would like to learn:")
	
	// Print the list of courses
	for i, course := range courses {
		fmt.Printf("%v) %v \n", i + 1, course)
	}
	
	// Prompt the user to select a course
	var courseChoice int
	fmt.Print("-> ")
	fmt.Scan(&courseChoice)
	
	// Print the course selection
	fmt.Printf("\nYou have selected course %v: %v\n", courseChoice, courses[courseChoice - 1])
	
	// Print a list of mentor options
	var mentors = []string{"Matt", "John", "Geoff", "Chloe", "David"}
	
	fmt.Println("\nPlease select your mentor from the following list:")
	
	// Print the list of mentors
	for i, mentor := range mentors {
		fmt.Printf("%v) %v \n", i + 1, mentor)
	}
	
	// Prompt the user to select a mentor
	var mentorChoice int
	fmt.Print("-> ")
	fmt.Scan(&mentorChoice)
	
	// Print the mentor selection
	fmt.Printf("\nYou have selected %v as your mentor.\n",  mentors[mentorChoice - 1])
	
	// Congratulations message
	fmt.Println("\nCongratulations! You have successfully enrolled in The Mentoring for Success Project!")
}