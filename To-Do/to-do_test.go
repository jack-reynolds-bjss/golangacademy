package main

import (
	"testing"
)

type ToDoExample struct {
	ToDo ToDo
	Expected string
}

var ToDoExamples = []ToDoExample{
	{ToDo{ Status: ToDoStatusNotStarted }, "Not Started \u274c"},
	{ToDo{ Status: ToDoStatusInProgress }, "In Progress \u25D4"},
	{ToDo{ Status: ToDoStatusCompleted }, "Completed \u2705"},
	{ToDo{}, "Unknown \u2753"},
}

func TestDoDo(t *testing.T) {
	t.Run("GetStatus work correctly", func(t *testing.T) {
		for _, todo := range ToDoExamples {
			result := todo.ToDo.GetStatus()
			if result != todo.Expected {
				t.Error(
					"For", todo.ToDo.Status,
					"Expected", todo.Expected,
					"Got", result,
				)
			}
		}
	})
}