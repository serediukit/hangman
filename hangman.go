package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	"time"
)

const MinLength = 4
const MaxLength = 5

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	wins := 0
	loses := 0

	numberOfLetters := 0

	reInitRandomValue(&numberOfLetters)
	again, hasWon := playHangman(numberOfLetters)

	for {
		if hasWon {
			wins++
		} else {
			loses++
		}

		printScore(wins, loses)

		if again {
			reInitRandomValue(&numberOfLetters)
			again, hasWon = playHangman(numberOfLetters)
		} else {
			break
		}
	}

	fmt.Println("Bye...")
}

func reInitRandomValue(toInit *int) {
	*toInit = rand.Intn(MaxLength-MinLength) + MinLength
}

func printScore(wins, loses int) {
	clearScreen()
	fmt.Printf("###########################\n")
	fmt.Printf("        Your score:        \n")
	fmt.Printf("        Wins:  %d          \n", wins)
	fmt.Printf("        Loses: %d          \n", loses)
	fmt.Printf("###########################\n")
	fmt.Println()
	return
}

func clearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	}
}

func chooseRandomWord(numberOfLetters int, gameType string) string {

	if gameType == "e" {
		numberOfLetters = min(numberOfLetters, 6)
	}

	var letterData []byte
	var err error

	path := fmt.Sprintf("words/%dletter.txt", numberOfLetters)
	letterData, err = os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	dataString := string(letterData)
	someWords := strings.Split(dataString, " ")
	randomNumber := rand.Intn(len(someWords) - 1)
	chosenWord := someWords[randomNumber]

	return chosenWord
}

func playHangman(numberOfLetters int) (playAgain bool, isWinner bool) {
	stageOfHangman := 0
	gameType := ""
	hasGuessedALetter := false
	hasWon := false
	guess := ""
	guessedLetters := ""
	//again := false
	dashes := ""
	newDashes := ""

	fmt.Printf("###########################\n")
	fmt.Printf("###    H A N G M A N    ###\n")
	fmt.Printf("###########################\n")
	fmt.Println()

	for {
		fmt.Println("Select game type:")
		fmt.Println("[e] Easy - use simple worlds (4, 5, 6 letters)")
		fmt.Println("[h] Hard - use more difficult words to 15 letters")
		fmt.Scanln(&gameType)
		gameType = strings.ToLower(gameType)

		if (gameType == "e") || (gameType == "h") {
			clearScreen()
			break
		} else {
			fmt.Println("Please, choose either 'e' or 'h' symbol")
		}
	}

	word := chooseRandomWord(numberOfLetters, gameType)

	fmt.Println()
	for {
		drawHangman(stageOfHangman)
		if stageOfHangman == 10 {
			fmt.Println("You lost the game")
			fmt.Printf("You could have saved him with word: %s\n", strings.ToUpper(word))
			return wantToPlayAgain(), false
		}

		if !hasGuessedALetter {
			dashes = hideTheWorld(len(word))
			fmt.Printf(" %s\n", dashes)
		} else {
			fmt.Printf(" %s\n", newDashes)
		}
		fmt.Printf("\nGuess the letter: ")
		fmt.Scanln(&guess)

		isALetter, err := regexp.MatchString("^[a-zA-Z]", guess)
		if err != nil {
			clearScreen()
			fmt.Println("Something went wrong")
			panic(err)
		}

		guess = strings.ToLower(guess)

		if !isALetter {
			clearScreen()
			fmt.Println("That is not a letter. Try again...")
		} else if len(guess) > 1 {
			clearScreen()
			fmt.Println("You entered more than 1 character. Try again...")
		} else if strings.Contains(guessedLetters, guess) {
			clearScreen()
			fmt.Println("You entered the guessed letter. Try again...")
		} else if strings.Contains(word, guess) {
			clearScreen()
			fmt.Println("Great! The letter you guessed is in the word!")
			guessedLetters += guess

			if !hasGuessedALetter {
				updatedDashes := revealDashes(word, guess, dashes)
				newDashes = updatedDashes
			} else {
				updatedDashes := revealDashes(word, guess, newDashes)
				newDashes = updatedDashes
			}

			hasGuessedALetter = true
			if newDashes == strings.ToUpper(word) {
				hasWon = true
			}

			if hasWon {
				clearScreen()
				fmt.Println("#######################################################")
				fmt.Println("###         C O N G R A T U L A T I O N S !         ###")
				fmt.Println("#######################################################")
				fmt.Println()
				fmt.Println(" _   _")
				fmt.Println("  \\0/")
				fmt.Println("   |")
				fmt.Println("   |")
				fmt.Println("  / \\")
				fmt.Println()
				fmt.Println("You have won!")
				fmt.Printf("The word was: %s\n", strings.ToUpper(word))
				fmt.Printf("You have passed hangman in %v of 10 guesses!\n", stageOfHangman)
				return wantToPlayAgain(), true
			}
		} else {
			clearScreen()
			fmt.Println("The letter is not in the word")
			stageOfHangman++
			guessedLetters += guess
		}
	}
}

