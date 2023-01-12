package repository

import (
	"errors"

	"skport/go-api-server-echo/domains"
)

// interface
type Repository interface {
	ReadAll() ([]domains.Album, error)
	ReadById(id int) (domains.Album, error)
	Post(newAlbum domains.Album) (error)
}

// A InMemoryRepository Class
type InMemoryRepository struct{
	albums []domains.Album
}

func NewInMemoryRepository() Repository {
	r := new(InMemoryRepository)
	r.init()
	return r
}

func (rp *InMemoryRepository) init() {
	rp.albums = []domains.Album{
		{ID: 1, Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: 2, Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}
}

func (rp *InMemoryRepository) ReadAll() ([]domains.Album, error) {
	return rp.albums, nil
}

func (rp *InMemoryRepository) ReadById(id int) (domains.Album, error) {
	for _, a := range rp.albums {
		if a.ID == id {
			return a, nil
		}
	}
	return domains.Album{}, errors.New("Not Found")
}

func (rp *InMemoryRepository) Post(newAlbum domains.Album) (error) {
	rp.albums = append(rp.albums, newAlbum)
	return nil
}