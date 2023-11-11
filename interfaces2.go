package main

import (
	"fmt"
)

type employee interface {
	getName() string //struct dapat dikatan "mengimplementasi" employee interface jika menggunakan struct-method getName() dan getSalary()
	getSalary() int
}

type contractor struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

//struct contractor mengimplementasi employee karena memiliki struct-method getName() dan getSalary()
func (c contractor) getName() string {
	return c.name
}

func (c contractor) getSalary() int {
	return c.hourlyPay * c.hoursPerYear
}

// don't touch below this line

type fullTime struct {
	name   string
	salary int
}

func (ft fullTime) getSalary() int {
	return ft.salary
}

func (ft fullTime) getName() string {
	return ft.name
}

func test(e employee) {//test menerima employee interface sebagai parameter, sehingga parameter yang dikirim pada test harus memiliki struct-method getName() dan getSalary()
	fmt.Println(e.getName(), e.getSalary())
	fmt.Println("====================================")
}

func test2(con contractor) { //sifat dari struct-method tidak akan berubah, dan dapat dipanggil seperti biasa seperti berikut
	fmt.Println(con.getSalary() , "Ini berasal dari test2 diluar implementasi employee")
	fmt.Println("====================================")
}

func main() {
	test(fullTime{
		name:   "Jack",
		salary: 50000,
	})
	test(contractor{
		name:         "Bob",
		hourlyPay:    100,
		hoursPerYear: 73,
	})
	test(contractor{
		name:         "Jill",
		hourlyPay:    872,
		hoursPerYear: 982,
	})
	test2(contractor{
		name:         "Miku",
		hourlyPay:    100,
		hoursPerYear: 10,
	})

}

//full course tentanng interfaces contoh 2 pada link berikut : https://boot.dev/course/3b39d0f6-f944-4f1b-832d-a1daba32eda4/7faa7ccf-4fc0-406d-90b7-dac16e72f6c3/8c123c3a-46d5-46d3-a4c4-d1e2555af05b