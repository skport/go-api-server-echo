package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"skport/go-api-server-echo/repository"
	"skport/go-api-server-echo/services"
)

func initConfig(e *echo.Echo) (*Handlers, *services.Services) {
	// Testing with in-memory data store
	rp := repository.NewInMemoryRepository()

	s := services.NewServices(&rp)
	h := NewHandler()
	h.Init(e, s)
	return h, s
}

func TestHealthCheck(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	
	h, _ := initConfig(e)

	// Assertions
	if assert.NoError(t, h.healthCheck(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "OK", rec.Body.String())
	}
}

func TestHello(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	
	_, s := initConfig(e)

	// Assertions
	if assert.NoError(t, s.Hello(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "Hello, World!", rec.Body.String())
	}
}

func TestGetAlbums(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/albums")
	
	_, s := initConfig(e)

	expectedJson := `[{"id":1,"title":"Blue Train","artist":"John Coltrane","price":56.99},{"id":2,"title":"Jeru","artist":"Gerry Mulligan","price":17.99},{"id":3,"title":"Sarah Vaughan and Clifford Brown","artist":"Sarah Vaughan","price":39.99}]`

	// Assertions
	if assert.NoError(t, s.GetAlbums(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expectedJson, strings.ReplaceAll(rec.Body.String(), "\n", ""))
	}
}

func TestGetAlbumByID(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/albums/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	
	_, s := initConfig(e)

	expectedJson := `{"id":1,"title":"Blue Train","artist":"John Coltrane","price":56.99}`

	// Assertions
	if assert.NoError(t, s.GetAlbumByID(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expectedJson, strings.ReplaceAll(rec.Body.String(), "\n", ""))
	}
}

func TestGetAlbumByIDOutOfRange(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/albums/:id")
	c.SetParamNames("id")
	c.SetParamValues("4")
	
	_, s := initConfig(e)

	// Assertions
	if assert.NoError(t, s.GetAlbumByID(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
	}
}

func TestPostAlbums(t *testing.T) {
	inputJson := `{"id":4, "title":"Sun", "artist":"Apple", "price":10.12}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(inputJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/albums")
	
	_, s := initConfig(e)

	// Assertions
	if assert.NoError(t, s.PostAlbums(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, `"Accepted"`, strings.ReplaceAll(rec.Body.String(), "\n", ""))
	}
}