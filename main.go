package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("invalid input, unless one command after cle")
	}
	args := os.Args[1:]
	argStr := ""
	for _, arg := range args {
		argStr += arg + " "
	}
	printResultWithGptCompletion(argStr)

	fmt.Printf("\nenter y to continue, or exit: ")
	var input string
	fmt.Scanln(&input)
	if input == "y" {
		cmd := exec.Command(os.Args[1], os.Args[2:]...)
		out, _ := cmd.CombinedOutput()
		fmt.Println(strings.TrimSpace(string(out)))
	}
	fmt.Println("over")
}
