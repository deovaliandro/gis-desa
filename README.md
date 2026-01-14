# 📍 Pemetaan Spasial Sulawesi Barat

Sistem Informasi Geografis (SIG) Desa berbasis **Go (Gin Framework)** dan
**MongoDB** untuk visualisasi peta tematik dan pengelolaan data desa di Provinsi
Sulawesi Barat.

Aplikasi ini mendukung:

- Halaman publik (guest)
- Visualisasi peta tematik
- Manajemen data desa
- Manajemen jenis peta (admin only)
- Sistem autentikasi berbasis peran (admin & user)

---

## 🧩 Teknologi yang Digunakan

|Komponen|Teknologi|
|---|---|
|Backend|Go (Gin Framework)|
|Database|MongoDB|
|Frontend|Go HTML Template|
|CSS|Tailwind CSS|
|Build CSS|Node.js 20|
|Peta|Leaflet.js|
|Session|gin-contrib/sessions|
|Auth|Role-based (Admin & User)|

---

## ⚙️ Prasyarat

### 1️⃣ Go (minimal versi 1.20)

```bash
go version
```

Unduh jika belum tersedia: <https://go.dev/dl/>

---

### 2️⃣ MongoDB

Unduh MongoDB Community Edition:
<https://www.mongodb.com/try/download/community>

Jalankan service MongoDB:

```bash
mongod
```

---

## 🔐 Konfigurasi Environment

Buat file `.env` atau ganti nama file `.env.example` menjadi `.env` di root
project:

```env
APP_PORT=8080

MONGO_URI=mongodb://localhost:27017
MONGO_DB_NAME=gis_desa

SESSION_SECRET=super-secret-key
```

> ⚠️ Jangan commit file `.env` ke repository publik.

---

## 📦 Instalasi Dependency

```bash
go mod tidy
```

---

## 🗄️ Inisialisasi & Import Database (WAJIB)

Database **tidak perlu diisi manual**, karena sudah disediakan **backup data**
di folder `db-backup`.

### 1️⃣ Masuk ke Mongo Shell

```bash
mongosh
```

Buat database:

```js
use gis_desa
```

Keluar dari shell:

```js
exit
```

---

### 2️⃣ Import Data dari Folder `db-backup`

Pastikan Anda berada di **root project**.

#### Import koleksi `desa`

```bash
mongoimport \
  --db gis_desa \
  --collection desa \
  --file db-backup/desa.json.gz \
  --gzip \
  --jsonArray
```

```bash
mongoimport \
  --db gis_desa \
  --collection users \
  --file db-backup/users.json.gz \
  --gzip \
  --jsonArray
```

```bash
mongoimport \
  --db gis_desa \
  --collection maps \
  --file db-backup/maps.json.gz \
  --gzip \
  --jsonArray
```

---

### 3️⃣ Verifikasi Import

Masuk ke `mongosh`:

```bash
mongosh
use gis_desa
```

Cek jumlah data:

```js
db.desa.countDocuments()
db.users.countDocuments()
db.maps.countDocuments()
```

Jika hasil > 0, maka import berhasil.

---

## ▶️ Menjalankan Aplikasi Secara Lokal

```bash
go run main.go
```

Akses aplikasi melalui browser:

- Home: <http://localhost:8080>
- Login: <http://localhost:8080/login>
- Peta: <http://localhost:8080/map/pendidikan>
- Admin Desa: <http://localhost:8080/admin/desa>
- Manajemen Peta: <http://localhost:8080/admin/maps>

---

## 👥 Akun Default

Gunakan akun admin berikut:

|Username|Password|Role|
|---|---|---|
|admin|pass123|Admin|
|kelompok-1|kelompok-1|User|

gunakan pola `kelompok-1` sampai `kelompok-8` untuk akun user kelompok.

---

## 🗺️ Fitur Peta

- Peta ditampilkan dalam **Card**
- Layout halaman peta terbagi dua:
  - Kiri: Peta interaktif (Leaflet)
  - Kanan: Interpretasi peta
- Warna desa merepresentasikan **tingkat pendidikan**
- Data berasal dari MongoDB (GeoJSON + atribut)

---

## 📄 Lisensi

CC BY-NC-SA 4.0 © Deo Valiandro M.

---

## ✍️ Catatan

Dokumentasi ini disusun untuk:

- Reproduksibilitas riset
- Pengembangan lanjutan
- Kebutuhan akademik (S2)

Silakan sesuaikan jika struktur backup berubah.
