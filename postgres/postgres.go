package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func ConnectPostgres() error {
	conn := "postgresql://neondb_owner:b7FsSlEq4PIV@ep-black-poetry-a1zc2yjj.ap-southeast-1.aws.neon.tech/neondb?sslmode=require"
	fmt.Println("connection:", conn)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return err
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		return err
	}

	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	log.Println("version:", version)

	rows, err := db.Query("SELECT id, name, value FROM playing_with_neon")
	if err != nil {
		return err
	}
	defer rows.Close()

	var ags []Neon
	for rows.Next() {
		var data Neon
		rows.Scan(&data.ID, &data.Name, &data.Value)

		ags = append(ags, data)
	}

	for _, a := range ags {
		log.Println("data:", a.ID, a.Name, a.Value)
	}

	var value float32
	if err := db.QueryRow("SELECT value FROM playing_with_neon WHERE name = $1", "c81e728d9d").Scan(&value); err != nil {
		return err
	}
	log.Println("single value:", value)

	return nil
}

type Neon struct {
	ID    int
	Name  string
	Value float32
}
