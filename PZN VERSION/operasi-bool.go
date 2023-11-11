package main

import "fmt"

func main() {
	nilaiAkhir := 90
	nilaiAbsen := 80
	nilaiAktif := 95

	lulusNilai := nilaiAkhir > 80
	lulusAbsen := nilaiAbsen > 80
	lulusAktif := nilaiAktif > 92

	lulus := lulusNilai && lulusAbsen && lulusAktif
	fmt.Println(lulus)
}