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
    ```Java
    ErrorHandler.Err(err).Check(string).Error()
    ```
  - Error
    untuk menampilkan log error secara biasa tanpa memberhentikan sistem ataupun proses
    ```Java
    ErrorHandler.Err(err).Check(string).Fatal()
    ```
- Hash
  untuk memproses hal-hal yang berhubungan dengan hashing / enskripsi string , seperti ,

  - Hashing string yang di kirimkan di propteri function

  ```Java
    Hash.Make(string)
  ```

  - Verifikasi kesamaan string yang dikirimkan dengan hash yang tersedia

  ```Java
      Hash.Verify(stringPassword, stringMathingPasswor)
  ```

- ResponseHandler _(chain method)_

  untuk menampilkan response dengan ginContext yang sudah di masukan ke dalam bagian properties function tersebut

  ```Java
      ResponseHandler.Go(ginContext).SetData(interface{}).SetHttpStatus(intHttpStatus).Get()
  ```

## Modul Implemented

- Autenikasi
  - Register
  - Login
  - Read Update Delete User
