package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	_ "github.com/jackc/pgx/v5/stdlib" // Driver PostgreSQL
// )

// var db *sql.DB

// type Album struct {
// 	ID     int64
// 	Title  string
// 	Artist string
// 	Price  float32
// }

// func main() {
// 	// Mendapatkan variabel koneksi dari environment
// 	dbUser := "kuliah"
// 	dbPass := "kuliah"
// 	dbName := "go-web"

// 	// Format DSN untuk PostgreSQL
// 	dsn := fmt.Sprintf("postgresql://%s:%s@localhost:1500/%s", dbUser, dbPass, dbName)

// 	// Membuka koneksi ke database
// 	var err error
// 	db, err = sql.Open("pgx", dsn)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Tes koneksi
// 	pingErr := db.Ping()
// 	if pingErr != nil {
// 		log.Fatal(pingErr)
// 	}
// 	fmt.Println("Connected!")

// 	// Query contoh
// 	albums, err := albumsByArtist("John Coltrane")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("Albums found: %v\n", albums)

// 	alb, err := albumByID(2)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("Album found: %v\n", alb)

// 	albID, err := addAlbum(Album{
// 		Title:  "The Modern Sound of Betty Carter",
// 		Artist: "Betty Carter",
// 		Price:  49.99,
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("ID of added album: %v\n", albID)
// }

// // albumsByArtist queries for albums that have the specified artist name.
// func albumsByArtist(name string) ([]Album, error) {
// 	var albums []Album

// 	rows, err := db.Query("SELECT id, title, artist, price FROM album WHERE artist = $1", name)
// 	if err != nil {
// 		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var alb Album
// 		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
// 			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
// 		}
// 		albums = append(albums, alb)
// 	}
// 	if err := rows.Err(); err != nil {
// 		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
// 	}
// 	return albums, nil
// }

// // albumByID queries for the album with the specified ID.
// func albumByID(id int64) (Album, error) {
// 	var alb Album

// 	row := db.QueryRow("SELECT id, title, artist, price FROM album WHERE id = $1", id)
// 	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
// 		if err == sql.ErrNoRows {
// 			return alb, fmt.Errorf("albumByID %d: no such album", id)
// 		}
// 		return alb, fmt.Errorf("albumByID %d: %v", id, err)
// 	}
// 	return alb, nil
// }

// // addAlbum adds the specified album to the database,
// // returning the album ID of the new entry
// func addAlbum(alb Album) (int64, error) {
// 	var id int64
// 	query := "INSERT INTO album (title, artist, price) VALUES ($1, $2, $3) RETURNING id"
// 	err := db.QueryRow(query, alb.Title, alb.Artist, alb.Price).Scan(&id)
// 	if err != nil {
// 		return 0, fmt.Errorf("addAlbum: %v", err)
// 	}
// 	return id, nil
// }
