clean code 
=> istilah pada suatu program yang bertujuan untuk mempermudah pembacaan, pemahaman dan pengubahan yang dilakukan oleh 
seorang programmer. 

kenapa harus di clean code ?
karena dengan menggunakan clean code dapat memudahkan development dalam
membangun suatu project.

clean code memiliki ciri yaitu :
1. mudah untuk dipahami
=> penggunaan variabel atau penamaan func pada suatu program mudah untuk dipahami ketika
menggunakn penamaan yang benar. 

contoh : 

*** PENAMAAN YANG MUDAH UNTUK DIPAHAMI
let userBalance = 125.0;
let primeList = [2, 3, 4, 7];
const locations = ['Austin', 'New York', ' San Francisco');
locations. forEach((location) =>
 doStuff();
 doSomeOtherStuff();
  //
     ..
  //
 dispatch(location);
 });


*** PENAMAAN YANG TIDAK MUDAH DIPAHAMI
let b = 125.0;
let data a = [2, 3, 5, 7];
const locations = ['Austin', 'New York', ' San Francisco'];
locations. forEach((1) =>
 doStuff();
 doSomeOtherStuff();
    Sebentar, '1' itu untuk apa lagi?
 dispatch(1);
 1);

2. mudah untuk dieja dan dicari 
=> penggunaan clean code pada suatu variabel atau func dapat memudahkan development atau pembaca 
untuk mengeja atau membaca dan membaca hal tersebu, sehingga dapat meminimalisir kesalahan dalam penulisan kembali code.

3. singkat namun mendeksripsikan konteks atau object
4. konsisten 
5. hindari penambahan konteks yang tidak perlu
6. komentar

pada clean code terdapat beberapa istilah :

1.KISS ( keep it so simple)
=> hindari penggunaan atau pembuatan func name untuk melakukan tugas A, sekaligus memodifikasi B dan mengecek C atau
hal yang rangkap dalam satu jenis func.

2. DRY (Don't Repeat Yourself
=>  fungsi yang digunakan untuk menghindari duplicat code secara berulang-ulang karena sering melakukan copy - paste. 

3. refactoring
=> proses restrukturisasi code yang telah dibuat, dengan cara merubah struktur internal tanpa mengubah perilaku eksternal. 


