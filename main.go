package main;

import (
	"fmt"
	"bufio"
	"os"
	// "strconv"
)

func main() {
	// scanner := bufio.NewScanner(os.Stdin);
	// scanner.Scan()
	// input, _ := strconv.ParseInt(scanner.Text(), 10, 64);
	// fmt.Printf("You typed: %d", 2020-input);

	// first scanner to identify username
	fmt.Println("Hello, What's your name?");
	scanner := bufio.NewScanner(os.Stdin);
	scanner.Scan()
	input := scanner.Text()

	if input != "" && input != " " {
		fmt.Printf("Hi " + input + "\n");
	} else {
		fmt.Println("Please type your name");
	}

	// second scanner to continue [Y/N]

	fmt.Println("You Want Continue? [Y/N]");
	continueScanner := bufio.NewScanner(os.Stdin)
	continueScanner.Scan()
	continueInput := continueScanner.Text()
	if continueInput == "Y" || continueInput == "y" {

	} else if continueInput == "N" || continueInput == "n" {
		fmt.Printf("You Choose: " + continueInput + " Thanks for using calculator");
	} else {
		fmt.Printf("That's bad choose format " + continueInput + " Bye.");
	}



	// fmt.Printf("You Want Continue? [Y/N]");
	// inpScanner := bufio.NewScanner(os.Stdin);
	// inpScanner.Scan()
	// inp := inpScanner.Text()
	// fmt.Printf(inp);
}
