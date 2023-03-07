sequential program
=> pencarian dalam suatu program dimana task pertama harus selesai terlebih dahulu sebelum task selanjutnya dimasukan. 

contoh :
pada sebuah program pencarian terdapat tiga menu yaitu video, website, dan iamge yang masing masing memiliki waktu eksekusi yang berbeda.
karena video memiliki waktu eksekusi yang lebih cepar yaitu 2 detik maka video di eksekusi terlebih dahulu sebelum website dengan waktu 3 detik dan image 4 detik. 
dengan total waktu yang dibutuhkan 9 detik. maka ketiga item tersebut akan muncul di tampilan website. 

paralel program 
=> cara kerja paralel program yaitu simultan artinya dapat mengerjakan sekaligus task yang tersedia tanpa harus menunggu task sebelumnya selesai.
dengan syarat harus ada multi core pada server. namun jika hanya 1 core hanya bisa bekerja secara sequential.  

concurrent program
=> menjalankan tugas secara bersaman dan independent.
maksud secara independent yaitu dapat berdiri sendiri walaupun core pada server hanya 1. tetap dapat berjalan secara bersama-sama, namun pada
saat proses terdapat pembagian.

ketika ingin menjalankan secara concurrent dan terdapat multicore. sehingga program tersebut akan berjalan secara concurrent dan paralel.
karena akan menggunakan source tersebut. 

mengimplentasikan concurrent pada golang yaitu :

1. goroutines
=> sebuah proses concurrent yang kita jalankan didalam golang. 
contoh concurrent yang sering ditemui pada golang yaitu func main() {} yang merupakan concurrent utama pada golang. 
pada goroutine terdapat channels yang berguna untuk berkomunikasi antara dua task. 
selain itu juga ada select untuk mengatur jalannya channel pada concurrent. 
 
