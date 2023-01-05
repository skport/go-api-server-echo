package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type httpResponce struct {
	Message string `json:"message"`
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.GET("/albums", getAlbums)
	e.GET("/albums/:id", getAlbumByID)
	e.POST("/albums", postAlbums)

	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}

// -------
// Handler

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c echo.Context) error {
	return c.JSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c echo.Context) error {
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
func getAlbumByID(c echo.Context) error {
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
