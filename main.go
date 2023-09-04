package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type Question struct {
	Prompt         string   `json:"prompt"`
	Options        []string `json:"options"`
	CorrectAnswers []string `json:"correct_answers"`
	Commentary     string   `json:"commentary"`
}

func main() {
	var questions []Question

	// Ask the user which type of questions they want
	var loadKnowledgeQuestions, loadPracticeQuestions bool

	fmt.Print("Do you want knowledge questions? (yes/no): ")
	var knowledgeResponse string
	_, err := fmt.Scan(&knowledgeResponse)
	if err != nil {
		fmt.Println("Input error.")
		return
	}
	if strings.ToLower(knowledgeResponse) == "yes" {
		loadKnowledgeQuestions = true
	}

	fmt.Print("Do you want practice questions? (yes/no): ")
	var practiceResponse string
	_, err = fmt.Scan(&practiceResponse)
	if err != nil {
		fmt.Println("Input error.")
		return
	}
	if strings.ToLower(practiceResponse) == "yes" {
		loadPracticeQuestions = true
	}

	// Load questions based on user's choice
	if loadKnowledgeQuestions {
		knowledgeQuestions, err := loadQuestions("knowledge_questions.json")
		if err != nil {
			fmt.Println("Error loading knowledge questions:", err)
			return
		}
		questions = append(questions, knowledgeQuestions...)
	}

	if loadPracticeQuestions {
		practiceQuestions, err := loadQuestions("practice_questions.json")
		if err != nil {
			fmt.Println("Error loading practice questions:", err)
			return
		}
		questions = append(questions, practiceQuestions...)
	}

	if len(questions) == 0 {
		fmt.Println("No questions selected.")
		return
	}

	// Initialize the random number generator
	rand.Seed(time.Now().UnixNano())

	// Ask the user how many questions they want for the training
	var numberOfQuestions int
	fmt.Print("How many questions do you want for the training? ")
	_, err = fmt.Scan(&numberOfQuestions)
	if err != nil || numberOfQuestions <= 0 || numberOfQuestions > len(questions) {
		fmt.Println("Invalid number of questions.")
		return
	}

	// Shuffle the questions to present them in random order
	rand.Shuffle(len(questions), func(i, j int) {
		questions[i], questions[j] = questions[j], questions[i]
	})

	// Initialize the score
	score := 0

	// Ask questions to the user
	for i := 0; i < numberOfQuestions; i++ {
		question := questions[i]
		fmt.Printf("\nQuestion %d:\n%s\n", i+1, question.Prompt)

		// Display options
		for j, option := range question.Options {
			if option != "" {
				fmt.Printf("%c. %s\n", 'A'+j, option)
			}
		}

		// Ask the user for their answer
		var answer string
		fmt.Print("Your answer (A, B, C, D, etc.): ")
		_, err := fmt.Scan(&answer)
		if err != nil {
			fmt.Println("Input error.")
			return
		}

		// Check the answer
		answer = answer[0:1] // Take only the first character
		if answer == question.CorrectAnswers[0] {
			fmt.Println("Correct answer!")
			score++
		} else {
			fmt.Printf("Wrong answer. The correct answers were: %s\n", strings.Join(question.CorrectAnswers, ", "))
		}

		// Display the commentary if it exists
		if question.Commentary != "" {
			fmt.Println("Commentary:", question.Commentary)
		}
	}

	// Display the final score
	fmt.Printf("\nFinal score: %d/%d\n", score, numberOfQuestions)
}

func loadQuestions(filename string) ([]Question, error) {
	// Read the JSON file
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Parse the JSON file into a slice of questions
	var questions []Question
	err = json.Unmarshal(file, &questions)
	if err != nil {
		return nil, err
	}

	return questions, nil
}