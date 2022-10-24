package response

import (
	"time"
	"weakly-task-3/businesses/blogs"
)

type Blogs struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CategoryID  uint      `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func FromDomain(domain blogs.Domain) Blogs {
	return Blogs{
		ID:          domain.ID,
		Title:       domain.Title,
		Description: domain.Description,
		CategoryID:  domain.CategoryID,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}
