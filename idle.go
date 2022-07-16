package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func getDarwinIdle() (idle time.Duration, err error) {
	ioRegOutput, err := exec.Command("ioreg", "-c", "IOHIDSystem").Output()
	if err != nil {
		return idle, err
	}

	lines := strings.Split(string(ioRegOutput), "\n")
	var idleInNs string

	for _, line := range lines {
		if !strings.Contains(line, "HIDIdleTime") {
			continue
		}

		cols := strings.Split(line, " ")
		idleInNs = cols[len(cols)-1]
		break
	}

	idle, err = time.ParseDuration(fmt.Sprintf("%sns", idleInNs))
	if err != nil {
		return idle, fmt.Errorf("Error parsing idle duration: %v", err)
	}

	return idle, nil
}

const IDLE_THRESHOLD = 5 * time.Second

func idleStateEmitter(events chan bool) error {
	var wasIdle bool
	for {
		idleDuration, err := getDarwinIdle()
		if err != nil {
			return fmt.Errorf("Error getting idle time: %v", err)
		}

		isIdle := idleDuration > IDLE_THRESHOLD
		if wasIdle != isIdle {
			events <- isIdle
			wasIdle = isIdle
		}

		if isIdle {
			time.Sleep(time.Second)
		} else  {
			time.Sleep(IDLE_THRESHOLD - idleDuration)
		}
	}
}