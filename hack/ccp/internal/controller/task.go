package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type tasks struct {
	Tasks map[uuid.UUID]*task `json:"tasks"`
}

func (t tasks) MarshalJSON() ([]byte, error) {
	if len(t.Tasks) == 0 {
		return nil, errors.New("map is empty")
	}
	type alias tasks
	return json.Marshal(&struct{ *alias }{alias: (*alias)(&t)})
}

type task struct {
	ID          uuid.UUID `json:"task_uuid"`
	CreatedAt   mcpTime   `json:"createdDate"`
	Args        string    `json:"arguments"`
	WantsOutput bool      `json:"wants_output"`
}

type taskResults struct {
	Results map[uuid.UUID]*taskResult `json:"task_results"`
}

type taskResult struct {
	ExitCode   int        `json:"exitCode"`
	FinishedAt pickleTime `json:"finishedTimestamp"`
	Stdout     string     `json:"stdout"`
	Stderr     string     `json:"stderr"`
}

// Custom time layout used by MCP.
//
// Using the reference Go time: 1136239445 (tz=MST):
//
//	>>> from datetime import datetime, timezone, timedelta
//	>>> datetime.fromtimestamp(1136239445, tz=timezone(timedelta(hours=-7))).isoformat(" ")
//	'2006-01-02 15:04:05-07:00'
type mcpTime struct {
	time.Time
}

const mcpTimeLayout = "2006-01-02 15:04:05-07:00"

func (t mcpTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Format(mcpTimeLayout))
}

func (t *mcpTime) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	pt, err := time.Parse(mcpTimeLayout, str)
	if err != nil {
		return err
	}
	t.Time = pt
	return nil
}

// Custom time layout used by the Gearman data encoder (regretfully).
type pickleTime struct {
	time.Time
}

type pickleTimeParts struct {
	Type        string `json:"__type__"`
	Year        int    `json:"year"`
	Month       int    `json:"month"`
	Day         int    `json:"day"`
	Hour        int    `json:"hour"`
	Minute      int    `json:"minute"`
	Second      int    `json:"second"`
	Microsecond int    `json:"microsecond"`
}

func (t *pickleTime) UnmarshalJSON(data []byte) error {
	parts := &pickleTimeParts{}
	if err := json.Unmarshal(data, &parts); err != nil {
		return err
	}

	if parts.Type != "datetime" {
		return fmt.Errorf("unexpected type: %q", parts.Type)
	}

	*t = pickleTime{
		time.Date(
			parts.Year, time.Month(parts.Month), parts.Day,
			parts.Hour, parts.Minute, parts.Second, parts.Microsecond,
			time.UTC,
		),
	}

	return nil
}
