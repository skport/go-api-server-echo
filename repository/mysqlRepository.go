package repository

import (
	"log"
	"fmt"
	"os"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"skport/go-api-server-echo/domains"
)

// -----
// Manage Database
var Db *sql.DB

func dbOpen() (*sql.DB, error) {
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@%s(%s:%s)/%s",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_PROTCOL"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
		),
	)
	if err != nil {
		return db, err
	}

	if err := db.Ping(); err != nil {
		return db, err
	}

	return db, nil
}

// -----
// Realization Class
type MySQLRepository struct{
	albums []domains.Album
}

func NewMySQLRepository() Repository {
	r := new(MySQLRepository)
	r.init()
	return r
}

func (rp *MySQLRepository) init() {
	rp.albums = []domains.Album{
		{ID: 1, Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: 2, Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}

	var err error
	Db, err = dbOpen()
	if err != nil {
		log.Fatal(err)
	}
}

func (rp *MySQLRepository) ReadAll() ([]domains.Album, error) {
	albums := []domains.Album{}

	rows, err := Db.Query("SELECT * FROM album ORDER BY id")
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

func (rp *MySQLRepository) ReadById(id int) (domains.Album, error) {
	album := domains.Album{}

	err := Db.QueryRow("SELECT * FROM album WHERE id = ?", id).
		Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
	switch {
	case err == sql.ErrNoRows:
	case err != nil:
		return domains.Album{}, err
	}

	return album, nil
}

func (rp *MySQLRepository) Post(newAlbum domains.Album) (error) {
	tx, err := Db.Begin()
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
