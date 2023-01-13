package repository

import (
	"skport/go-api-server-echo/domains"
)

// interface
type Repository interface {
	ReadAll() ([]domains.Album, error)
	ReadById(id int) (domains.Album, error)
	Post(newAlbum domains.Album) (error)
}