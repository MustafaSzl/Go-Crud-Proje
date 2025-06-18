package main

import (
	"fmt"
	"log"
)

func main() {
	// 1. Veritabanı bağlantısı oluşturuluyor
	db := Connect()
	defer db.Close() // Program sonunda bağlantıyı kapat

	fmt.Println("Veritabanına bağlanıldı.")

	// 2. Yeni kullanıcı ekle
	err := createUser(db, "Mehmet", "Mehmet@gmail.com", "12345")
	if err != nil {
		log.Println("Kullanıcı ekleme hatası:", err)
	} else {
		fmt.Println("Yeni kullanıcı eklendi: ali")
	}

	// 3. Kullanıcıları veritabanından çek ve yazdır
	users, err := getUsers(db)
	if err != nil {
		log.Fatal("Kullanıcıları çekme hatası:", err)
	}
	fmt.Println("Mevcut kullanıcılar:")
	for _, u := range users {
		fmt.Printf("ID: %d, Kullanıcı Adı: %s, Email: %s\n", u.ID, u.Username, u.Email)
	}

	if len(users) == 0 {
		fmt.Println("Kullanıcı yok, program sonlandırılıyor.")
		return
	}

	// 4. Kullanıcının bilgilerini güncelle
	idGuncelle := 0
	err = updateUser(db, idGuncelle, "yeni_kullanici_adi", "yeni_email@example.com", "yeni_sifre")
	if err != nil {
		log.Println("Güncelleme hatası:", err)
	} else {
		fmt.Println("Kullanıcı güncellendi:", idGuncelle)
	}

	// 5. Güncellenmiş kullanıcıları tekrar çek ve yazdır
	users, _ = getUsers(db)
	fmt.Println("Güncellenmiş kullanıcılar:")
	for _, u := range users {
		fmt.Printf("ID: %d, Kullanıcı Adı: %s, Email: %s\n", u.ID, u.Username, u.Email)
	}

	// 6. Kullanıcıyı sil
	idSil := 0 // Silmek istediğin kullanıcının ID'si
	err = deleteUser(db, idSil)
	if err != nil {
		log.Println("Silme hatası:", err)
	} else {
		fmt.Println("Kullanıcı silindi:", idSil)
	}

	// 7. Silme sonrası kullanıcıları tekrar çek ve yazdır
	users, _ = getUsers(db)
	fmt.Println("Silme sonrası kullanıcılar:")
	for _, u := range users {
		fmt.Printf("ID: %d, Kullanıcı Adı: %s, Email: %s\n", u.ID, u.Username, u.Email)
	}
}
