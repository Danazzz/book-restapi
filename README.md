**BookRESTAPI**

BookRESTAPI adalah REST API untuk mengelola data buku dan kategorinya yang menggunakan framework **Gin** dan database  **PostgreSQL** . REST API ini mendukung fitur autentikasi JWT dan menyediakan endpoint untuk operasi CRUD buku, kategori, serta manajemen pengguna.

**Fitur**

1. **CRUD Books**

* Tambah buku
* Lihat semua buku
* Lihat detail buku berdasarkan ID
* Perbarui buku
* Hapus buku

2. **CRUD Categories**

* Tambah kategori
* Lihat semua kategori
* Lihat detail kategori berdasarkan ID
* Lihat buku berdasarkan kategori
* Perbarui buku
* Hapus kategori

3. **Authentication**

* Login (JWT)
* Registrasi user

4. **Validasi**

* Validasi tahun rilis buku (1980-2024)
* Validasi total halaman untuk menentukan ketebalan buku

**Teknologi yang Digunakan**

* **Golang** : Bahasa pemrograman utama.
* **Gin** : Framework web untuk REST API.
* **PostgreSQL** : Database relasional.
* **JWT** : Untuk autentikasi pengguna.
* **Railway** : Platform deployment.

**Struktur Folder**

.
├── controllers/            # Logika API untuk buku, kategori, dan pengguna
├── database/              # Koneksi database dan migrasi
├── middleware/          # Middleware untuk autentikasi JWT
├── models/                 # Model untuk buku, kategori, dan pengguna
├── routes/                  # Definisi rute API
├── config/                 # Konfigurasi aplikasi
├── main.go               # Entry point aplikasi
├── go.mod               # Dependency Go
├── .env                    # File konfigurasi variabel lingkungan

**Setup Proyek**

1. **Clone Repository**

`git clone https://github.com/Danazzz/book-restapi.git`

`cd book-restapi`

2. **Setup Local**

* Buat file **.env** atau gunakan **.env.example** dan tambahkan konfigurasi berikut:

```
DATABASE_URL=postgres://username:password@localhost:5432/your_database?sslmode=disablePORT=8080
JWT_SECRET=your_jwt_secret
```

3. **Instal Dependency**

`go mod tidy`

4. **Jalankan Aplikasi**

`go run main.go`

**Endpoint API**

**1. CRUD Books**

**Method**		**Endpoint**		**Deskripsi**

**POST**		**/api/books**		Menambahkan buku baru

**GET**			**/api/books**		Menampilkan semua buku

**GET**			**/api/books/:id**	Menampilkan detail buku berdasarkan ID

**PUT**			**/api/books/:id**	Memperbarui buku berdasarkan ID

**DELETE**		**/api/books/:id**	Menghapus buku berdasarkan ID

**2. CRUD Categories**

**Method**		**Endpoint**				**Deskripsi**

**POST** 		**/api/categories**			Menambahkan kategori baru

**GET**			**/api/categories**			Menampilkan semua kategori

**GET**			**/api/categories/:id**		Menampilkan detail kategori berdasarkan ID

**GET**			**/api/categories/:id/books**	Menampilkan buku berdasarkan kategori ID

**PUT**			**/api/categories/:id**		Memperbarui kategori berdasarkan ID

**DELETE**		**/api/categories/:id**		Menghapus kategori berdasarkan ID

**3. Authentication**

**Method**		**Endpoint**			**Deskripsi**

**POST**		**/api/users/login**		Login untuk mendapatkan JWT

**POST**		**/api/users/register**	Registrasi pengguna baru

**Autentikasi**

* Gunakan **JWT** untuk autentikasi.
* Tambahkan header **Authorization** dengan token JWT pada endpoint yang memerlukan autentikasi:

`Authorization: Bearer <your_token>`

**Validasi**

1. **Tahun Rilis Buku** :

    Hanya menerima nilai antara**1980** hingga **2024**.

2. **Ketebalan Buku** :

* tebal: Jika jumlah halaman > 100.
* tipis: Jika jumlah halaman <= 100.
