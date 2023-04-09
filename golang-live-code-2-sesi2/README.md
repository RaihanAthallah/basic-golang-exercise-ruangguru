# Dasar Pemrograman Backend

## Live Coding - Student Management System

### Implementation technique

Siswa akan melaksanakan sesi live code di 15 menit terakhir dari sesi mentoring dan di awasi secara langsung oleh Mentor. Dengan penjelasan sebagai berikut:

- **Durasi**: 15 menit pengerjaan
- **Submit**: Maximum 10 menit setelah sesi mentoring menggunakan `grader-cli`
- **Obligation**: Wajib melakukan _share screen_ di breakout room yang akan dibuatkan oleh Mentor pada saat mengerjakan Live Coding.

### Description

Proyek ini adalah sebuah program _command-line interface (CLI)_ sederhana yang ditulis dalam bahasa Go, yang memungkinkan kita untuk mengelola data mahasiswa, seperti menambahkan mahasiswa baru, menghapus mahasiswa, menampilkan data mahasiswa, dan memperbaharui skor mahasiswa.

Program ini memiliki 4 fungsi utama:

- `ViewStudents`: untuk menampilkan data mahasiswa
- `AddStudent`: untuk menambahkan mahasiswa baru ke dalam sistem
- `RemoveStudent`: untuk menghapus mahasiswa dari sistem
- `UpdateScore`: memperbarui skor mahasiswa dalam sistem

Program ini menggunakan `map` untuk menyimpan data mahasiswa, dengan key berupa nama mahasiswa dan value berupa slice yang berisi alamat, nomor telepon, dan nilai mahasiswa.

### Constraints

Program ini dibagi menjadi 4 bagian:

- **Main**: mengontrol keseluruhan aliran program dan memanggil fungsi lainnya
- **ViewStudents**: menampilkan data mahasiswa dalam format tabel
- **AddStudent**: menambahkan mahasiswa baru ke dalam sistem
- **RemoveStudent**: menghapus mahasiswa dari sistem
- **UpdateScore**: memperbarui skor mahasiswa dalam sistem

Program ini menggunakan switch statement untuk memproses input pengguna dan memanggil fungsi yang sesuai untuk mengelola data mahasiswa. Jika pengguna memasukkan input yang tidak valid, program akan menampilkan pesan kesalahan.

Berikut adalah penjelasan dari fungsi-fungsi yang harus diimplementasi:

- Fungsi **RemoveStudent** menerima parameter `*map[string][]interface{}` dan mengembalikan sebuah fungsi yang menerima parameter string, yang akan menghapus data mahasiswa dengan nama yang diberikan dari `map`.

  ```go
  func RemoveStudent(students *map[string][]interface{}) func(string) {
    // TODO: answer here
  }
  ```

- Fungsi **UpdateScore** menerima parameter `*map[string][]interface{}` dan mengembalikan sebuah fungsi yang menerima 2 parameter yaitu name (string) dan score (int). Fungsi tersebut akan memperbarui skor siswa dengan name yang diberikan di pointer `map` of `students` yang diberikan dan tidak mengembalikan nilai apa pun.

  ```go
  func UpdateScore(students *map[string][]interface{}) func(string, int) {
    return func(name string, score int) {
      // TODO: answer here
    }  
  }
  ```

### Test Case Examples

**Input/Output**:

```bash
> Enter command (add, remove, update-score, view): add
> Enter name: John
> Enter address: Sudirman
> Enter phone: 555-1234
> Enter score: 90
> Enter command (add, remove, update-score, view): add
> Enter name: Jane
> Enter address: Ciganjur
> Enter phone: 555-5678
> Enter score: 85
> Enter command (add, remove, update-score, view): view
> Student data:
> Name  Address Phone   Score
> Jane  Ciganjur  555-5678    85
> John  Sudirman  555-1234    90

> Enter command (add, remove, update-score, view): update-score
> Enter name: John
> Enter new score: 95
> Score updated:
> Name  Address Phone   Score
> Jane  Ciganjur  555-5678    85
> John  Sudirman  555-1234    95
> Enter command (add, remove, update-score, view): remove
> Enter name: Jane
> Enter command (add, remove, update-score, view): view
> Student data:
> Name  Address Phone   Score
> John  Sudirman  555-1234    95
```

**Explanation**:

> Program ini memungkinkan pengguna untuk menambah, menghapus, memperbarui score, atau melihat data siswa. Dalam contoh ini, pengguna menambahkan dua siswa, melihat datanya, memperbarui skor John, menghapus Jane dari data, dan melihat data yang diperbarui.
