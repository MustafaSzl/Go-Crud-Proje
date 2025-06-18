package main

import (
	"database/sql"
)

// getUsers fonksiyonu users tablosundaki tüm kullanıcıları veritabanından çeker
// db parametresi veritabanı bağlantısıdır
// dönen değerler: kullanıcı listesi ve hata (varsa)
func getUsers(db *sql.DB) ([]User, error) {
	// SQL sorgusunu çalıştırıyorum, tüm kullanıcıları id sırasına göre alıyorum
	rows, err := db.Query("SELECT id, username, email, password FROM users ORDER BY id")
	if err != nil {
		// sorgu hata verirse buraya gelir, hata döndürürüm
		return nil, err
	}
	// sorgudan dönen sonuçları sonunda kapatmak için defer kullanıyorum
	defer rows.Close()

	// boş bi liste oluşturdum kullanıcıları içine atacağım
	var users []User

	// rows içindeki her satıra yani her kullanıcıya tek tek bakıyorum
	for rows.Next() {
		// her satır için boş bi User yapısı açıyorum
		var u User

		// satırdaki verileri User içindeki alanlara yerleştiriyorum
		err := rows.Scan(&u.ID, &u.Username, &u.Email, &u.Password)
		if err != nil {
			// verileri çekerken bi hata olursa buraya gelir, fonksiyonu hata ile bitiririm
			return nil, err
		}

		// başarılıysa o kullanıcıyı listeye ekliyorum
		users = append(users, u)
	}

	// tüm kullanıcılar listeye eklendi, listeyi döndürüyorum, hata yok (nil)
	return users, nil
}
