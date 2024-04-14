package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func SplitName(name string) (string, string) {
	split := strings.Split(name, ".")
	return split[0], split[1]
}
func GetName(key string) {
	switch key {
	case "add":
		{
			fmt.Print("Filename: ")
			fmt.Scan(&info.Fullname)
			info.Filename, info.Extension = SplitName(info.Fullname)
		}
	case "search":
		{
			fmt.Print("File to search: ")
			fmt.Scan(&search.fullNotation)
			search.name, search.ext = SplitName(search.fullNotation)
		}
	case "write":
		{
			fmt.Print("File to update: ")
			fmt.Scan(&search.fullNotation)
			update.name, update.ext = SplitName(update.fullNotation)
		}
	case "file create":
		{
			fmt.Print("File to create: ")
			fmt.Scan(&createFile.fullNotation)
			createFile.name, createFile.ext = SplitName(createFile.fullNotation)
		}
	default:
		{
			log.Fatal("Unknown key word")
		}
	}
}
func GetData() error {
	fmt.Print("Data to read \033[32m(Press Ctrl+Z to end reading)\033[0m: ")
	info.Data, err = reader.ReadBytes('\x04')
	if err == io.EOF {
		bytes.TrimSpace(info.Data)
		return nil
	}
	return err
}
func DbName() string {
	fmt.Print("Database name: ")
	var dbName string
	fmt.Scan(&dbName)
	return dbName
}
func DirName() string {
	fmt.Print("Direction name: ")
	var dir string
	fmt.Scan(&dir)
	return dir
}
func DirCreate() (string, error) {
	dir := DirName()
	name := DbName()
	err := os.Mkdir("./"+dir, 0755)
	dbPath := filepath.Join(dir, name)
	return dbPath, err
}
func Find(db *sql.DB, fln, ext string) (*sql.Rows, error) {
	query := Form()
	rows, err := db.Query(query, fln, ext)
	if !rows.Next() {
		fmt.Println("No rows found")
		return nil, err
	}
	return rows, err
}
func Form() string {
	return `SELECT * FROM files WHERE filename = ? AND extension = ? `
}

func Write(db *sql.DB) error {
	_, err = db.Exec(`UPDATE files SET data = ?, status = ? WHERE filename = ?`, info.Data, Statuses[1], update.name)
	return err
}
func RunFile(fileToRun string) error {
	return err
}
