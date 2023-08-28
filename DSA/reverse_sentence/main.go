package main

import (
	"fmt"
	"strings"
)

func reverseSentence(sentence string) string {
	words := strings.Split(sentence, " ")
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func reverseSentenceWithReverseWord(sentence string) string {
	words := strings.Split(sentence, " ")
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		// fmt.Println(words[i])
		// fmt.Println(words[j])
		wordj, wordi := ReverseTheWord(words[j], words[i])
		words[i], words[j] = wordj, wordi
	}
	return strings.Join(words, " ")
}
func ReverseTheWord(wordj string, wordi string) (string, string) {
	runei := []rune(wordi)
	runej := []rune(wordj)
	for i, j := 0, len(runei)-1; i < j; i, j = i+1, j-1 {
		runei[i], runei[j] = runei[j], runei[i]
	}
	for i, j := 0, len(runej)-1; i < j; i, j = i+1, j-1 {
		runej[i], runej[j] = runej[j], runej[i]
	}
	return string(runej), string(runei)
}

func main() {
	input := "This is a sample sentence to reverse"
	reversed := reverseSentence(input)
	fmt.Println("Original sentence:", input)
	fmt.Println("Reversed sentence:", reversed)
	reversedSentenceWithWord := reverseSentenceWithReverseWord(input)
	fmt.Println("Reversed sentence with reversed word:", reversedSentenceWithWord)
}
