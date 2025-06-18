package main

// Bu dosya main paketi içinde yer alır, diğer Go dosyalarıyla birlikte çalışabilir

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)


// PostgreSQL veritabanına bağlanmayı sağlayan Connect fonksiyonu
func Connect() *sql.DB {

	conStr := "user=postgres password=1235 dbname=go_crud_db sslmode=disable"

	// Veritabanı bağlantı bilgileri: kullanıcı adı, şifre, veritabanı adı ve SSL ayarı

	db, err := sql.Open("postgres", conStr)

	// sql.Open: PostgreSQL'e bağlanmak için bağlantıyı hazırlar, ancak hemen test etmez

	if err != nil {
		log.Fatal("VERİ TABANI BAĞLANTI HATASI", err)
	}

	// Hata varsa, log.Fatal ile hata mesajı yazdırılır ve program durdurulur

	return db
}

// Hata yoksa, bağlantı nesnesi (db) geri döndürülür ve program bu bağlant
