package main

import (
	"log"
	"net/http"
	"os"

	"github.com/sigmamonster69/CICDpipelineingolang/internal/app"
	"github.com/sigmamonster69/CICDpipelineingolang/internal/ui"
)

// main is the entry point for the cicd-dashboard web server.
// This application serves an interactive HTML dashboard for tracking CI/CD learning progress.
// The dashboard runs on port 8080 and provides a task board with persistent state storage.
// Users can mark tasks as done/not done, and the state is saved to a JSON file.
// The CI pipeline builds this binary and can deploy it to a web server or container.
func main() {
	// Get current working directory for storing task state file
	// The .task-state.json file will be created in this directory
	stateDir, err := os.Getwd()
	if err != nil {
		// Log fatal error and exit if we can't determine the working directory
		log.Fatal(err)
	}

	// Register HTTP handler for the root path "/"
	// This handler serves the dashboard HTML and handles task state updates
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Load current task state from disk (or use defaults if file doesn't exist)
		tasks, err := ui.LoadTasks(stateDir)
		if err != nil { 
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} //all good 

		// Check if request includes task update parameters
		taskName := r.URL.Query().Get("task")
		if taskName != "" {
			// Handle "mark as done" action
			if r.URL.Query().Get("done") == "1" {
				tasks = ui.UpdateTaskStatus(tasks, taskName, true)
			}
			// Handle "undo/mark as not done" action
			if r.URL.Query().Get("undo") == "1" {
				tasks = ui.UpdateTaskStatus(tasks, taskName, false)
			}
			// Save updated task state to disk
			if err := ui.SaveTasks(stateDir, tasks); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			// Redirect to refresh the page and show updated state
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		// Render the dashboard HTML with current application state
		html, err := ui.RenderDashboard(ui.PageData{
			Title:       "CI/CD learning dashboard",
			Description: "A tiny browser view for the Go CI/CD learning project.",
			Message:     app.Message(),
			Report:      app.BuildReport("sam", 2, 3),
			Tasks:       tasks,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set response content type header for HTML
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		// Write the rendered HTML to the response
		_, _ = w.Write([]byte(html))
	})

	// Log startup message indicating server is ready
	log.Println("dashboard running on http://localhost:8080")
	
	// Start HTTP server on port 8080
	// This call blocks until the server stops or encounters an error
	log.Fatal(http.ListenAndServe(":8080", nil))
}
