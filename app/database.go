package app

import (
	"caturilham05/product/helper"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func NewDB() *sql.DB {
	err := godotenv.Load()
	helper.PanicIfError(err)

	// Ambil variabel dari .env
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	// gunakan parseTime=true untuk mengkonversi kolom datetime ke tipe time.Time
	loc := "Asia%2FJakarta"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=%s", user, pass, host, port, name, loc)
	// fmt.Println(dsn)
	// db, err := sql.Open("mysql", "root@tcp(localhost:3306)/restapigov2")
	// helper.PanicIfError(err)
	db, err := sql.Open("mysql", dsn)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	return db
}
