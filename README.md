# test
Test_kerja 

TEST KERJA PT. SEMESTA ARUS TEKNOLOGI

Soal :
1. REST API data Siswa
    # go command : 
      - intitialize: go mod init test.com/siswa-apis
      - gin dependencies : go get github.com/gin-gonic/gin
      - mongodb dependencies:  go get go.mongodb.org/mongo-driver
    
    # routing CRUD siswa:
      - POST, insert data siswa :http://localhost:9090/siswa/insert 
      - GET, ambil semua data siswa: http://localhost:9090/siswa/all
      - GET, ambil data siswa by id: http://localhost:9090/siswa/:id            //Masih belum berhasil catch id nya (filter nya)
      - PUT, update data siswa: http://localhost:9090/siswa/update              //Masih belum berhasil mendapatkan filternya
      - DELETE, delete data siswa by id: http://localhost:9090/siswa/delete/:id //Masih belum berhasil mendapatkan filter id nya

    # JWT masih belum berhasil

2. Algoritma 
   a. Bilangan Prima : Run
   b. Pola Bintang : Run
   c. Sorting angka : Run
   d. Palindrome : Masih belum berhasil

3, Front End, Tidak sempat.


Quick Note : Mohon maaf sebesar-besarnya kepada Bpk/Ibu Pengetest, 
             saya hanya mampu mengerjakan sebagian kecil dikarenakan saya baru belajar Golang 1 minggu yang lalu 
             dan basic saya sebenarnya menggunakan Javascript, NodeJS, ExpressJS dan baru lulus dari bootcamp beberapa bulan yang lalu 
