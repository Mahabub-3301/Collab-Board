package main

import (
	"errors"
	"strings"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func validate(t Task) error {
	if strings.TrimSpace(t.Title) == "" {
		return errors.New("title is required")
	}

	switch t.Status {
	case "todo", "in-progress", "done":
		return nil

	default:
		return errors.New("invalid status: must be todo, in-progress or done")
	}
}n
