Database adalah sekumpulan data yang terorganisir.
ACCOUNT MODEL:
Display Name    = Nurhadi - Aldo
Username        = @nurhadi_aldo
Bio             = Akun Resmi Relawan..
Location        = Indonesia
Join Date       = 20/12/2018
Data Base Relationship antara lain :
1. One To One Relationship # 1 user hanya memiliki 1 foto profil
2. One To Many #1 user bisa memiliki banyak tweets
3. Many To Many Relationship # 1 user bisa memiliki banyak follower user, 1 user bisa difollow banyak user
DDL (Data Definition Language) adalah perintah untuk membuatstruktur dari objek
Statement Operatin DDL antara lain :
1. Like / Between   = Tampilan data user_id dan message table tweets yang message mengandung huruf H didepan dan antara 1 dan 3 seperti contoh di materi 
2. And / Or         = Tampilan data user_id dan message table tweets yang message mengandung huruf H didepan dan user_id antara 1 dan 3 seperti contoh di materi 
3. Order By         = Tampilan data user_id dan message table tweets yang message mengandung huruf H didepan dan user_id antara 1 dan 3 diurutkan berdasarkan id tweets berurutan dari terbesar kecil 
4. Limit            = Tampilan data user_id dan message table tweets yang message mengandung huruf H didepan dan user_id antara 1 dan 3 diurutkan berdasarkan id tweets berurutan dari terbesar ke terkecil di batasi 2 data
DML (Data Manipulation Language) adalah perintah yang digunakan untuk memanipulasi data dalam tabel dari suatu database.
Statement Operatin DML antara lain :
1. Insert = input data ke tabel user
2. Select = menampilkan semua data pada tabel user
3. Update = ubah data fullname ke tabel user dengan id 1
4. Delete = hapus data pada tabel user dengan id 1