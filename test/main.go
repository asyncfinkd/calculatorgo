package main;

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
	count := false;
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
		// second scanner to continue [Y/N]
		fmt.Println(input + ", " + "You Want Continue? [Y/N]");
		continueScanner := bufio.NewScanner(os.Stdin)
		continueScanner.Scan()
		continueInput := continueScanner.Text()
		if continueInput == "Y" || continueInput == "y" {
			/* Continue Calculator Section */
			if count != true {
				fmt.Printf("Choose Operator (+, -, *, /, %, \n");
				chooseOperator := bufio.NewScanner(os.Stdin);
				chooseOperator.Scan()
				chi := chooseOperator.Text();
				if chi != "" && chi != " " {
					fmt.Printf("You want to show enter result? [Y/N] \n");
					sr := bufio.NewScanner(os.Stdin);
					sr.Scan()
					sri := sr.Text()
					fmt.Printf("You Choose: " + sri + "\n");

					if sri == "Y" || sri == "y" || sri == "n" || sri == "N" {
						// component to building calc;
						// first number enter;
						firstNumberScanner := bufio.NewScanner(os.Stdin)
						fmt.Printf("Please enter a first number \n");
						firstNumberScanner.Scan()
						fnsi, _ := strconv.ParseInt(firstNumberScanner.Text(), 10, 64);
						if sri == "n" || sri == "N" {
							// empty;
						} else {
							fmt.Println(fnsi);
						}
						// second number enter;
						fmt.Printf("Please enter a second number \n");
						secondNumberScanner := bufio.NewScanner(os.Stdin);
						secondNumberScanner.Scan()
						snsi, _ := strconv.ParseInt(secondNumberScanner.Text(), 10, 64);
						if sri == "n" || sri == "N" {
							// empty;
						} else {
							fmt.Println(snsi);
						}
						switch chi {
							case "+":
								z := fnsi + snsi;
								fmt.Println(z);
								break;
							case "-":
								z := fnsi - snsi;
								fmt.Println(z);
								break;
							case "*":
								z := fnsi * snsi;
								fmt.Println(z);
								break;
							case "/":
								z := fnsi / snsi;
								fmt.Println(z);
								break;
							case "%":
								z := fnsi % snsi;
								fmt.Println(z);
								break;
							default:
								z := "choose bad format error";
								fmt.Printf(z);
						}
					} else {
						fmt.Printf("That's bad choose format " + sri + " bye.");
					}
					
					// fmt.Printf(input + ", " + "Please Enter a first number");
				} else {
					fmt.Println("You Choose " + chi + " and that's bad choose format bye.");
				}
			}
		} else if continueInput == "N" || continueInput == "n" {
			fmt.Printf("You Choose: " + continueInput + " Thanks for using calculator");
		} else {
			fmt.Printf("That's bad choose format " + continueInput + " Bye.");
		}
	} else {
		fmt.Println("Please type your name");
	}

	// fmt.Printf("You Want Continue? [Y/N]");
	// inpScanner := bufio.NewScanner(os.Stdin);
	// inpScanner.Scan()
	// inp := inpScanner.Text()
	// fmt.Printf(inp);
}
