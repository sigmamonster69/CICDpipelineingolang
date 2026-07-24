package ui

import (
\t"strings"
\t"testing"
)

func TestRenderDashboard(t *testing.T) {
\thtml, err := RenderDashboard(PageData{
\t\tTitle:   "Demo",
\t\tMessage: "OK",
\t\tReport:  "hello world",
\t})
\tif err != nil {
\t\tt.Fatalf("RenderDashboard() error = %v", err)
\t}
\tif html == "" {
\t\tt.Fatal("RenderDashboard() returned empty html")
\t}
\tif want := "<title>Demo</title>"; !strings.Contains(html, want) {
\t\tt.Fatalf("RenderDashboard() missing %q", want)
\t}
}

func TestUpdateTaskStatus(t *testing.T) {
\ttasks := DefaultTasks()
\ttasks = UpdateTaskStatus(tasks, "Read ROADMAP.txt", true)

\tif !tasks[0].Done {
\t\tt.Fatal("expected task to be marked done")
\t}
\tif tasks[0].DoneAt == "" {
\t\tt.Fatal("expected done date to be set")
\t}
}
