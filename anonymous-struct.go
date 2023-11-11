package main

import "fmt"

//membuat anonymous struct seperti ini tidak terlalu disarankan karen susah dibaca
/* myCar := struct {
	Make string
	Model string
  } {
	Make: "tesla",
	Model: "model 3"
  } */


//lebih baik lakukan seperti ini
type car struct {
	Make string
	Model string
	Height int
	Width int
	// Wheel adalah anonymous struct yang tidak memiliki type, dan hanya akan dapat digunakan saat car struct digunakan
	Wheel struct {
	  Radius int
	  Material string
	}
  }

func main(){

	//instansiasi car struct
	myCar := car{
		Make: "Avanza",
		Model: "Seri A",
		Height: 15,
		Width: 15,
		Wheel: struct {
			Radius int
			Material string
		}{
			Radius: 45,
			Material: "Kayu jati",
		},
	}

	fmt.Printf(`Rizky telah membuat %s dengan model %s tingginya %v cm lebar %v cm dan rasio roda %v derajat menggunakan %s` , myCar.Make , myCar.Model , myCar.Height , myCar.Width , myCar.Wheel.Radius , myCar.Wheel.Material)
	fmt.Println()
}

//full course tentanng anonymous struct pada link berikut : https://boot.dev/course/3b39d0f6-f944-4f1b-832d-a1daba32eda4/59d90390-f479-4472-bb16-9af599285229/7256e43f-aea0-47bb-9671-9c55ebca7095