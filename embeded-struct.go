package main

import "fmt"

type sender struct {
	rateLimit int
	user //embed user struct kedala sender
}

type user struct {
	name   string
	number int
}

// don't edit below this line

//embeded struct merupakah shorthand untuk menyederhanakan pemanggilan field dari instansiasi struct misal seperti berikut
func test(s sender) {
	fmt.Println("Sender name:", s.name)
	fmt.Println("Sender number:", s.number)
	fmt.Println("Sender rateLimit:", s.rateLimit)
	fmt.Println("====================================")
}
//jika tidak menggunakan user struct tidak diembed ke dalam sender, seharusnya pemanggilan diatas seperti ini s.user.name, s.user.number

func main() {
	test(sender{
		rateLimit: 10000,
		user: user{ //cara instansiasi embeded user struct
			name:   "Deborah",
			number: 18055558790,
		},
	})
	test(sender{
		rateLimit: 5000,
		user: user{
			name:   "Sarah",
			number: 19055558790,
		},
	})
	test(sender{
		rateLimit: 1000,
		user: user{
			name:   "Sally",
			number: 19055558790,
		},
	})
}

//full course tentanng embeded struct pada link berikut : https://boot.dev/course/3b39d0f6-f944-4f1b-832d-a1daba32eda4/59d90390-f479-4472-bb16-9af599285229/e30f8be3-30e4-4528-ab93-54461538314e