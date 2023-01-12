package services

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"skport/go-api-server-echo/domains"
	"skport/go-api-server-echo/repository"
)

// -----
// Variables

type httpResponce struct {
	Message string `json:"message"`
}

// -----
// Services
type Services struct {
	rp repository.Repository
}

func NewServices(rp *repository.Repository) *Services {
	s := new(Services)
	s.rp = *rp
	return s
}

// Hello godoc
// @Summary      Show Hello
// @Description  Show Hello
// @Tags         Basic
// @Accept       */*
// @Produce      plain
// @Success      200 {string} string "Hello, World!"
// @Router       / [get]
func (s *Services) Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// GetAlbums godoc
// @Summary      Show Albums
// @Description  getAlbums responds with the list of all albums as JSON.
// @Tags         Album
// @Accept       */*
// @Produce      json
// @Success      200 {string} string "GetAlbums"
// @Router       /albums [get]
func (s *Services) GetAlbums(c echo.Context) error {
	r, err := s.rp.ReadAll()
	if err != nil {
		log.Println(err)
	}
	return c.JSON(http.StatusOK, r)
}

// PostAlbums godoc
// @Summary      Post Album
// @Description  postAlbums adds an album from JSON received in the request body.
// @Tags         Album
// @Accept       json
// @Produce      json
// @Success      200 {string} string "Post album"
// @Router       /albums [post]
func (s *Services) PostAlbums(c echo.Context) error {
	var newAlbum domains.Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.Bind(newAlbum); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Add the new album to data store.
	err := s.rp.Post(newAlbum)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, newAlbum)
}

// GetAlbumByID godoc
// @Summary      Show Album for id
// @Description  getAlbumByID locates the album whose ID value matches the id parameter
// @Tags         Album
// @Accept       */*
// @Produce      json
// @Success      200 {string} string "Album for id"
// @Router       /albums/{id} [get]
// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func (s *Services) GetAlbumByID(c echo.Context) error {
	id := c.Param("id")

	// Looking for an album whose ID value match in data store.
	i, _ := strconv.Atoi(id)
	r, err := s.rp.ReadById(i)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, r)

	/*
	// Not Found
	re := &httpResponce{
		Message: "album not found",
	}
	return c.JSON(http.StatusNotFound, re)
	*/
}
