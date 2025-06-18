package main

import "database/sql"

// createUser fonksiyonu, verilen bilgileri kullanarak yeni bir kullanıcı ekler
func createUser(db *sql.DB, username, email, password string) error {
	// SQL INSERT sorgusunu çalıştırır
	_, err := db.Exec(
		"INSERT INTO users (username, email, password) VALUES ($1, $2, $3)",
		username, email, password,
	)
	// Hata varsa döndürür, yoksa nil döner
	return err
}
