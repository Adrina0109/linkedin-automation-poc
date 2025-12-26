package storage

import (
	"encoding/json"
	"os"
	"time"
)

type Report struct {
	Timestamp       time.Time `json:"timestamp"`
	Logins          int       `json:"logins"`
	ProfilesVisited int       `json:"profiles_visited"`
	ConnectionsSent int       `json:"connections_sent"`
	MessagesSent    int       `json:"messages_sent"`
}

func SaveReport(r Report) error {
	data, _ := json.MarshalIndent(r, "", "  ")
	return os.WriteFile("report.json", data, 0644)
}
