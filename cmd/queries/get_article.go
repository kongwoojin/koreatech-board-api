package queries

import (
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"koreatech-board-api/cmd/db"
	"koreatech-board-api/cmd/model"
	"net/http"
)

// @Summary		Get article
// @Description	Get article by UUID
// @Tags			Notice
// @Accept			json
// @Produce		json
// @Param			uuid	query		string	true	"uuid of article"
// @Success		200		{object}	model.ApiArticle
// @Failure		400
// @Failure		404
// @Router			/article [get]
func GetArticle(c echo.Context) error {
	uuid := c.QueryParam("uuid")
	status := http.StatusOK
	apiError := ""

	var article model.Article

	query := `
		SELECT 
			id::text, num, title, writer, to_char(write_date, 'YYYY-MM-DD') as write_date, 
			article_url, content, is_notice
		FROM notice
		WHERE id = $1::uuid
	`

	rows, err := db.Pool.Query(c.Request().Context(), query, uuid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ApiArticle{
			StatusCode: http.StatusBadRequest,
			Error:      "Query error: " + err.Error(),
		})
	}

	collected, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.Article])
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ApiArticle{
			StatusCode: http.StatusBadRequest,
			Error:      "Scan error: " + err.Error(),
		})
	}

	if len(collected) == 0 {
		return c.JSON(http.StatusOK, model.ApiArticle{
			StatusCode: http.StatusOK,
			Error:      "",
		})
	}

	article = collected[0]

	fileQuery := `
		SELECT f.file_name, f.file_url
		FROM file f
		JOIN notice_files nf ON f.id = nf.file_id
		WHERE nf.notice_id = $1::uuid
	`

	fileRows, err := db.Pool.Query(c.Request().Context(), fileQuery, uuid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ApiArticle{
			StatusCode: http.StatusBadRequest,
			Error:      "File Query error: " + err.Error(),
		})
	}

	files, err := pgx.CollectRows(fileRows, pgx.RowToStructByName[model.Files])
	if err != nil {
		// Log error but proceed
	}
	article.Files = files

	return c.JSON(status, model.ApiArticle{
		StatusCode: status,
		Error:      apiError,
		Num:        article.Num,
		Id:         article.Id,
		Title:      article.Title,
		Writer:     article.Writer,
		WriteDate:  article.WriteDate,
		ArticleUrl: article.ArticleUrl,
		Content:    article.Content,
		IsNotice:   article.IsNotice,
		Files:      article.Files,
	})
}
