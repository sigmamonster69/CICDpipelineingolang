package main

import (
	"log"
	"net/http"
	"os"

	"github.com/sigmamonster69/CICDpipelineingolang/internal/app"
	"github.com/sigmamonster69/CICDpipelineingolang/internal/ui"
)

func main() {
	stateDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err) // all used for the ci cd pieplines
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tasks, err := ui.LoadTasks(stateDir)
		if err != nil { 
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} //all good 

		taskName := r.URL.Query().Get("task")
		if taskName != "" {
			if r.URL.Query().Get("done") == "1" {
				tasks = ui.UpdateTaskStatus(tasks, taskName, true)
			}
			if r.URL.Query().Get("undo") == "1" {
				tasks = ui.UpdateTaskStatus(tasks, taskName, false)
			}
			if err := ui.SaveTasks(stateDir, tasks); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

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

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write([]byte(html))
	})

	log.Println("dashboard running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
