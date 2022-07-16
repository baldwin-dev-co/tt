package main

import (
	"fmt"
	"os"
	"time"

	cli "github.com/jawher/mow.cli"
)

func main() {
	app := cli.App("tt", "time tracker")
	app.Spec = "SESSION [-d]"

	var (
		_ = app.StringArg("SESSION", "default", "The name of the session to store the data under")
		_ = app.StringOpt("d dir", "", "The directory to save the log file in")
	)

	app.Action = func() {
		// logPath, err := resolveLogPath(logDir)
		// if err != nil {
		// 	fmt.Println("Error finding log file, run using the [-d] flag to specify a log directory.")
		// 	cli.Exit(1)
		// }

		// log, err := readLog(logPath)
		// if err != nil {
		// 	create := BoolPrompt(
		// 		fmt.Sprintf(
		// 			"Log file does not exist in %s, would you like to create one?",
		// 			filepath.Dir(logPath),
		// 		),
		// 		true,
		// 	)

		// 	if !create {
		// 		cli.Exit(1)
		// 	}

		// 	log = new(Log)
		// }

		sesh := NewSesh()
		// (*log)[*seshKey] = append((*log)[*seshKey], sesh)

		idleEvents := make(chan bool)
		go func() {
			err := idleStateEmitter(idleEvents)
			if err != nil {
				fmt.Printf("Error listening for idle events: %v\n", err)
				cli.Exit(1)
			}

			close(idleEvents)
		}()

		for idle := range idleEvents {
			sesh.Pause()
			event := "Paused"
			if !idle {
				event = "Resumed"
			}
			fmt.Printf("%s at: %v\n", event, time.Now())
		}
	}

	app.Run(os.Args)
}
