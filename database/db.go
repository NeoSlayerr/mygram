package database

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
)

var (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "postgres"
	dbname   = "mygram"
	dialect  = "postgres"
)

var (
	db  *sql.DB
	err error
)

func handleDatabaseConnection() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err = sql.Open(dialect, psqlInfo)

	if err != nil {
		log.Panic("error occured while trying to validate database arguments:", err)
	}

	err = db.Ping()

	if err != nil {
		log.Panic("error occured while trying to connect to database:", err)
	}

}

func handleCreateRequiredTables() {
	userTable := `
		CREATE TABLE IF NOT EXISTS "users" (
			id SERIAL PRIMARY KEY,
			username VARCHAR(255) UNIQUE NOT NULL,
			email VARCHAR(255) UNIQUE NOT NULL,
			password TEXT NOT NULL,
			age int NOT NULL,
			created_at timestamptz DEFAULT now(),
			updated_at timestamptz DEFAULT now()
		);
	`
	
	photoTable := `
		CREATE TABLE IF NOT EXISTS "photos" (
			id SERIAL PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			caption TEXT NOT NULL,
			photo_url TEXT NOT NULL,
			user_id int NOT NULL,
			created_at timestamptz DEFAULT now(),
			updated_at timestamptz DEFAULT now(),
			CONSTRAINT photos_user_id_fk
				FOREIGN KEY(user_id)
					REFERENCES users(id)
						ON DELETE CASCADE
		);
	`

	commentTable := `
		CREATE TABLE IF NOT EXISTS "comments" (
			id SERIAL PRIMARY KEY,
			user_id int NOT NULL,
			photo_id int NOT NULL,
			message TEXT NOT NULL,
			created_at timestamptz DEFAULT now(),
			updated_at timestamptz DEFAULT now(),
			CONSTRAINT comments_user_id_fk
				FOREIGN KEY(user_id)
					REFERENCES users(id)
						ON DELETE CASCADE,
			CONSTRAINT comments_photo_id_fk
				FOREIGN KEY(photo_id)
					REFERENCES photos(id)
						ON DELETE CASCADE
		);
	`

	socialmediaTable := `
		CREATE TABLE IF NOT EXISTS "socialmedias" (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			social_media_url TEXT NOT NULL,
			user_id int NOT NULL,
			created_at timestamptz DEFAULT now(),
			updated_at timestamptz DEFAULT now(),
			CONSTRAINT socialmedias_user_id_fk
				FOREIGN KEY(user_id)
					REFERENCES users(id)
						ON DELETE CASCADE
		);
	`

	createTableQueries := fmt.Sprintf("%s %s %s %s", userTable, photoTable, commentTable, socialmediaTable)

	_, err = db.Exec(createTableQueries)

	if err != nil {
		log.Panic("error occured while trying to create required tables:", err)
	}
}

func InitiliazeDatabase() {
	handleDatabaseConnection()
	handleCreateRequiredTables()
}

func GetDatabaseInstance() *sql.DB {
	return db
}
