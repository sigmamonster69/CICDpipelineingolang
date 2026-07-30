package ui

import (
	"strings"
	"testing"
)

// TestRenderDashboard verifies that the RenderDashboard function generates valid HTML.
// This test ensures the template renders correctly and includes expected elements.
// The dashboard is a key UI component for tracking CI/CD learning progress.
func TestRenderDashboard(t *testing.T) {
	// Render dashboard with sample data
	html, err := RenderDashboard(PageData{
		Title:   "Demo",
		Message: "OK",
		Report:  "hello world",
	})
	if err != nil {
		t.Fatalf("RenderDashboard() error = %v", err)
	}
	
	// Verify HTML output is not empty
	if html == "" {
		t.Fatal("RenderDashboard() returned empty html")
	}
	
	// Check that the title appears in the rendered HTML
	want := "<title>Demo</title>"
	if !strings.Contains(html, want) {
		t.Fatalf("RenderDashboard() missing %q", want)
	}
	
	// Additional check: verify message appears in output
	if !strings.Contains(html, "OK") {
		t.Fatal("RenderDashboard() missing status message")
	}
}

// TestUpdateTaskStatus verifies that task completion status updates correctly.
// This test validates that marking a task done sets both the Done flag and DoneAt timestamp.
// Task state persistence is important for user progress tracking.
func TestUpdateTaskStatus(t *testing.T) {
	// Get default task list
	tasks := DefaultTasks()
	taskName := "Read the task board in the website"
	
	// Mark the first task as done
	tasks = UpdateTaskStatus(tasks, taskName, true)

	// Find the updated task and verify it's marked done
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
	
	// Test undoing (marking as not done)
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