func drawHangman(stageOfHangman int) {
	switch stageOfHangman {
	case 0:
		fmt.Println("   +---+")
		fmt.Println("       |")
		fmt.Println("       |")
		fmt.Println("       |")
		fmt.Println("       |")
		fmt.Println("       |")
		fmt.Println("       |")
		fmt.Println("       |")
		fmt.Printf(" ======== %v/10 Guesses\n", stageOfHangman)
		fmt.Println()
		break

	case 1:
		fmt.Println("   +---+")
		fmt.Println("   |   |")
		fmt.Println("   0   |")
		fmt.Println("       |")
		fmt.Println("       |")
		fmt.Println("       |")
		fmt.Println("       |")
		fmt.Println("       |")
		fmt.Printf(" ======== %v/10 Guesses\n", stageOfHangman)
		fmt.Println()
		break

	case 2:
		fmt.Println("   +---+")
		fmt.Println("   |   |")
		fmt.Println("   0   |")
		fmt.Println("   |   |")
		fmt.Println("       |")
		fmt.Println("       |")
		fmt.Println("       |")
		fmt.Println("       |")
		fmt.Printf(" ======== %v/10 Guesses\n", stageOfHangman)
		fmt.Println()
		break

	case 3:
		fmt.Println("   +---+")
		fmt.Println("   |   |")
		fmt.Println("   0   |")
		fmt.Println("  /|   |")
		fmt.Println("       |")
		fmt.Println("       |")
		fmt.Println("       |")
		fmt.Println("       |")
		fmt.Printf(" ======== %v/10 Guesses\n", stageOfHangman)
		fmt.Println()
		break

	case 4:
		fmt.Println("   +---+")
		fmt.Println("   |   |")
		fmt.Println("   0   |")
		fmt.Println(" _/|   |")
		fmt.Println("       |")
		fmt.Println("       |")
		fmt.Println("       |")
		fmt.Println("       |")
		fmt.Printf(" ======== %v/10 Guesses\n", stageOfHangman)
		fmt.Println()
		break

	case 5:
		fmt.Println("   +---+")
		fmt.Println("   |   |")
		fmt.Println("   0   |")
		fmt.Println(" _/|\\  |")
		fmt.Println("       |")
		fmt.Println("       |")
		fmt.Println("       |")
		fmt.Println("       |")
		fmt.Printf(" ======== %v/10 Guesses\n", stageOfHangman)
		fmt.Println()
		break

	case 6:
		fmt.Println("   +---+")
		fmt.Println("   |   |")
		fmt.Println("   0   |")
		fmt.Println(" _/|\\_ |")
		fmt.Println("       |")
		fmt.Println("       |")
		fmt.Println("       |")
		fmt.Println("       |")
		fmt.Printf(" ======== %v/10 Guesses\n", stageOfHangman)
		fmt.Println()
		break

	case 7:
		fmt.Println("   +---+")
		fmt.Println("   |   |")
		fmt.Println("   0   |")
		fmt.Println(" _/|\\_ |")
		fmt.Println("   |   |")
		fmt.Println("       |")
		fmt.Println("       |")
		fmt.Println("       |")
		fmt.Printf(" ======== %v/10 Guesses\n", stageOfHangman)
		fmt.Println()
		break

	case 8:
		fmt.Println("   +---+")
		fmt.Println("   |   |")
		fmt.Println("   0   |")
		fmt.Println(" _/|\\_ |")
		fmt.Println("   |   |")
		fmt.Println("  /    |")
		fmt.Println("       |")
		fmt.Println("       |")
		fmt.Printf(" ======== %v/10 Guesses\n", stageOfHangman)
		fmt.Println()
		break

	case 9:
		fmt.Println("   +---+")
		fmt.Println("   |   |")
		fmt.Println("   0   |")
		fmt.Println(" _/|\\_ |")
		fmt.Println("   |   |")
		fmt.Println("  / \\  |")
		fmt.Println("       |")
		fmt.Println("       |")
		fmt.Printf(" ======== %v/10 Guesses\n", stageOfHangman)
		fmt.Println()
		break

	case 10:
		fmt.Println("   +---+")
		fmt.Println("   |   |")
		fmt.Println("   0   |")
		fmt.Println("  /|\\  |")
		fmt.Println(" ° | ° |")
		fmt.Println("  / \\  |")
		fmt.Println("       |")
		fmt.Println(" R.I.P.|")
		fmt.Printf(" ======== %v/10 Guesses\n", stageOfHangman)
		fmt.Println()
		break
	}
}

func wantToPlayAgain() bool {
	for {
		var again string
		fmt.Println("Want to play again? [y/n]")
		fmt.Scanln(&again)
		isMatch, err := regexp.MatchString("^Y|y|N|n", again)
		if err != nil {
			fmt.Println("Something went wrong")
			panic(err)
		} else if !isMatch {
			fmt.Println("You need to type [y] or [n]. Try again...")
		} else if len(again) > 1 {
			fmt.Println("You entered more than 1 character. Try again...")
		} else if strings.ToLower(again) == "y" {
			return true
		} else if strings.ToLower(again) == "n" {
			return false
		} else {
			fmt.Println("Sorry, I don't know how to play again.")
		}
	}
}

func hideTheWorld(length int) string {
	dashes := ""
	for i := 0; i < length; i++ {
		dashes += "_"
	}
	return dashes
}

func revealDashes(word string, guess string, dashes string) string {
	newDashes := ""
	for i, r := range dashes {
		if c := string(r); c != "_" {
			newDashes += c
		} else {
			var letter = string(word[i])
			if guess == letter {
				newDashes += strings.ToUpper(guess)
			} else {
				newDashes += "_"
			}
		}
	}
	return newDashes
}
