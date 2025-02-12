package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"notion-htmx-blog/internal/config"
	"notion-htmx-blog/internal/handler"
	"notion-htmx-blog/internal/repository"
	"notion-htmx-blog/internal/service"
)

// customFileServer wraps http.FileServer to set correct MIME types
func customFileServer(dir string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set correct MIME type for CSS files
		if filepath.Ext(r.URL.Path) == ".css" {
			w.Header().Set("Content-Type", "text/css")
		}
		http.FileServer(http.Dir(dir)).ServeHTTP(w, r)
	})
}

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize templates with custom functions
	funcMap := template.FuncMap{
		"iterate": func(count int) []struct{} {
			return make([]struct{}, count)
		},
	}

	log.Printf("Loading templates from %s", cfg.TemplatesPath)
	tmpl := template.Must(template.New("").Funcs(funcMap).ParseFS(
		os.DirFS(cfg.TemplatesPath),
		"*.html",
	))
	log.Printf("Loaded templates: %v", tmpl.DefinedTemplates())

	// Initialize dependencies
	repo := repository.NewNotionRepository(cfg.NotionAPIKey, cfg.NotionDBID)
	svc := service.NewReviewService(repo)
	h := handler.NewReviewHandler(svc, tmpl)

	// Create router and register routes
	mux := http.NewServeMux()
	h.RegisterRoutes(mux)

	// Serve static files with correct MIME types
	mux.Handle("/static/", http.StripPrefix("/static/", customFileServer(cfg.StaticFilePath)))

	// Start server
	log.Printf("Server starting on :%s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, mux); err != nil {
		log.Fatal(err)
	}
}
