package queries

import (
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"koreatech-board-api/cmd/db"
	"koreatech-board-api/cmd/enums"
	"koreatech-board-api/cmd/model"
	"net/http"
)

// @Summary		Get minumum notice list
// @Description	Get minumum notice list for board widget, only 5 new notices
// @Tags			Notice
// @Accept			json
// @Produce		json
// @Param   department  path     string     true  "name of the department"       Enums(arch, cse, dorm, mse, ace, ide, ite, mechanical, mechatronics, school, sim)
// @Param			board			path		string	true	"name of the board" Enums(notice, free, job, pds, lecture, bachelor, scholar)
// @Success		200				{object}	model.APIData
// @Failure		400
// @Failure		404
// @Router			/{department}/{board}/widget [get]
func GetMinimumNotices(c echo.Context) error {
	departmentRaw := c.Param("department")
	boardRaw := c.Param("board")

	apiError := ""
	status := http.StatusOK

	department, ok := enums.ParseDepartment(departmentRaw)

	if !ok {
		status = http.StatusNotFound
		apiError = fmt.Sprintf("Department \"%s\" not found!", departmentRaw)
	}

	board, ok := enums.ParseBoard(boardRaw)

	if !ok {
		status = http.StatusNotFound
		apiError = fmt.Sprintf("Board \"%s\" not found!", boardRaw)
	}

	var results []model.Board = nil

	if department != enums.UNKNOWN_DEPARTMENT && board != enums.UNKNOWN_BOARD {
		query := `
			SELECT 
				id::text, num, title, writer, to_char(write_date, 'YYYY-MM-DD') as write_date, 
				read_count, (init_crawled_time = update_crawled_time) as is_new, is_notice
			FROM notice
			WHERE department = $1 AND board = $2
			ORDER BY write_date DESC, num DESC
			LIMIT 5
		`

		rows, err := db.Pool.Query(c.Request().Context(), query, department.String(), board.String())
		if err != nil {
			status = http.StatusBadRequest
			apiError = "Query error: " + err.Error()
		} else {
			results, err = pgx.CollectRows(rows, pgx.RowToStructByName[model.Board])
			if err != nil {
				status = http.StatusBadRequest
				apiError = "Scan error: " + err.Error()
			}
		}
	}

	if results != nil {
		for i := range results {
			results[i].Num = 5 - int64(i)
		}
	} else {
		results = []model.Board{}
	}

	apiData := model.APIData{
		StatusCode: status,
		LastPage:   1,
		Error:      apiError,
		Posts:      results,
	}

	return c.JSON(status, apiData)
}
