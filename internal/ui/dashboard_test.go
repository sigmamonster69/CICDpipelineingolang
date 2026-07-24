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
	if want := "<title>Demo</title>"; !strings.Contains(html, want) {
		t.Fatalf("RenderDashboard() missing %q", want)
	}
}

func TestUpdateTaskStatus(t *testing.T) {
	tasks := DefaultTasks()
	tasks = UpdateTaskStatus(tasks, "Read the task board in the website", true)

	if !tasks[0].Done {
		t.Fatal("expected task to be marked done")
	}
	if tasks[0].DoneAt == "" {
		t.Fatal("expected done date to be set")
	}
}
