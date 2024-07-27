package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
)

const (
	markName  = "REMINDER_TOOL"
	markValue = "1"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Printf("USAGE: <hh:mm> <text message>\n")
		os.Exit(1)
	}

	now := time.Now()

	w := when.New(nil)
	w.Add(en.All...)
	w.Add(common.All...)

	r, err := w.Parse(os.Args[1], time.Now())
	if err != nil {
		panic(err)
	}
	if r == nil {
		fmt.Println(fmt.Errorf("couldn't read the text provided"))
		os.Exit(2)
	}
	if now.After(r.Time) {
		fmt.Println(fmt.Errorf("please set future time"))
		os.Exit(3)
	}

	diff := r.Time.Sub(now)

	if os.Getenv(markName) == markValue {
		time.Sleep(diff)
		err := beeep.Alert("Reminder", strings.Join(os.Args[2:], " "), "assets/information")
		if err != nil {
			fmt.Println(err)
			os.Exit(4)
		}
	} else {
		cmd := exec.Command(os.Args[0], os.Args[1:]...)
		cmd.Env = append(os.Environ(), fmt.Sprintf("%s=%s", markName, markValue))
		if err := cmd.Start(); err != nil {
			fmt.Println(err)
			os.Exit(5)
		}

		fmt.Println("Reminder will be run in ", diff.Round(time.Second))
		os.Exit(0)
	}

}
