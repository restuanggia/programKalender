package main

import (
	"fmt"
)

func getDayNumber(day, month, year int) int {
	t := []int{0, 3, 2, 5, 0, 3, 5, 1, 4, 6, 2, 4}
	if month < 3 {
		year--
	}
	return (year + year/4 - year/100 + year/400 + t[month-1] + day) % 7
}

func main() {
	var month, year int
	fmt.Print("Masukkan Bulan (1 - 12): ")
	fmt.Scan(&month)
	fmt.Print("Masukkan Tahun: ")
	fmt.Scan(&year)

	months := []string{"", "Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"}
	dayInMonth := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	// Cek tahun kabisat
	if year%400 == 0 || (year%100 != 0 && year%4 == 0) {
		dayInMonth[2] = 29
	}

	day := 1
	dayNumber := getDayNumber(day, month, year)

	fmt.Printf("\n\nKalender %s %d\n", months[month], year)
	fmt.Println("Minggu  Senin  Selasa Rabu  Kamis  Jumat  Sabtu")

	// Cetak spasi untuk hari pertama
	for i := 0; i < dayNumber; i++ {
		fmt.Print("       ")
	}

	// Cetak tanggal-tanggal dalam bulan
	for i := day; i <= dayInMonth[month]; i++ {
		fmt.Printf("%2d     ", i)
		dayNumber++
		if dayNumber > 6 {
			dayNumber = 0
			fmt.Println()
		}
	}
}
