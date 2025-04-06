package service

import (
	"notion-htmx-blog/internal/domain"
)

type ReviewService struct {
	repo domain.ReviewRepository
}

func NewReviewService(repo domain.ReviewRepository) *ReviewService {
	return &ReviewService{
		repo: repo,
	}
}

func (s *ReviewService) ListReviews(limit int, tags []domain.Tag) ([]domain.Review, error) {
	if limit <= 0 {
		limit = 10 // Default limit
	}
	return s.repo.ListReviews(limit, tags)
}

func (s *ReviewService) GetReview(id string) (*domain.Review, error) {
	return s.repo.GetReview(id)
}

func (s *ReviewService) GetReviewBySlug(slug string) (*domain.Review, error) {
	return s.repo.GetReviewBySlug(slug)
}
