DOCKER
*************
docker
=> cara untuk melakukan deploying aplication kedalam suatu server.
untuk melakukan docker ada beberapa hal yang harus diinstall seperti, operating system, library, run time application, web server dan database. 
image 
=> hasil distribusi file atau package dan image akan di deploy ke dalam registry.
container
=> image yang kita running 

infrastruktur docker 
************************
ada 3 bagian :
1. client
=> perintah atau comand untuk pengoperasian pada docker host.
2. docker host
=> terdapat docker deamon, untuk mengatur komponen yang ada didalam docker host atau di control unix
3. registry
=> untuk menyimpan image pada docker agar lebih mudah dalm mengatur. 

hal yang dapat dilakukan dengan docker 
******************************************
1. build docker image
=> disimpan kedalam image repository atau docker hub
2. upload and publish images
3. download and run
=> di download kedalam docker file atau sistem atau laptop. 

sintaks yang digunakan pada docker 
**************************************
1. from
=> mendapatkan image yang akan kita pake.
2. run
=> mengeksekusi bash comand nya ketika membuild container.
3. env
=> untuk melakukan konfigurasi
4. add
=> untuk copy file kedalam image
5. copy
=> untuk mengcopy file
6. workdir
=> membuat folder/direktory kerja
7. entrypoint 
=> untuk mengeksekusi perintah ketikan selesai membuild container.
8. cmd
=> untuk mengeksekusi.


perbedaan container dgn virtual machine
****************************************
1. container != virtual machine.
=> berorientasi aplikasi dan kebutuhan aplikasi itu sendiri. maka dri itu membutuhkan dependency seperti library, package.
2. virtual machine
=> berorientasi mesin itu sendiri, menjalankan aplikasi dan operasi sistem.



PENGGUNAAN DOCKER
**********************
1. buatlah file docker didalam folder kerja.
2. didalam file nya masukkan command untuk proses docker. (dilakukan di vscode).
3. tambahkan  query (FROM golang:1.17-alpine).
4. jika sudah buat folder kerja menggunakan perintah workdir/folder ( workdir/app).
5. tambahkan perintah copy go.mod ./ dan copy go.sum ./
6. jika sudah masukkan perintah RUN go mod download
7. lalu ketikkan copy .. ( untuk mengcopy semua source code pada folder yg kita gunakan)
8. jika sudah run go build -o/dist (untuk menentukan output berada. di contoh itu nnti outpyt ada di dist).
9. ketik expose 3222 untuk mendifine dimana port yg digunakan 3222
10. ketik cmd["/dist"]





#mengecek versi
docker version

#mengecek images
docker images

#mengecek container semua
docker container ls --all

#mengecek container yang jalan
docker container ls 

#membuat container
docker container create --name golangserver1 golang  

note :
1. tidk bisa membuat container dengan nama yang sama, krna nama dalam container bersifat unik.
2. untuk golangserver1 ( nama container nya) , sedangkan golang (images nya).


#mendownload images
docker pull golang

#menjalankan container 
docker container start nama container (golangserver1)

#menstop container
docker container stop nama container (golangserver1)

#menghapus docker container
docker container rm golangserver1

#menghapus imager 
docker image rm (nama images) golang

#membuka port pada docker 
docker container create --name (nama container : golang1) -p 8000:27017  (nama images : golang). (8000 merupkan port eksternal sedangkan 27017 port internal).

#membuild images
docker build -- tag (nama tag) (nama image)  (direktori/folder tampat docker).
ex :docker build -- tag app-golang:1.0 . (lalu enter)

#push docker
docker push username dockerhub/nama images
ex : docker push username ravitaAsmi/app1:1.1

#membuat tag baru dari images yang sudah ada  
docker tag (nama image yang sudah ada) (username repo/ username dockerhub) (versi tag nya)
ex : docker tag app-golang:1.0 ravitaAsmi:1.0

#login registry
docker login 

note :
ketika di enter akan diminta untuk masukkan username dan password apda dockerhub.





note : 
1. jika salah membuat container langkah yang dilakukan hapus container, karna walaupun container dihapus namun images masi ada didalam docker.
2. jika ingin menghapus container lalu container tersebut masih jalan, maka di stop dlu baru hapus.
3. docker push untuk mempush images baru 
4. docker pull untuk download images yang sudah ada.
