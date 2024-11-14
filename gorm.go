package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Album struct {
	ID     int64 `gorm:"primaryKey"`
	Title  string
	Artist string
	Price  float32
}

var db *gorm.DB

func main() {
	// Inisialisasi koneksi database dengan GORM
	dbUser := "kuliah"
	dbPass := "kuliah"
	dbName := "go-web"
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=1500 sslmode=disable", dbUser, dbPass, dbName)

	// Membuka koneksi ke database dengan GORM
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal terhubung ke database:", err)
	}

	// Auto migrasi untuk membuat tabel sesuai dengan struct Album
	db.AutoMigrate(&Album{})

	// Contoh penggunaan fungsi CRUD
	addAlbum := Album{Title: "The Modern Sound of Betty Carter", Artist: "Betty Carter", Price: 49.99}
	if err := createAlbum(addAlbum); err != nil {
		log.Fatal(err)
	}

	albums, err := getAlbumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)

	alb, err := getAlbumByID(2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Album found: %v\n", alb)
}

func createAlbum(album Album) error {
	result := db.Create(&album)
	if result.Error != nil {
		return fmt.Errorf("createAlbum: %v", result.Error)
	}
	fmt.Printf("Album created with ID: %v\n", album.ID)
	return nil
}

func getAlbumsByArtist(artist string) ([]Album, error) {
	var albums []Album
	result := db.Where("artist = ?", artist).Find(&albums)
	fmt.Printf("Query executed, found %d albums\n", result.RowsAffected) // Debug log
	if result.Error != nil {
		return nil, fmt.Errorf("getAlbumsByArtist: %v", result.Error)
	}
	return albums, nil
}

func getAlbumByID(id int64) (Album, error) {
	var album Album
	result := db.First(&album, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return album, fmt.Errorf("getAlbumByID %d: no such album", id)
		}
		return album, fmt.Errorf("getAlbumByID: %v", result.Error)
	}
	return album, nil
}
