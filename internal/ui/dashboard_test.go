package ui

import (
	"strings"
	"testing"
)

func TestRenderDashboard(t *testing.T) {
	html, err := RenderDashboard(PageData{
		Title:   "Demo",
		Message: "OK",
		Report:  "hello world",
	})
	if err != nil {
		t.Fatalf("RenderDashboard() error = %v", err)
	}
	
	if html == "" {
		t.Fatal("RenderDashboard() returned empty html")
	}
	
	want := "<title>Demo</title>"
	if !strings.Contains(html, want) {
		t.Fatalf("RenderDashboard() missing %q", want)
	}
	
	if !strings.Contains(html, "OK") {
		t.Fatal("RenderDashboard() missing status message")
	}
}

func TestUpdateTaskStatus(t *testing.T) {
	tasks := DefaultTasks()
	taskName := "Read the task board in the website"
	
	tasks = UpdateTaskStatus(tasks, taskName, true)

	var found bool
	for _, task := range tasks {
		if task.Name == taskName {
			found = true
			if !task.Done {
				t.Fatal("expected task to be marked done")
			}
			if task.DoneAt == "" {
				t.Fatal("expected done date to be set")
			}
			break
		}
	}
	
	if !found {
		t.Fatalf("task %q not found in default tasks", taskName)
	}
	
	tasks = UpdateTaskStatus(tasks, taskName, false)
	for _, task := range tasks {
		if task.Name == taskName {
			if task.Done {
				t.Fatal("expected task to be marked not done after undo")
			}
			if task.DoneAt != "" {
				t.Fatal("expected done date to be cleared after undo")
			}
			break
		}
	}
}
