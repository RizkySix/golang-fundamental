package main

import "fmt"

// nil hanya  bisa digunakan untuk tipe data map,slice,function,interface,pointer,dan channel
func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	data := newMap("Yak")

	if data == nil {
		fmt.Println("Nilai kosong atau nil")
	} else {
		fmt.Println(data["name"])
	}

}
