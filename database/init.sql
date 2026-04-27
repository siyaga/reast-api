-- Mengaktifkan ekstensi UUID jika dibutuhkan ke depannya
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- (Tabel akan dibuat oleh GORM, jadi tidak perlu CREATE TABLE di sini)