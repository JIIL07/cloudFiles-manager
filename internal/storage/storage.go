package storage

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"

	"github.com/JIIL07/cloudFiles-manager/internal/config"
)

type Storage struct {
	DB *sqlx.DB
}

type UserData struct {
	UserID   int    `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
	Protocol string `db:"hashprotocol" json:"hashprotocol"`
	Admin    int    `db:"admin" json:"admin"`
}
type File struct {
	Id        int    `db:"id" json:"id"`
	UserID    int    `db:"user_id" json:"user_id"`
	Filename  string `db:"filename" json:"filename"`
	Extension string `db:"extension" json:"extension"`
	Filesize  int    `db:"filesize" json:"filesize"`
	Status    string `db:"status" json:"status"`
	Data      []byte `db:"data" json:"data"`
}

func InitDatabase(config *config.Config) (*Storage, error) {
	if config.Env == "prod" || config.Env == "debug" {
		config.Database.DataSourceName = os.Getenv("DATABASE_PATH")
	}
	db, err := sqlx.Open(config.Database.DriverName, config.Database.DataSourceName)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS users (
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,
		"username" TEXT NOT NULL UNIQUE,
		"email" TEXT NOT NULL,
		"password" TEXT NOT NULL,
		"hashprotocol" TEXT,
		"admin" INTEGER DEFAULT 1
	);`)

	if err != nil {
		return nil, fmt.Errorf("failed to create table: %v", err)
	}

	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS files (
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,
		"user_id" INTEGER NOT NULL,
		"filename" TEXT, 
		"extension" TEXT, 
		"filesize" INTEGER, 
		"status" TEXT, 
		"data" BLOB,
		FOREIGN KEY (user_id) REFERENCES users(id)
	);`)

	if err != nil {
		return nil, fmt.Errorf("failed to create table: %v", err)
	}

	return &Storage{DB: db}, nil
}

func (s *Storage) CloseDatabase() error {
	return s.DB.Close()
}

func (s *Storage) SaveNewUser(u *UserData) error {
	_, err := s.DB.Exec(`INSERT INTO users 
		(username, email, password, hashprotocol, admin) VALUES (?, ?, ?, ?, ?)`,
		u.Username, u.Email, u.Password, u.Protocol, u.Admin,
	)
	if err != nil {
		return fmt.Errorf("failed to save new user: %v", err)
	}
	return nil
}

func (s *Storage) GetAllUsers() ([]UserData, error) {
	var users []UserData
	var u = &UserData{}
	rows, err := s.DB.Query(`SELECT * FROM users`)
	if err != nil {
		return nil, fmt.Errorf("failed to get all users: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&u.Username, &u.Email, &u.Password, &u.Protocol, &u.Admin); err != nil {
			return nil, fmt.Errorf("failed to scan user: %v", err)
		}
		users = append(users, *u)
	}
	return users, nil
}

func (s *Storage) DeleteUser(username string) error {
	_, err := s.DB.Exec(`DELETE FROM users WHERE username = ?`, username)
	if err != nil {
		return err
	}
	return nil
}
