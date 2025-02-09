package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"notion-htmx-blog/internal/domain"

	"github.com/jomei/notionapi"
)

type NotionRepository struct {
	client *notionapi.Client
	dbID   string
}

func NewNotionRepository(apiKey, dbID string) *NotionRepository {
	return &NotionRepository{
		client: notionapi.NewClient(notionapi.Token(apiKey)),
		dbID:   dbID,
	}
}

func (r *NotionRepository) ListReviews(limit int, tag domain.Tag) ([]domain.Review, error) {
	query := &notionapi.DatabaseQueryRequest{
		PageSize: limit,
		Sorts: []notionapi.SortObject{
			{
				Property:  "Created time",
				Direction: "descending",
			},
		},
	}

	// If tag is specified, filter by tag
	if tag != "" {
		query.Filter = &notionapi.PropertyFilter{
			Property: "Tag",
			MultiSelect: &notionapi.MultiSelectFilterCondition{
				Contains: string(tag),
			},
		}
	}

	log.Printf("Querying Notion with filter: %+v", query.Filter)
	result, err := r.client.Database.Query(context.Background(), notionapi.DatabaseID(r.dbID), query)
	if err != nil {
		return nil, fmt.Errorf("failed to query Notion database: %w", err)
	}

	reviews := make([]domain.Review, 0, len(result.Results))
	for _, page := range result.Results {
		review, err := r.mapPageToReview(&page)
		if err != nil {
			return nil, err
		}
		// Only include published reviews
		if review.Published {
			reviews = append(reviews, *review)
		}
	}

	return reviews, nil
}

func (r *NotionRepository) GetReview(id string) (*domain.Review, error) {
	page, err := r.client.Page.Get(context.Background(), notionapi.PageID(id))
	if err != nil {
		return nil, fmt.Errorf("failed to get Notion page: %w", err)
	}

	review, err := r.mapPageToReview(page)
	if err != nil {
		return nil, err
	}

	if !review.Published {
		return nil, fmt.Errorf("review not found: %s", id)
	}

	return review, nil
}

func (r *NotionRepository) GetReviewBySlug(slug string) (*domain.Review, error) {
	query := &notionapi.DatabaseQueryRequest{
		Filter: &notionapi.PropertyFilter{
			Property: "Slug",
			RichText: &notionapi.TextFilterCondition{
				Equals: slug,
			},
		},
	}

	result, err := r.client.Database.Query(context.Background(), notionapi.DatabaseID(r.dbID), query)
	if err != nil {
		return nil, fmt.Errorf("failed to query Notion database: %w", err)
	}

	if len(result.Results) == 0 {
		return nil, fmt.Errorf("review not found with slug: %s", slug)
	}

	review, err := r.mapPageToReview(&result.Results[0])
	if err != nil {
		return nil, err
	}

	if !review.Published {
		return nil, fmt.Errorf("review not found with slug: %s", slug)
	}

	return review, nil
}

func (r *NotionRepository) mapPageToReview(page *notionapi.Page) (*domain.Review, error) {
	review := &domain.Review{
		ID: page.ID.String(),
	}

	log.Printf("Mapping Notion page %s", page.ID)
	log.Printf("Available properties: %v", getPropertyNames(page.Properties))

	if prop, ok := page.Properties["Title"]; ok {
		if title, ok := prop.(*notionapi.TitleProperty); ok && len(title.Title) > 0 {
			review.Title = title.Title[0].PlainText
			log.Printf("Found title: %s", review.Title)
		} else {
			log.Printf("Title property is not in expected format: %T", prop)
		}
	} else {
		log.Printf("Title property not found")
	}

	if prop, ok := page.Properties["Cover Image"]; ok {
		if url, ok := prop.(*notionapi.URLProperty); ok {
			review.CoverImage = url.URL
			log.Printf("Found cover image: %s", review.CoverImage)
		} else {
			log.Printf("Cover Image property is not in expected format: %T", prop)
		}
	} else {
		log.Printf("Cover Image property not found")
	}

	if prop, ok := page.Properties["Slug"]; ok {
		if slug, ok := prop.(*notionapi.RichTextProperty); ok && len(slug.RichText) > 0 {
			review.Slug = slug.RichText[0].PlainText
			log.Printf("Found slug: %s", review.Slug)
		} else {
			log.Printf("Slug property is not in expected format: %T", prop)
		}
	} else {
		log.Printf("Slug property not found")
	}

	if prop, ok := page.Properties["Description"]; ok {
		if desc, ok := prop.(*notionapi.RichTextProperty); ok && len(desc.RichText) > 0 {
			review.Description = desc.RichText[0].PlainText
			log.Printf("Found description: %s", review.Description)
		} else {
			log.Printf("Description property is not in expected format: %T", prop)
		}
	} else {
		log.Printf("Description property not found")
	}

	if prop, ok := page.Properties["Published"]; ok {
		if published, ok := prop.(*notionapi.CheckboxProperty); ok {
			review.Published = published.Checkbox
			log.Printf("Found published status: %v", review.Published)
		} else {
			log.Printf("Published property is not in expected format: %T", prop)
		}
	} else {
		log.Printf("Published property not found")
	}

	if prop, ok := page.Properties["Date"]; ok {
		if date, ok := prop.(*notionapi.DateProperty); ok && date.Date != nil {
			parsedTime, err := time.Parse("2006-01-02", date.Date.Start.String()[:10])
			if err == nil {
				review.Date = parsedTime
				log.Printf("Found date: %v", review.Date)
			} else {
				log.Printf("Error parsing date: %v", err)
			}
		} else {
			log.Printf("Date property is not in expected format: %T", prop)
		}
	} else {
		log.Printf("Date property not found")
	}

	if prop, ok := page.Properties["Created time"]; ok {
		if created, ok := prop.(*notionapi.CreatedTimeProperty); ok {
			review.CreatedTime = created.CreatedTime
			log.Printf("Found created time: %v", review.CreatedTime)
		} else {
			log.Printf("Created time property is not in expected format: %T", prop)
		}
	} else {
		log.Printf("Created time property not found")
	}

	if prop, ok := page.Properties["Author"]; ok {
		if author, ok := prop.(*notionapi.RichTextProperty); ok && len(author.RichText) > 0 {
			review.Author = author.RichText[0].PlainText
			log.Printf("Found author: %s", review.Author)
		} else {
			log.Printf("Author property is not in expected format: %T", prop)
		}
	} else {
		log.Printf("Author property not found")
	}

	if prop, ok := page.Properties["Tag"]; ok {
		if tags, ok := prop.(*notionapi.MultiSelectProperty); ok {
			review.Tags = make([]domain.Tag, len(tags.MultiSelect))
			for i, tag := range tags.MultiSelect {
				review.Tags[i] = domain.Tag(tag.Name)
			}
			log.Printf("Found tags: %v", review.Tags)
		} else {
			log.Printf("Tag property is not in expected format: %T", prop)
		}
	} else {
		log.Printf("Tag property not found")
	}

	return review, nil
}

// Helper function to get property names
func getPropertyNames(props map[string]notionapi.Property) []string {
	names := make([]string, 0, len(props))
	for name := range props {
		names = append(names, name)
	}
	return names
}
