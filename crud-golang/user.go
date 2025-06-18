package main

// Veritabanındaki 'users' tablosunun Go karşılığı olan yapı
type User struct { //DEĞİŞKEN OLUŞTURUYORUZ KISACA
	ID       int    // id SERIAL PRIMARY KEY — otomatik artan birincil anahtar
	Username string // username VARCHAR(100) — kullanıcının adı
	Email    string // email VARCHAR(100) — kullanıcının e-posta adresi
	Password string // password VARCHAR(100) — kullanıcının şifresi
}
