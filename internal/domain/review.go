package domain

import "time"

type Tag string

const (
	TagBook    Tag = "book"
	TagArticle Tag = "article"
	TagFood    Tag = "food"
	// Cuisine tags
	TagThai     Tag = "thai"
	TagItalian  Tag = "italian"
	TagJapanese Tag = "japanese"
	TagChinese  Tag = "chinese"
	TagIndian   Tag = "indian"
)

// Review represents a book review entity
type Review struct {
	ID          string
	Title       string
	CoverImage  string
	URL         string
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
	ListReviews(limit int, tags []Tag) ([]Review, error)
	GetReview(id string) (*Review, error)
	GetReviewBySlug(slug string) (*Review, error)
}

// ReviewService defines the interface for review business logic
type ReviewService interface {
	ListReviews(limit int, tags []Tag) ([]Review, error)
	GetReview(id string) (*Review, error)
	GetReviewBySlug(slug string) (*Review, error)
}
