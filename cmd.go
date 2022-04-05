package menu

import "fmt"

func main() {

	for {
		var cmd string
		fmt.Scan(&cmd)
		switch cmd {
		case "help":
			fmt.Println("*This is a help cmd! You can enter the following command.")
			fmt.Println("*help\tTo get some help about commands.")
			fmt.Println("*version\tGet the current version of the menu.")
			fmt.Println("*quit\tExit the menu.")

		case "version":
			fmt.Println("This version is 1.1")

		case "quit":
			return
		default:
			fmt.Println("It's a wrong cmd")
		}

	}
}
