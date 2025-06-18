package main

import (
	"database/sql"
	"fmt"
)

// deleteUser fonksiyonu, verilen id'ye sahip kullanıcıyı siler
func deleteUser(db *sql.DB, id int) error {
	// users tablosundan id'ye göre kayıt silinir
	res, err := db.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		return err // Hata varsa geri döndür
	}

	// Silinen satır sayısını al
	affected, err := res.RowsAffected()
	if err != nil {
		return err // Satır sayısı alınamazsa hata döndür
	}

	// Eğer hiç satır silinmediyse (id yoksa) hata ver
	if affected == 0 {
		return fmt.Errorf("silme yapılmadı, id bulunamadı: %d", id)
	}

	return nil // Başarılıysa nil döndür
}
