package repository

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"skport/go-api-server-echo/domains"
)

// Realization Class
type DevDBRepository struct{
	albums []domains.Album
}

func NewDevDBRepository() Repository {
	r := new(DevDBRepository)
	r.init()
	return r
}

func (rp *DevDBRepository) init() {
	rp.albums = []domains.Album{
		{ID: 1, Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: 2, Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}
}

func (rp *DevDBRepository) ReadAll() ([]domains.Album, error) {
	albums := []domains.Album{}

	db, err := sql.Open("mysql", "root:pass@tcp(127.0.0.1:3306)/api")
	if err != nil {
		return albums, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM album ORDER BY id")
	if err != nil {
		return albums, err
	}
	defer rows.Close()

	for rows.Next() {
		var a domains.Album
		rows.Scan(&a.ID, &a.Title, &a.Artist, &a.Price)
		albums = append(albums, a)
	}

	return albums, nil
}

func (rp *DevDBRepository) ReadById(id int) (domains.Album, error) {
	album := domains.Album{}

	db, err := sql.Open("mysql", "root:pass@tcp(127.0.0.1:3306)/api")
	if err != nil {
		return album, err
	}
	defer db.Close()

	err = db.QueryRow("SELECT * FROM album WHERE id = ?", id).
					Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
	switch {
	case err == sql.ErrNoRows:
	case err != nil:
		return domains.Album{}, err
	}

	return album, nil
}

func (rp *DevDBRepository) Post(newAlbum domains.Album) (error) {
	db, err := sql.Open("mysql", "root:pass@tcp(127.0.0.1:3306)/api")
	if err != nil {
		return err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("INSERT INTO album VALUES(?, ?, ?, ?)", newAlbum.ID, newAlbum.Title, newAlbum.Artist, newAlbum.Price)
	if err != nil {
		return err
	}

	if err != nil {
		_ = tx.Rollback()
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
