package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

var connectedPattern *regexp.Regexp = regexp.MustCompile(`Got connection SteamID (\d+)`)
var disconnectedPattern *regexp.Regexp = regexp.MustCompile(`Closing socket (\d+)`)

func whosOnline() (users map[string]struct{}) {
	scanner := bufio.NewScanner(os.Stdin)
	users = make(map[string]struct{})

	for scanner.Scan() {
		line := scanner.Text()

		if match := connectedPattern.FindStringSubmatch(line); match != nil {
			users[match[1]] = struct{}{}
		} else if match := disconnectedPattern.FindStringSubmatch(line); match != nil {
			delete(users, match[1])
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error in scanner", err)
	}

	return
}

func main() {
	who := whosOnline()

	for user := range who {
		fmt.Println(user)
	}
}
