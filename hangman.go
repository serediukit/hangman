package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const MinLength = 4
const MaxLength = 15

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
	*toInit = rand.Intn(MaxLength-MinLength) + MaxLength
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
	switch gameType {

	if gameType == "e" {
		numberOfLetters = max(numberOfLetters, 6)
	}

	var letterData []byte
	var err error


	path := fmt.Sprintf("words/%dletters.txt", numberOfLetters)
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
	//stageOfHangman := 0
	gameType := ""
	//hasGuessedALetter := false
	//hasWon := false
	//guess := ""
	//guessedLetters := ""
	//again := false
	//dashes := ""
	//newDashes := ""

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
}
