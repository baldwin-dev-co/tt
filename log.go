package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Log map[string][]*Session

const LOG_FILE_NAME = "ttlog.json"

func resolveLogPath(dir *string) (path string, err error) {
	if dir != nil && *dir != "" {
		return filepath.Join(*dir, LOG_FILE_NAME), nil
	}

	cwd, err := os.Getwd()
	if err != nil {
		return path, fmt.Errorf("Error resolving cwd: %v", err)
	}

	// step through the cwd and each of its parents looking for a ttlog.json file
	for dir := cwd; dir != "/"; dir = filepath.Dir(dir) {
		path = filepath.Join(dir, LOG_FILE_NAME)
		if _, err := os.Stat(path); err == nil {
			return path, nil
		}
	}

	return path, fmt.Errorf("No log file found")
}

func readLog(path string) (log *Log, err error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Error reading file %s: %v", path, err)
	}

	err = json.Unmarshal(file, &log)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshaling json: %v", err)
	}

	return log, nil
}

func writeLog(path string, log *Log) error {
	data, err := json.Marshal(log)
	if err != nil {
		return fmt.Errorf("Error marshaling json %v", err)
	}

	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return fmt.Errorf("Error writing file: %v", err)
	}

	return nil
}

func logSession(log *Log, subject string, sesh *Session) {
	(*log)[subject] = append((*log)[subject], sesh)
}

