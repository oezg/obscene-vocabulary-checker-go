package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	/*
		Read the name of the file that contains taboo words from the user input;
		Read the file (back-up link is: words.);
		Read words from the input. If the word is taboo, replace it with * , according to the length of the words,
		and print the result, otherwise print the word. Repeat this until the exit command is entered;
		When the program receives the exit command, print Bye! and exit the program.
	*/

	//Take the name of the file as input from the user
	var filename string
	fmt.Println("Enter the name of the file that contains taboo words:")
	fmt.Scanln(&filename)

	//Open the file and defer closing the file to the end of the program
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//Initialize the scanner with the file and scan words
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	//Loop through the file word by word,
	//Add each word to a dictionary
	dictionary := make(map[string]struct{})
	for scanner.Scan() {
		dictionary[strings.ToLower(scanner.Text())] = struct{}{}
	}

	//Initialize the reader and loop continuously
	reader := bufio.NewReader(os.Stdin)
	for {
		//Read input until the new line
		fmt.Println("Enter a sentence to censor obscene words or 'exit':")
		sentence, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		//Get rid of the new line character at the end of the string
		sentence = strings.Trim(sentence, "\n")
		//Exit the loop if the input is 'exit'
		if "exit" == sentence {
			break
		}
		//Split the sentence into words
		words := strings.Split(sentence, " ")
		//Loop through the words and replace the obscene words with *'s
		for i, word := range words {
			if _, ok := dictionary[strings.ToLower(word)]; ok {
				words[i] = strings.Repeat("*", len(word))
			}
		}
		//Join the words back into a sentence and show it
		fmt.Println(strings.Join(words, " "))
	}
	//Finish the program with a bye
	fmt.Println("Bye!")
}
