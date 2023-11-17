package main

import (
	"fmt"
	"project-golang/database"
	_ "project-golang/internal"
)

func main() {
	fmt.Println(database.GetDatabase())
}
