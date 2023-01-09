package services

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

// -----
// Variables
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

type httpResponce struct {
	Message string `json:"message"`
}

// -----
// Services
type Services struct {
}

func NewServices() *Services {
	s := new(Services)
	return s
}

func (s *Services) Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// getAlbums responds with the list of all albums as JSON.
func (s *Services) GetAlbums(c echo.Context) error {
	return c.JSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func (s *Services) PostAlbums(c echo.Context) error {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.Bind(newAlbum); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	return c.JSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func (s *Services) GetAlbumByID(c echo.Context) error {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			return c.JSON(http.StatusOK, a)
		}
	}

	r := &httpResponce{
		Message: "album not found",
	}
	return c.JSON(http.StatusNotFound, r)
}
