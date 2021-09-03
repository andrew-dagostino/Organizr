package tests

import (
	"net/http"
	"net/http/httptest"
	"organizr/server/models"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// Mocks
type handler struct {
	db []models.Board
}

func (h *handler) getBoards(c echo.Context) error {
	boards := h.db
	return c.JSON(http.StatusOK, boards)
}

func (h *handler) getBoard(c echo.Context) error {
	board_gid := c.Param("board_gid")

	var board models.Board
	for _, b := range h.db {
		if b.Gid == board_gid {
			board = b
		}
	}

	if (models.Board{}) == board {
		return echo.NewHTTPError(http.StatusNotFound, "Board Not Found")
	}
	return c.JSON(http.StatusOK, board)
}

// Test Data
var (
	mockDB = &[]models.Board{
		{Id: 1, Gid: "abc123", Title: "title", MemberCount: 3},
		{Id: 2, Gid: "def456", Title: "title2", MemberCount: 5},
	}
	boardsJSON = `[{"id":1,"gid":"abc123","title":"title","board_member_count":3},{"id":2,"gid":"def456","title":"title2","board_member_count":5}]`
	boardJSON  = `{"id":2,"gid":"def456","title":"title2","board_member_count":5}`
)

func Test_GetBoards(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/board")
	h := &handler{*mockDB}

	// Assertions
	if assert.NoError(t, h.getBoards(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, boardsJSON, strings.TrimSpace(rec.Body.String()))
	}
}

func Test_GetBoard(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/board/:board_gid")
	c.SetParamNames("board_gid")
	c.SetParamValues("def456")
	h := &handler{*mockDB}

	// Assertions
	if assert.NoError(t, h.getBoard(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, boardJSON, strings.TrimSpace(rec.Body.String()))
	}
}

func Test_GetBoard_NotFound(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/board/:board_gid")
	c.SetParamNames("board_gid")
	c.SetParamValues("dne")
	h := &handler{*mockDB}

	// Assertions
	assert.Error(t, h.getBoard(c))
}
