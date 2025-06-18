package main

import (
	"database/sql"
	"fmt"
)

// updateUser fonksiyonu:
// Verilen id'ye sahip kullanıcıyı günceller.
// Yeni kullanıcı adı, email ve şifreyi parametre olarak alır.
// Hata varsa döner.
func updateUser(db *sql.DB, id int, username, email, password string) error {
	// UPDATE sorgusu ile kullanıcı bilgilerini güncelliyoruz
	res, err := db.Exec("UPDATE users SET username=$1, email=$2, password=$3 WHERE id=$4",
		username, email, password, id)
	if err != nil {
		// Sorgu çalıştırılamadıysa hata döndür
		return err
	}

	// Sorgudan kaç satır etkilendi öğreniliyor
	affected, err := res.RowsAffected()
	if err != nil {
		// RowsAffected çağrılırken hata olursa hata döndür
		return err
	}

	// Eğer etkilenen satır sayısı 0 ise, bu id'ye sahip kullanıcı yok demektir
	if affected == 0 {
		return fmt.Errorf("güncelleme yapılmadı, id bulunamadı: %d", id)
	}

	// Başarılı güncelleme
	return nil
}
