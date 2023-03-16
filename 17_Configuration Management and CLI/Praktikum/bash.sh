#!/bin/bash

# Buat direktori utama
mkdir -p tegar/about_me/personal
mkdir -p tegar/about_me/professional
mkdir -p tegar/my_friends
mkdir -p tegar/my_system_info

# Buat file dalam setiap direktori
touch tegar/about_me/personal/facebook.txt
touch tegar/about_me/professional/linkedin.txt
touch tegar/my_friends/list_of_my_friends.txt
touch tegar/my_system_info/about_this_laptop.txt
touch tegar/my_system_info/internet_connection.txt

# ini struktur nya

#!/bin/bash

# Buat direktori utama
mkdir -p tegar/about_me/personal
mkdir -p tegar/about_me/professional
mkdir -p tegar/my_friends
mkdir -p tegar/my_system_info

# Buat file dalam setiap direktori
touch tegar/about_me/personal/facebook.txt
touch tegar/about_me/professional/linkedin.txt
touch tegar/my_friends/list_of_my_friends.txt
touch tegar/my_system_info/about_this_laptop.txt
touch tegar/my_system_info/internet_connection.txt

# Menambahkan isi pada file about_this_laptop.txt
echo "My username: tegarimansyah" > tegar/my_system_info/about_this_laptop.txt
echo "With host: Darwin cyberhome 18.6.0 Darwin Kernel Version 18.6.0: Thu Apr 25 23:16:27 PDT 2019; root:xnu-4903.261.4~2/RELEASE_X86_64 x86_64" >> tegar/my_system_info/about_this_laptop.txt

# Menambahkan isi pada file facebook.txt
echo "https://www.facebook.com/tegarimansyahfb" > tegar/about_me/personal/facebook.txt

# Menambahkan teks pada file internet_connection.txt
echo "Connection to google:" > tegar/my_system_info/internet_connection.txt
echo "PING forcesafesearch.google.com (216.239.38.120): 56 data bytes" >> tegar/my_system_info/internet_connection.txt
echo "64 bytes from 216.239.38.120: icmp_seq=0 ttl=51 time=30.073 ms" >> tegar/my_system_info/internet_connection.txt
echo "64 bytes from 216.239.38.120: icmp_seq=1 ttl=51 time=73.019 ms" >> tegar/my_system_info/internet_connection.txt
echo "64 bytes from 216.239.38.120: icmp_seq=2 ttl=51 time=31.850 ms" >> tegar/my_system_info/internet_connection.txt
echo "" >> tegar/my_system_info/internet_connection.txt
echo "--- forcesafesearch.google.com ping statistics ---" >> tegar/my_system_info/internet_connection.txt
echo "3 packets transmitted, 3 packets received, 0.0% packet loss" >> tegar/my_system_info/internet_connection.txt
echo "round-trip min/avg/max/stddev = 30.073/44.981/73.019/19.839 ms" >> tegar/my_system_info/internet_connection.txt

# Menambahkan teks pada file linkedin.txt
echo "https://www.linkedin.com/in/tegarimansyahlinkedin" > tegar/about_me/professional/linkedin.txt

# Menambahkan teks pada file list_of_my_friends.txt
echo "1) Achmad Miftahul - amulum
2) Achmad Rafiq - achmadrafiq97
3) Adiststi - adiststi
4) Agung - agungajin19
5) Azzahra - al7262
6) Charisma - fadzrichar
7) Farida - ulfarida
8) Garry - garryarielcussoy
9) Hamdi - hamdiranu
10) Hedy Gading - Gading09
11) Ilham - aamfatur
12) Lelianto - Lelianto
13) Mohammad - daffa99
14) Muhammad Fadhil - fabdulkarim
15) Ofbimon - bimon-alta
16) Purnama Syafitri - pipitmnr
17) Setyo - setyoyo07
18) Wildan - wiflash
19) Willy - sumarnowilly94
20) Woka - woka20" > tegar/my_friends/list_of_my_friends.txt

# Menampilkan struktur folder dan file yang telah dibuat
echo "Struktur folder dan file yang telah dibuat:"
ls -R tegar