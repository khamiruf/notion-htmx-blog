package domain

import "time"

type Tag string

const (
	TagBook    Tag = "book"
	TagArticle Tag = "article"
)

// Review represents a book review entity
type Review struct {
	ID          string
	Title       string
	CoverImage  string
	Slug        string
	Description string
	Published   bool
	Date        time.Time
	CreatedTime time.Time
	Author      string
	Tags        []Tag
}

// ReviewRepository defines the interface for review data operations
type ReviewRepository interface {
	ListReviews(limit int, tag Tag) ([]Review, error)
	GetReview(id string) (*Review, error)
	GetReviewBySlug(slug string) (*Review, error)
}

// ReviewService defines the interface for review business logic
type ReviewService interface {
	ListReviews(limit int, tag Tag) ([]Review, error)
	GetReview(id string) (*Review, error)
	GetReviewBySlug(slug string) (*Review, error)
}
