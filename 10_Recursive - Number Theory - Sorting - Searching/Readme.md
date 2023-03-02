recrusif : situasi ketika kita membuat sebuah function kemudian function tersebut di panggil kembali ( pemanggilan diri nya sendiri ).
dengan recrusive jika terdpat masalah yg ringan maka solusi didapat lebih cepat. jika masalah yang didapat itu besar maka recrusive akan membagi solusi tersebut menjadi beberapa bagian. 

contoh masalah faktorial :

public class algoritma {

	static int faktorial(int n) {
		if (n == 1) {
			return 1;
		} else {
			return n * faktorial(n - 1); // akan melakukan pengurangan terhadap nilai n.
		}
}


number theory : adalah cabang matematika yang mempelajari tentang bilangan.
number theory yang terkenal yaitu seperti prima, komposit, bilangan ganjil, bilangan genap, dll.

searching : proses pencarian dalam menemukan posisi nilai tertentu dalam suatu daftar nilai. 
salah satu jenis searching yaitu linier search dengan notasi O(n).\

sedangkan sorthin yaitu menyususn dara dalam urutan tertentu. jenis pengurutan yaitu selection sort, merge sort, 