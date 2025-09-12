package recordings

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
)

type Album struct {
	Id     string
	Title  string
	Artist string
	Price  float64
}

var db *sql.DB

func InitDb() error {
	var dbuserName string = os.Getenv("DBUSER")
	if len(dbuserName) == 0 {
		return fmt.Errorf("dbusername pas renseigné. il faut remplir la variable d'environnement %s", "DBUSER")
	}

	var dbuserPass string = os.Getenv("DBPASS")
	if len(dbuserPass) == 0 {
		return fmt.Errorf("dbuserpass pas renseigné. Il faut remplir la variable d'environnement %s", "DBPASS")
	}

	// Capture connection properties.
	cfg := mysql.NewConfig()
	cfg.User = dbuserName
	cfg.Passwd = dbuserPass
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "recordings"

	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return err
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return pingErr
	}
	fmt.Println("Connected!")

	return nil
}

func GetAlbums() ([]Album, error) {

	// if true {
	// 	return nil, fmt.Errorf("lplluff")
	// }

	albums := make([]Album, 0)

	rows, err := db.Query("SELECT * FROM album")
	if err != nil {
		return nil, fmt.Errorf("albums: %v", err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.Id, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albums: %v", err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albums: %v", err)
	}

	return albums, nil
}

func GetAlbum(id string) (Album, error) {

	album := Album{
		Id:     id,
		Title:  id,
		Artist: id,
	}
	return album, nil
}
