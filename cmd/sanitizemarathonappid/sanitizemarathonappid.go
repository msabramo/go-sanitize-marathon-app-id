package main

import (
	"fmt"
	"os"

	sanitizemarathonappid "github.com/msabramo/go-sanitize-marathon-app-id"
)

func process(args []string) (outputs []string) {
	for _, arg := range args {
		sanitizedMarathonAppID := sanitizemarathonappid.Sanitize(arg)
		outputs = append(outputs, sanitizedMarathonAppID)
	}
	return outputs
}

func main() {
	outputs := process(os.Args[1:])
	for _, output := range outputs {
		fmt.Println(output)
	}
}
