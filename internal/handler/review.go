package handler

import (
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"

	"notion-htmx-blog/internal/domain"
)

type ReviewHandler struct {
	service domain.ReviewService
	tmpl    *template.Template
}

func NewReviewHandler(service domain.ReviewService, tmpl *template.Template) *ReviewHandler {
	return &ReviewHandler{
		service: service,
		tmpl:    tmpl,
	}
}

func (h *ReviewHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", h.handleHome())
	mux.HandleFunc("/books", h.handleBooks())
	mux.HandleFunc("/articles", h.handleArticles())
	mux.HandleFunc("/food", h.handleFood())
	mux.HandleFunc("/reviews/", h.handleGetReview())
}

func (h *ReviewHandler) handleHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		data := map[string]interface{}{
			"Year":     time.Now().Year(),
			"Template": "about",
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		// Check if this is an HTMX request
		if r.Header.Get("HX-Request") == "true" {
			log.Printf("HTMX request detected, rendering about template")
			if err := h.tmpl.ExecuteTemplate(w, "about", data); err != nil {
				log.Printf("Error executing about template: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			log.Printf("About template rendered successfully")
		} else {
			log.Printf("Regular request detected, rendering base template")
			if err := h.tmpl.ExecuteTemplate(w, "base", data); err != nil {
				log.Printf("Error executing base template: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			log.Printf("Base template rendered successfully")
		}
	}
}

func (h *ReviewHandler) handleBooks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Fetching book reviews...")
		reviews, err := h.service.ListReviews(10, []domain.Tag{domain.TagBook})
		if err != nil {
			log.Printf("Error fetching book reviews: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		log.Printf("Got %d book reviews", len(reviews))
		data := map[string]interface{}{
			"Year":    time.Now().Year(),
			"Reviews": reviews,
			"Type":    "Book",
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		// Check if this is an HTMX request
		if r.Header.Get("HX-Request") == "true" {
			log.Printf("HTMX request detected, rendering content template")
			if err := h.tmpl.ExecuteTemplate(w, "content", data); err != nil {
				log.Printf("Error executing content template: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			log.Printf("Content template rendered successfully")
		} else {
			log.Printf("Regular request detected, rendering base template")
			if err := h.tmpl.ExecuteTemplate(w, "base", data); err != nil {
				log.Printf("Error executing base template: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			log.Printf("Base template rendered successfully")
		}
	}
}

func (h *ReviewHandler) handleArticles() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Fetching article reviews...")
		reviews, err := h.service.ListReviews(10, []domain.Tag{domain.TagArticle})
		if err != nil {
			log.Printf("Error fetching article reviews: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		log.Printf("Got %d article reviews", len(reviews))
		data := map[string]interface{}{
			"Year":    time.Now().Year(),
			"Reviews": reviews,
			"Type":    "Article",
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		// Check if this is an HTMX request
		if r.Header.Get("HX-Request") == "true" {
			log.Printf("HTMX request detected, rendering content template")
			if err := h.tmpl.ExecuteTemplate(w, "content", data); err != nil {
				log.Printf("Error executing content template: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			log.Printf("Content template rendered successfully")
		} else {
			log.Printf("Regular request detected, rendering base template")
			if err := h.tmpl.ExecuteTemplate(w, "base", data); err != nil {
				log.Printf("Error executing base template: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			log.Printf("Base template rendered successfully")
		}
	}
}

func (h *ReviewHandler) handleFood() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Fetching food reviews...")

		// Get cuisine from query parameter
		cuisine := r.URL.Query().Get("cuisine")
		tags := []domain.Tag{domain.TagFood}

		// Add cuisine tag if specified
		if cuisine != "" {
			tags = append(tags, domain.Tag(cuisine))
		}

		reviews, err := h.service.ListReviews(10, tags)
		if err != nil {
			log.Printf("Error fetching food reviews: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		log.Printf("Got %d food reviews", len(reviews))
		data := map[string]interface{}{
			"Year":           time.Now().Year(),
			"Reviews":        reviews,
			"Type":           "Food",
			"CurrentCuisine": cuisine,
			"CuisineTags": []string{
				"thai",
				"italian",
				"japanese",
				"chinese",
				"indian",
			},
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		// Check if this is an HTMX request
		if r.Header.Get("HX-Request") == "true" {
			log.Printf("HTMX request detected, rendering food template")
			if err := h.tmpl.ExecuteTemplate(w, "content", data); err != nil {
				log.Printf("Error executing food template: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			log.Printf("Content template rendered successfully")
		} else {
			log.Printf("Regular request detected, rendering base template")
			if err := h.tmpl.ExecuteTemplate(w, "base", data); err != nil {
				log.Printf("Error executing base template: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			log.Printf("Base template rendered successfully")
		}
	}
}

func (h *ReviewHandler) handleGetReview() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slug := strings.TrimPrefix(r.URL.Path, "/reviews/")
		if slug == "" {
			http.NotFound(w, r)
			return
		}

		review, err := h.service.GetReviewBySlug(slug)
		if err != nil {
			log.Printf("Error fetching review: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		data := map[string]interface{}{
			"Year":   time.Now().Year(),
			"Review": review,
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		// Check if this is an HTMX request
		if r.Header.Get("HX-Request") == "true" {
			if err := h.tmpl.ExecuteTemplate(w, "review_content", data); err != nil {
				log.Printf("Error executing template: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
		} else {
			if err := h.tmpl.ExecuteTemplate(w, "base", data); err != nil {
				log.Printf("Error executing template: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
		}
	}
}
