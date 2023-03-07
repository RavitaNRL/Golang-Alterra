package main

import "fmt"

func angkaRomawi(romawi int) string {
	// slice untuk menyimpan nilai angka dan simbol romawi.
	romanNum := []struct {
		nilai  int
		simbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	// variabel untuk menyimpan hasil konversi antara nilai menjadi romawi
	result := ""

	// loop melalui slice romanNumerals
	for _, rn := range romanNum {
		// selama nilai arabic lebih besar atau sama dengan Value
		for romawi >= rn.nilai {
			// tambahkan Symbol ke result
			result += rn.simbol

			// kurangi nilai arabic dengan Value
			romawi -= rn.nilai
		}
	}

	// mengulangi hasil konversi
	return result
}

func main() {
	// data yang akan dimasukan ke perulangan
	cases := []int{4, 9, 23, 2021, 1646}

	// perulangan melalui slice cases
	for _, c := range cases {
		// konversi angka menjadi angka romawi
		roman := angkaRomawi(c)

		// cetak hasil konversi
		fmt.Printf("Input nilai : %d\nOutput angka romawi : %s\n\n", c, roman)
	}
}
