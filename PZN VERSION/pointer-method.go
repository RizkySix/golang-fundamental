package main

import "fmt"

type Man struct {
	Name string
}

// method struct menggunakan * pointer/reference sehingga value sumber dari variable rizky akan dirubah dan tidak diduplicate
func (man *Man) married() {
	man.Name = "Mr." + man.Name
}

func main() {
	rizky := Man{"rizky"}
	rizky.married()
	fmt.Println(rizky.Name)
}
