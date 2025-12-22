package storage

import (
	"encoding/json"
	"os"
	"sync"
	"time"
)

type LoginRecord struct {
	Timestamp string `json:"timestamp"`
	Status    string `json:"status"`
}

type State struct {
	Logins []LoginRecord `json:"logins"`
}

var mu sync.Mutex

const stateFile = "state.json"

// LoadState WITHOUT locking (caller controls locking)
func loadStateUnsafe() (*State, error) {
	if _, err := os.Stat(stateFile); os.IsNotExist(err) {
		return &State{Logins: []LoginRecord{}}, nil
	}

	data, err := os.ReadFile(stateFile)
	if err != nil {
		return nil, err
	}

	var state State
	if err := json.Unmarshal(data, &state); err != nil {
		return nil, err
	}

	return &state, nil
}

func SaveLogin(status string) error {
	mu.Lock()
	defer mu.Unlock()

	state, err := loadStateUnsafe()
	if err != nil {
		return err
	}

	record := LoginRecord{
		Timestamp: time.Now().Format(time.RFC3339),
		Status:    status,
	}

	state.Logins = append(state.Logins, record)

	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(stateFile, data, 0644)
}
