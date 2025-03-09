# ExerciseGolang

**1. Temukan dan Perbaiki Kesalahan dalam Kode**  
Peserta harus menemukan dan memperbaiki setidaknya 5 kesalahan dalam kode.  
Berikut adalah beberapa kesalahan yang harus diperbaiki:

- **Status code salah** (500 pada `getBook`, harusnya 404).
- **Tidak ada response JSON pada `deleteBook`** (harus mengembalikan pesan sukses).
- **Memperbarui buku bisa mengganti ID** (`updateBook` tidak mempertahankan ID lama).
- **Jika `books` kosong, `getBooks` tetap mengembalikan array kosong tanpa status code yang sesuai**.
- **Logika penghapusan tidak mengembalikan response yang benar**.

❗ **Catatan:** Peserta tidak perlu menambahkan validasi input baru (misalnya pengecekan apakah ID unik atau tidak).  
Tugas hanya memperbaiki bug yang sudah ada tanpa menambahkan validasi tambahan.

---

**2. Buat Branch Baru dan Upload Perbaikan**  
- Fork atau Clone repository utama.
- Buat branch baru dengan nama peserta.
- Commit setiap perubahan dengan deskripsi yang jelas.

---

**3. Tambahkan Logging untuk Setiap Logic**  
Peserta harus menambahkan `log.Println()` di setiap fungsi untuk mempermudah debugging, misalnya:

- **Sebelum menampilkan daftar buku**.
- **Sebelum menghapus buku**, log **ID buku** yang dihapus.
- **Sebelum memperbarui buku**, log **ID dan judul buku**.

---
