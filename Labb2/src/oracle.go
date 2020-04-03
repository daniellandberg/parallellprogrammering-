// Daniel Landberg

// This program implements an ELIZA-like oracle (en.wikipedia.org/wiki/ELIZA).
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

const (
	star   = "Pythia"
	venue  = "Delphi"
	prompt = "> "
)

func main() {
	fmt.Printf("Welcome to %s, the oracle at %s.\n", star, venue)
	fmt.Println("Your questions will be answered in due time.")

	questions := Oracle()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fmt.Printf("%s heard: %s\n", star, line)
		questions <- line // The channel doesn't block.
	}
}
//skapade en funktion som tar in 2 channels för att få programmet i ett neverending state genom for loopen
//skickar sedan ned question som en sträng a och hela kanalen answer till prophecy
func questionsansw(questions chan string, answer chan string){

	for a := range questions {
		go prophecy(a, answer)
		go printa(answer)
	}
}
//funktion för att printa dem svar vi får på frågan.
func printa (answer chan string){
	c := <- answer
	fmt.Printf(c+"\n")
}
//funktion som konstant skickar ut var 30 sekund en random prophecy i en egen kanal.
func rando(questions chan string, answer chan string){
	for {
		time.Sleep(30 * time.Second)
		go prophecy("", answer)
		go printa(answer)

	}
}


// Oracle returns a channel on which you can send your questions to the oracle.
// You may send as many questions as you like on this channel, it never blocks.
// The answers arrive on stdout, but only when the oracle so decides.
// The oracle also prints sporadic prophecies to stdout even without being asked.

//la till answer som kanal
func Oracle() chan<- string {
	questions := make(chan string)
	answer := make(chan string)
	go questionsansw(questions, answer)
	go rando(questions, answer)
	// TODO: Answer questions.
	// TODO: Make prophecies.
	// TODO: Print answers.
	return questions
}

// This is the oracle's secret algorithm.
// It waits for a while and then sends a message on the answer channel.
// TODO: make it better.
func prophecy(question string, answer chan<- string) {
	// Keep them waiting. Pythia, the original oracle at Delphi,
	// only gave prophecies on the seventh day of each month.


	//la till några regex, väldigt enkla men går göra mer avancerade om man vill
	time.Sleep(time.Duration(2+rand.Intn(3)) * time.Second)
	match, _ := regexp.MatchString("live", question)
	match1, _ := regexp.MatchString("love", question)
	match2, _ := regexp.MatchString("i need (.*)", question)
	match3, _ := regexp.MatchString("grade|exam", question)
	if match3 ==true{
		bullshit3 := []string{
			"You will get an A",
			"A",
			"Study hard",
		}
		answer <- bullshit3[rand.Intn(len(bullshit3))]
	}
	if match2 ==true{
		bullshit2 := []string{
			"You will get it",
			"One day, you will get what you need",
			"did you know that deen is need backwards?",
		}
		answer <- bullshit2[rand.Intn(len(bullshit2))]
	}
	if match1 == true{
		bullshit1 := []string{
			"Everyone loves",
			"Love is in the air",
		}
		answer <- bullshit1[rand.Intn(len(bullshit1))]
	}
	if match ==true {
		bullshit := []string{
			"You will live forever",
			"You are a string person",
			"What is life?",
			"Life is life na na na na na",
		}
		answer <- bullshit[rand.Intn(len(bullshit))]

	}else{
	// Find the longest word.
	longestWord := ""
	words := strings.Fields(question) // Fields extracts the words into a slice.
	for _, w := range words {
		if len(w) > len(longestWord) {
			longestWord = w
		}
	}

	// Cook up some pointless nonsense.
	nonsense := []string{
		"The moon is dark.",
		"The sun is bright.",
	}
	answer <- longestWord + "... " + nonsense[rand.Intn(len(nonsense))]
}}

func init() { // Functions called "init" are executed before the main function.
	// Use new pseudo random numbers every time.
	rand.Seed(time.Now().Unix())
}