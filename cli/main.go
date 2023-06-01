package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Enter your first note")
	firstNote := getEnteredText()
	notes := []string{firstNote}

	fmt.Println("Note is saved")
	for {
		displayMenu()
		choice := getEnteredText()
		choice = choice[0:1]
		switch choice {
		case "1":
			note := getEnteredText()
			notes = append(notes, note)
			fmt.Println("Note is added")
		case "2":
			displayNotes(notes)
			fmt.Println("Which note would you like to update?")
			selectedNote := getEnteredText()
			selectedNote = selectedNote[0:1]
			fmt.Println("Enter a new value for a note")
			updatedNote := getEnteredText()
			selectedNumberInt, _ := strconv.Atoi(selectedNote)
			notes[selectedNumberInt-1] = updatedNote
			displayNotes(notes)
		case "3":
			displayNotes(notes)
			fmt.Println("Which note would you like to delete?")
			selectedNote := getEnteredText()
			selectedNote = selectedNote[0:1]
			fmt.Printf("Are you sure you want to delete this note? (y/n)")
			isConfirmNote := getEnteredText()
			isConfirmNote = isConfirmNote[0:1]
			selectedNumberInt, _ := strconv.Atoi(selectedNote)
			if isConfirmNote == "y" {
				notes = append(notes[:selectedNumberInt-1], notes[selectedNumberInt:]...)
				fmt.Println("Note deleted successfully.")
			} else if isConfirmNote == "n" {
				fmt.Println("Deletion canceled.")
			} else {
				fmt.Println("Wrong answer, please choose from y/n")

			}
		case "4":
			fmt.Println("Thank you for using the program!")
			return

		default:
			fmt.Println("not correct choice")
		}
	}
}

func getEnteredText() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	return text
}

func displayMenu() {
	fmt.Println("What would you like to do?")

	fmt.Println("1. Create a new note")
	fmt.Println("2. Update a previous note")
	fmt.Println("3. Delete a previous note")
	fmt.Println("4. Exit ")
}

func displayNotes(notes []string) {
	fmt.Println("List of your current notes:")
	for i, note := range notes {
		fmt.Printf("%d. %s", i+1, note)
	}
}
