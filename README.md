# Web Application

## Live Coding - Go EduHub 3

### Implementation technique

Siswa akan melaksanakan sesi live code di 15 menit terakhir dari sesi mentoring dan di awasi secara langsung oleh Mentor. Dengan penjelasan sebagai berikut:

- **Durasi**: 15 menit pengerjaan
- **Submit**: Maximum 10 menit setelah sesi mentoring menggunakan `grader-cli submit`
- **Obligation**: Wajib melakukan _share screen_ di breakout room yang akan dibuatkan oleh Mentor pada saat mengerjakan Live Coding.

### Description

**Go Eduhub** adalah sebuah aplikasi yang dirancang untuk membantu pengelolaan dan manajemen data siswa dan kursus menggunakan bahasa pemrograman Go. Aplikasi ini memungkinkan pengguna untuk melakukan berbagai operasi seperti menambah dan menghapus data siswa serta kursus yang terkait dengan siswa tersebut.

Dalam live-code ini, kita akan mengimplementasikan API menggunakan _Golang web framework Gin_ untuk mengelola data _student_ dan _course_. API harus mengizinkan client untuk:

- Menambahkan siswa baru
- Menghapus siswa yang ada
- Menambahkan kursus ke daftar kursus siswa
- Menghapus kursus dari daftar kursus siswa

Disini sudah ditentukan endpoint untuk setiap operasi untuk mengimplementasikan logika yang diperlukan dari setiap operasi menggunakan repository student dan course.

### Constraints

Pada live code ini, kamu harus melengkapi fungsi dari repository `student` dan `course` ini memiliki implementasi function-function berikut:

📁 **repository**

Ini adalah fungsi yang berinteraksi dengan database Postgres menggunakan GORM:

- `repository/student.go`
  - `Delete`: Function ini menggunakan library GORM untuk menghapus data mahasiswa yang memiliki `id` yang sesuai dengan nilai yang diberikan sebagai argumen. Pertama-tama, function akan mengeksekusi sebuah query `DELETE` untuk menghapus data mahasiswa tersebut dari tabel `students`.
    - Jika proses tersebut berhasil, function akan mengembalikan `nil` sebagai `error`.
    - Namun jika terjadi error pada proses tersebut, function akan mengembalikan `error` yang terjadi.

- `repository/course.go`
  - `Delete`: Function ini akan menghapus data kursus yang memiliki `id` yang sesuai dengan nilai yang diberikan sebagai argumen. Pertama-tama, function akan mengeksekusi sebuah query untuk menghapus data kursus pada tabel `courses` dengan `id` yang sesuai.
    - Jika proses tersebut berhasil, function akan mengembalikan `nil` sebagai `error`.
    - Namun jika terjadi error pada proses tersebut, function akan mengembalikan `error` yang terjadi.

📁 **api**

- `api/student.go`
  - `DeleteStudent`: fungsi ini akan menghapus data siswa yang memiliki `id` yang sesuai dengan nilai yang diberikan sebagai argumen. Pertama-tama, fungsi akan mengubah `id` dari string ke integer menggunakan fungsi `strconv.Atoi`. Kemudian, fungsi akan memanggil fungsi `api.studentRepo.Delete` untuk menghapus data siswa pada tabel `students` dengan `id` yang sesuai.
    - Jika proses tersebut berhasil, fungsi akan mengembalikan status code `200` dan sebuah JSON response dengan pesan `success`.
    - Namun jika terjadi error pada proses tersebut, fungsi akan mengembalikan status code `http.StatusInternalServerError` dan sebuah JSON response dengan pesan error yang terjadi menggunakan struct `model.ErrorResponse`.

- `api/course.go`
  - `DeleteCourse`: fungsi ini digunakan untuk menghapus data course yang sudah ada di dalam sistem. Course yang akan dihapus diidentifikasi melalui `courseID` yang diambil dari parameter URL.
    - Jika terjadi error dalam proses validasi ID atau proses penghapusan data dari database, maka API akan mengembalikan response JSON dengan status HTTP `400` Bad Request atau `500` Internal Server Error masing-masing beserta pesan error yang dihasilkan.
    - Jika operasi berhasil dilakukan, maka API akan mengembalikan response JSON dengan status HTTP `404` Not Found (terkesan sedikit keliru, karena seharusnya HTTP status code yang digunakan adalah `200` OK) dan pesan sukses.

### Perhatian

Sebelum kalian menjalankan `grader-cli test`, pastikan kalian sudah mengubah database credentials pada file **`main.go`** (line 24) dan **`main_test.go`** (line 36) sesuai dengan database kalian. Kalian cukup mengubah nilai dari  `"username"`, `"password"` dan `"database_name"`saja.

Contoh:

```go
dbCredentials = Credential{
    Host:         "localhost",
    Username:     "postgres", // <- ubah ini
    Password:     "postgres", // <- ubah ini
    DatabaseName: "kampusmerdeka", // <- ubah ini
    Port:         5432,
}
```

### Test Case Examples

#### Test Case

**Input**:

```http
DELETE /student/delete/{id} HTTP/1.1
Host: localhost:8080
```

**Expected Output / Behavior**:

- Jika permintaan berhasil dan ID siswa valid, server harus mengembalikan kode status HTTP `200 OK` dan respons JSON dengan pesan sukses.

  ```json
  {
      "message": "student delete success"
  }
  ```

- Jika permintaan gagal karena ID siswa tidak valid, server harus mengembalikan kode status HTTP `400 Bad Request` dan respons JSON dengan pesan kesalahan.

  ```json
  {
      "error": "Invalid student ID"
  }
  ```

- Jika siswa dengan ID yang diberikan tidak ditemukan, server harus mengembalikan kode status HTTP `404 Not Found` dan respons JSON dengan pesan kesalahan.

  ```json
  {
      "error": "[error messages]"
  }
  ```

- Jika terjadi kesalahan saat menghapus siswa, server harus mengembalikan kode status HTTP `500 Internal Server Error` dan respons JSON dengan pesan kesalahan.

  ```json
  {
      "error": "[error messages]"
  }
  ```
