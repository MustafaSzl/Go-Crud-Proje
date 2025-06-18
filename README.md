# Go-Crud-Proje

Bu proje, **Go (Golang)** programlama dili kullanılarak yazılmış basit bir **CRUD (Create, Read, Update, Delete)** uygulamasıdır. Backend tarafında PostgreSQL veritabanı kullanılmaktadır.

## 🔧 Özellikler

- Kullanıcı ekleme
- Kullanıcı listeleme
- Kullanıcı güncelleme
- Kullanıcı silme

## 🗃 Veritabanı Bilgisi

Bu projeyle birlikte bir `.sql` dosyası sağlanmamaktadır. PostgreSQL veritabanını **kendiniz oluşturmanız gerekmektedir**.

### Veritabanı Oluşturma Önerisi

Aşağıdaki örneğe benzer bir tablo yapısı oluşturabilirsiniz:

```sql
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  username VARCHAR(100),
  email VARCHAR(100),
  password VARCHAR(100)
);


connStr := "user=postgres password=12345 dbname=go_crud_db sslmode=disable"

