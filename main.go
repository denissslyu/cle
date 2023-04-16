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

	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	out, _ := cmd.CombinedOutput()
	fmt.Println("\nCombined out:\n", strings.TrimSpace(string(out)))
}
