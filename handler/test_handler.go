package handler

import (
	"net/http"

	"go-tech-blog/repository"

	"github.com/labstack/echo/v4"
)

// Test ...
func Test(c echo.Context) error {
	article1, _ := repository.ArticleGetWithWriterName(9001)
	article2, _ := repository.ArticleGetWithWriterName(9002)
	article3, _ := repository.ArticleGetWithWriterName(9003)
	article4, _ := repository.ArticleGetWithWriterName(9004)
	article5, _ := repository.ArticleGetWithWriterName(9005)

	article6, _ := repository.ArticleGetWithWriter(9001)
	article7, _ := repository.ArticleGetWithWriter(9002)
	article8, _ := repository.ArticleGetWithWriter(9003)
	article9, _ := repository.ArticleGetWithWriter(9004)
	article10, _ := repository.ArticleGetWithWriter(9005)

	writer1, _ := repository.WriterGetByID(1001)
	writer2, _ := repository.WriterGetByID(1002)

	article11, _ := repository.ArticleGetWithTags(9001)
	article12, _ := repository.ArticleGetWithTags(9004)

	articles, _ := repository.ArticleListWithTags()

	data := map[string]interface{}{
		"Article1":  article1,
		"Article2":  article2,
		"Article3":  article3,
		"Article4":  article4,
		"Article5":  article5,
		"Article6":  article6,
		"Article7":  article7,
		"Article8":  article8,
		"Article9":  article9,
		"Article10": article10,
		"Writer1":   writer1,
		"Writer2":   writer2,
		"Article11": article11,
		"Article12": article12,
		"Articles":  articles,
	}

	// JSON 形式でレスポンスを返却します。
	return c.JSONPretty(http.StatusOK, data, "  ")
}
