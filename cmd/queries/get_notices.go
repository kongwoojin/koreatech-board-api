package queries

import (
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"koreatech-board-api/cmd/db"
	"koreatech-board-api/cmd/enums"
	"koreatech-board-api/cmd/model"
	"math"
	"net/http"
	"strconv"
)

// @Summary		Get notice list
// @Description	Get notice list
// @Tags			Notice
// @Accept			json
// @Produce		json
// @Param   department  path     string     true  "name of the department"       Enums(arch, cse, dorm, mse, ace, ide, ite, mechanical, mechatronics, school, sim)
// @Param			board			path		string	true	"name of the board" Enums(notice, free, job, pds, lecture, bachelor, scholar)
// @Param			page			query		integer	false	"page of board"
// @Param			num_of_items	query		integer	false	"items per page"
// @Success		200				{object}	model.APIData
// @Failure		400
// @Failure		404
// @Router			/{department}/{board} [get]
func GetNotices(c echo.Context) error {
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

	page, pageErr := strconv.Atoi(c.QueryParam("page"))
	numOfItems, noiErr := strconv.Atoi(c.QueryParam("num_of_items"))

	if pageErr != nil {
		page = 1
	}

	if noiErr != nil {
		numOfItems = 20
	}

	var results []model.Board = nil
	var count int64
	lastPage := 1

	if department != enums.UNKNOWN_DEPARTMENT && board != enums.UNKNOWN_BOARD {
		offset := int64((page - 1) * numOfItems)

		query := `
			SELECT 
				id::text, num, title, writer, to_char(write_date, 'YYYY-MM-DD') as write_date, 
				read_count, (init_crawled_time = update_crawled_time) as is_new, is_notice
			FROM notice
			WHERE department = $1 AND board = $2
			ORDER BY is_notice DESC, write_date DESC, num DESC
			OFFSET $3 LIMIT $4
		`

		rows, err := db.Pool.Query(c.Request().Context(), query, department.String(), board.String(), offset, numOfItems)
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

		if status == http.StatusOK {
			countQuery := `SELECT count(*) FROM notice WHERE department = $1 AND board = $2`
			err = db.Pool.QueryRow(c.Request().Context(), countQuery, department.String(), board.String()).Scan(&count)
			if err != nil {
				status = http.StatusBadRequest
				apiError = "Count query error: " + err.Error()
			}
			lastPage = int(math.Ceil(float64(count) / float64(numOfItems)))
		}
	}

	if results != nil {
		for i := range results {
			results[i].Num = count - int64(page-1)*int64(numOfItems) - int64(i)
		}
	} else {
		results = []model.Board{} // Ensure empty array not null
	}

	apiData := model.APIData{
		StatusCode: status,
		LastPage:   lastPage,
		Error:      apiError,
		Posts:      results,
	}

	return c.JSON(status, apiData)
}
