package main

import "fmt"

func main() {
	//array
	names := [...]string {"Buk", "Agus", "Hot" , "Asem", "B1nal", "Kecut"}

	// sebelah kiri adalah index awal yang ingin dicari, sebelah kanan adalah index-1 yang akan dicari jadi 4:6 = 4-5
	slice1 := names[4:6]
	slice2 := names[:3]
	slice3 := names[3:]
	slice4 := names[:]
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)



	//append slice seperti ini, lengkapnya cek di ytb Programmer zaman now versi golang dasar 5jam

	days := [...]string {
		"senin",
		"selasa",
		"rabu",
		"kamis",
		"jumat",
		"sabtu",
		"minggu",
	}

	//ambil hari sabtu dan minggu
	daySlice1 := days[5:7]
	fmt.Println(daySlice1)

	//rubah nilai sabtu dan minggu menggunakan index pada slice
	daySlice1[0] = "sabtu-baru"
	daySlice1[1] = "minggu-baru"

	fmt.Println(daySlice1)
	fmt.Println(days) //value dari arraynya akan berubah juga

	//menambahkan index baru ke slice 2 tidak akan lagi mempengaruhi value dari array asal dan slice1
	daySlice2 := append(daySlice1 , "libur-baru")
	daySlice2[0] = "hanya-pada-slice2"

	fmt.Println(daySlice1)
	fmt.Println(days)
	
	fmt.Println(daySlice2)


	//membuat slice langsung dengan make(array, len, capacity)
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Janu"
	newSlice[1] = "Asem"
	
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	//append data baru, yang akan masuk ke array yang sama
	newSlice2 := append(newSlice , "Ganteng")
	
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "MAMA"
	fmt.Println(newSlice)
	fmt.Println(newSlice2)

	//copy slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice) , cap(fromSlice)) //buat slice

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	//hati2 saat membuat array karena cara membuatnya cukup mirip dengan slice
	iniArray := [...]int {1,2,3,4,5}
	iniSlice := []int {1,2,3,4,5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}	