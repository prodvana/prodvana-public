package cmdutil

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	common_config_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/common_config"
)

func Prompt(message string) string {
	// fmt.Scan* do not handle whitespace correctly.
	// https://stackoverflow.com/questions/43843477/scanln-in-golang-doesnt-accept-whitespace
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("%s: ", message)
	scanner.Scan()
	if scanner.Err() != nil {
		panic(fmt.Sprintf("Unexpected error reading input: %+v", scanner.Err()))
	}
	return scanner.Text()
}

func PromptYesNo(message string) bool {
	var answer string
	fmt.Printf("%s [Y/n]: ", message)
	fmt.Scanln(&answer)
	answer = strings.ToLower(strings.TrimSpace(answer))
	if answer != "" && answer != "y" && answer != "yes" {
		return false
	}
	return true
}

func PromptDangerousActions(actions []*common_config_pb.DangerousAction) {
	first := true
	for {
		var message string
		if first {
			var explanations []string
			for _, action := range actions {
				explanations = append(explanations, "  "+action.Explanation)
			}
			message = fmt.Sprintf("The following dangerous actions will be performed:\n\n%s\n\nPlease type \"I AM SURE\" in all caps to continue, or ctrl-c to abort", strings.Join(explanations, "\n"))
		} else {
			message = "Invalid input. Please type \"I AM SURE\" in all caps to continue, or ctrl-c to abort"
		}
		first = false
		answer := Prompt(message)
		if strings.TrimSpace(answer) == "I AM SURE" {
			return
		}
	}
}
