package request

import "weakly-task-3/businesses/blogs"

type Blogs struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	CategoryID  uint   `json:"category_id"`
}

func (req *Blogs) ToDomain() *blogs.Domain {
	return &blogs.Domain{
		Title:       req.Title,
		Description: req.Description,
		CategoryID:  req.CategoryID,
	}
}
