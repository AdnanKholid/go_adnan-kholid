package blogs

import (
	"net/http"
	"weakly-task-3/businesses/blogs"
	"weakly-task-3/controllers/blogs/request"
	"weakly-task-3/controllers/blogs/response"

	"github.com/labstack/echo/v4"
)

type BlogsController struct {
	blogsUseCase blogs.Usecase
}

func NewBlogsController(blogsUC blogs.Usecase) *BlogsController {
	return &BlogsController{
		blogsUseCase: blogsUC,
	}
}

func (bc *BlogsController) GetAll(c echo.Context) error {
	blogsData := bc.blogsUseCase.GetAll()

	blogs := []response.Blogs{}

	for _, blog := range blogsData {
		blogs = append(blogs, response.FromDomain(blog))
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Alert": "Success Get Data",
		"Data":  blogs,
	})
}

func (bc *BlogsController) GetByID(c echo.Context) error {
	id := c.Param("id")

	blogs := bc.blogsUseCase.GetByID(id)

	if blogs.ID == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"Alert": "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Alert": "Success Get Data",
		"Data":  response.FromDomain(blogs),
	})

}

func (bc *BlogsController) Create(c echo.Context) error {
	input := request.Blogs{}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Alert": "Failed Create",
		})
	}

	blogs := bc.blogsUseCase.Create(input.ToDomain())

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"Alert": "Success Create",
		"Data":  response.FromDomain(blogs),
	})
}

func (bc *BlogsController) Update(c echo.Context) error {
	id := c.Param("id")
	input := request.Blogs{}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Alert": "Failed Create",
		})
	}
	blogs := bc.blogsUseCase.Update(id, input.ToDomain())

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"Alert": "Success Update",
		"Data":  response.FromDomain(blogs),
	})
}

func (bc *BlogsController) Delete(c echo.Context) error {
	id := c.Param("id")

	blogs := bc.blogsUseCase.Delete(id)

	if !blogs {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"Alert": "Failed Delete",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"Alert": "Success Delete",
	})
}

func (bc *BlogsController) GetByTitle(c echo.Context) error {
	title := c.Param("title")

	blogs := bc.blogsUseCase.GetByTitle(title)

	if blogs.Title == "" {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"Alert": "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Alert": "Success Get Data",
		"Data":  response.FromDomain(blogs),
	})

}

func (bc *BlogsController) GetByCategory(c echo.Context) error {
	idCategory := c.Param("category_id")

	blogsByCategory := bc.blogsUseCase.GetByCategory(idCategory)

	blogs := []response.Blogs{}

	for _, blog := range blogsByCategory {
		blogs = append(blogs, response.FromDomain(blog))
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Alert": "Success Get Data",
		"Data":  blogs,
	})

}
