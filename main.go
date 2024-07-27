package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
)

func reminder(input string) {
	inputArray := strings.Split(input, " ")

	if len(inputArray) < 2 {
		fmt.Printf("Please enter in following format <hh:mm> <text message> for this reminder->%s\n", input)
		return
	}

	now := time.Now()

	w := when.New(nil)
	w.Add(en.All...)
	w.Add(common.All...)

	r, err := w.Parse(inputArray[0], time.Now())
	if err != nil {
		fmt.Printf("Error while reading time for this reminder->%s\n", input)
		return
	}
	if r == nil {
		fmt.Printf("couldn't read the text provided for this reminder->%s\n", input)
		return
	}
	if now.After(r.Time) {
		fmt.Printf("Time should be of future for this reminder->%s\n", input)
		return
	}

	diff := r.Time.Sub(now)

	fmt.Println("Reminder will be run in ", diff.Round(time.Second))
	time.Sleep(diff)

	err = beeep.Alert("Reminder", strings.Join(inputArray[1:], " "), "assets/information")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {

	var takeInput = true

	fmt.Println("*********Welcome to GO REMINDER TOOL**********")
	fmt.Println("USAGE: Please enter your time and note to add a reminder in following format\n <hh:mm> <text message>\n For exiting the program please enter 'exit'")

inputLoop:
	for takeInput {
		var input string
		reader := bufio.NewReader(os.Stdin)
		input, _ = reader.ReadString('\n')
		// have to use trimspace since we reading till \n
		if strings.TrimSpace(input) == "exit" {
			takeInput = false
			break inputLoop
		}
		go reminder(input)
	}

}
