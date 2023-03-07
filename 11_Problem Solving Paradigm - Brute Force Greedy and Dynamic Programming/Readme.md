problem solving paradigma dibagi menjadi tiga yaitu :
1. brute force
2. divide & conqure
3. greeedy

brute force = complate search 
=> menyelesaikan maslah dengan melakukan pengecekan terhadap semua data.
 
complate search digunakan ketika tdk ditemukan algoritma yang lebih cepat untuk menyelesaikan masalah tersebut.
kompleksitas untuk complate search yaitu O(n). 
ex : ketik mencari nilai max dan min

devide & conqure 
=> menyelesaikan masalah menjadi beberapa bagian, yaitu secara devide , conqure dan combine. 
ex : permasalahan binary search. dengan mencari indek

procedure Binarysearcn (n, X)
                                     
   hasil <- 0
   kiri <- 1
   kanan <-N
   while ((kiri <= kanan) ^ (hasil=0)) do
     tengah <- (kiri+kanan) div 2
     if (X < h[tengah]) then
       kanan <- tengah - 1
     else if (X > h[tengah]) then
       kiri <- tengah + 1
     else
       hasil <- tengah
     end if
   end while
   if (hasil = 0) then
     print "beri hadiah lain"
   else
     print hasil
   end if
 end procedure

greedy 
=>  merupakan algoritma yang tamak, dia akan mengambil segala kemungkinan untuk memecahkan masalah. 
namun untuk beberapa kondisi greedy lebih unggul dibandingkan brute force. 
contoh : coin change 

dynamic programming 
=> algoritma untuk memecahkan masalah mengenai optimasi menajdi beberapa submasalah agar menjadi lebih sederhana.
cara kerja dari algoritma ini yaitu memecahkan masalah menjadi submaslah dan menyimpan solusi submasalah untuk menyelesaikan 
masalah yang lebih besar lagi. 
teknik penyelesaian masalah dengan dynamic programming yaitu bisa maju dari depan ke belakang, atau mundur dari belakang ke depan.
dimana hasil yang dapat nanti ny sama. 
