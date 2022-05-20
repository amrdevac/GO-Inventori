# Login Sistem JWT 

## Package

- JWT : https://github.com/golang-jwt/jwt
- MYSQL Driver : github.com/jinzhu/gorm/dialects/mysql
- ORM : github.com/jinzhu/gorm
- Framework : github.com/gin-gonic/gin

### Implements

- Chain Method
- Mysql Connection
- Implement Model, Controller, Service

## Service Implemented

Untuk service yang bertujuan pemakaian generic , terdapat di folder app/Provider didalamnya terdapat beberapa service yang sudah di implementasikan :

- RequestJson
  Untuk melakukan validasi pada request json yang dikirimkan

- ErrorHandler _(chain method)_
  untuk menampilkan log sat sebuah case mengalami kejanggalan. pada implementasinya , menggunakan chain method , terdapat beberapa pilihan yang sudah di implemtnasikan pada chain tersebut sepert ,
  - Fatal
    untuk menampilkan log error secara fatal dan memberhentikan sistem
    ErrorHandler.Err(**err**).Check(**"Text Log Gagal"**).Error() //Java
    ```ErrorHandler.Err(err).Check("Text").Error() //Go```
  - Error
    untuk menampilkan log error secara biasa tanpa memberhentikan sistem ataupun proses
    _ErrorHandler.Err(**err**).Check(**"Text Log Gagal"**).Fatal()_

- Hash
  untuk memproses hal hal yang berhubungan dengan hash , seperti ,
  - Hashing string yang di kirimkan di propteri function
  - Verifikasi kesamaan string yang dikirimkan dengan hash yang tersedia
- ResponseHandler _(chain method)_
  untuk menampilkan response dengan ginContext yang sudah di masukan ke dalam bagian properties function tersebut

## Modul Implemented

- Autenikasi
  - Register
  - Login 
  - Read Update Delete User
