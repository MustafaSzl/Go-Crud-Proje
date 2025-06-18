package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type User struct {
	ID       int
	Username string
	Email    string
	Password string
}

// ----------------------------------------
// 📌 İşlem Kontrol Değişkenleri
// ----------------------------------------
var enableCreate = false // true ise kullanıcı ekleme aktif
var enableUpdate = false // true ise kullanıcı güncelleme aktif
var enableDelete = true  // true ise kullanıcı silme aktif

func main() {
	connStr := "user=postgres password=12345 dbname=go_crud_db sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Bağlantı hatası:", err)
	}
	defer db.Close()

	fmt.Println("PostgreSQL bağlantısı başarılı!")

	// ----------------------------------------
	// 1. Yeni kullanıcı ekle (aktif ise)
	// ----------------------------------------
	if enableCreate {
		err = createUser(db, "", "", "")
		if err != nil {
			log.Println("Ekleme hatası:", err)
		} else {
			fmt.Println("Yeni kullanıcı eklendi")
		}
	} else {
		fmt.Println("Kullanıcı ekleme işlemi devre dışı bırakıldı.")
	}

	// ----------------------------------------
	// 2. Kullanıcıları oku ve yazdır
	// ----------------------------------------
	users, err := getUsers(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nKullanıcılar:")
	for _, u := range users {
		fmt.Printf("%d: %s - %s\n", u.ID, u.Username, u.Email)
	}

	if len(users) == 0 {
		fmt.Println("Hiç kullanıcı yok, işlemler iptal")
		return
	}

	// ----------------------------------------
	// 3. İlk kullanıcının bilgilerini güncelle (aktif ise)
	// ----------------------------------------
	firstID := users[0].ID
	if enableUpdate {
		err = updateUser(db, firstID, "ali_guncel", "ali_guncel@example.com", "54321")
		if err != nil {
			log.Println("Güncelleme hatası:", err)
		} else {
			fmt.Printf("\nKullanıcı güncellendi: ID %d\n", firstID)
		}
	} else {
		fmt.Println("\nKullanıcı güncelleme işlemi devre dışı bırakıldı.")
	}

	// ----------------------------------------
	// 4. Güncellenmiş kullanıcıları tekrar oku ve yazdır
	// ----------------------------------------
	users, _ = getUsers(db)
	fmt.Println("\nGüncellenmiş kullanıcılar:")
	for _, u := range users {
		fmt.Printf("%d: %s - %s\n", u.ID, u.Username, u.Email)
	}

	// ----------------------------------------
	// 5. Kullanıcı silme işlemi (aktif ise)
	// ----------------------------------------
	deleteID := 5 // Silmek istediğin kullanıcının ID'si
	err = deleteUser(db, deleteID)
	if err != nil {
		log.Println("Silme hatası:", err)
	} else if enableDelete {
		fmt.Printf("\nKullanıcı silindi: ID %d\n", deleteID)
	}

	// ----------------------------------------
	// 6. Silme sonrası kullanıcıları tekrar oku ve yazdır
	// ----------------------------------------
	users, _ = getUsers(db)
	fmt.Println("\nSilme sonrası kullanıcılar:")
	for _, u := range users {
		fmt.Printf("%d: %s - %s\n", u.ID, u.Username, u.Email)
	}
}

// ----------------------------------------
// 📌 Kullanıcı Ekleme Fonksiyonu
// ----------------------------------------
func createUser(db *sql.DB, username, email, password string) error {
	_, err := db.Exec("INSERT INTO users (username, email, password) VALUES ($1, $2, $3)", username, email, password)
	return err
}

// ----------------------------------------
// 📌 Kullanıcıları Okuma Fonksiyonu
// ----------------------------------------
func getUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT id, username, email, password FROM users ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.Username, &u.Email, &u.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

// ----------------------------------------
// 📌 Kullanıcı Güncelleme Fonksiyonu
// ----------------------------------------
func updateUser(db *sql.DB, id int, username, email, password string) error {
	res, err := db.Exec("UPDATE users SET username=$1, email=$2, password=$3 WHERE id=$4", username, email, password, id)
	if err != nil {
		return err
	}
	affected, _ := res.RowsAffected()
	if affected == 0 {
		return fmt.Errorf("güncelleme yapılmadı, id bulunamadı: %d", id)
	}
	return nil
}

// ----------------------------------------
// 📌 Kullanıcı Silme Fonksiyonu
// ----------------------------------------
func deleteUser(db *sql.DB, id int) error {
	if !enableDelete {
		fmt.Println("Silme işlemi devre dışı bırakıldı, silme yapılmadı.")
		return nil
	}

	res, err := db.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return fmt.Errorf("silme yapılmadı, id bulunamadı: %d", id)
	}

	return nil
}
