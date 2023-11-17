package main

import "fmt"

type validationError struct {
	Message string
}

type notFoundError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"Validation error"}
	}

	if id != "Janu" {
		return &notFoundError{"Not found " + id}
	}

	//jika tidak ada error
	return nil
}

func main() {

	err := SaveData("kk", nil)
	// konversi menggunakan type assertion if, ok artinya boolean apakah ada error atau tidak

	//if err != nil {
	//	// ada error
	//	if validationErr, ok := err.(*validationError); ok {
	//		fmt.Println("Error:", validationErr.Error())
	//	} else if notFoundErr, ok := err.(*notFoundError); ok {
	//		fmt.Println("Error:", notFoundErr.Error())
	//	} else {
	//		fmt.Println("What the error?", err.Error())
	//	}
	//} else {
	//	// success
	//	fmt.Println("Success added")
	//}

	// menggunakan switch case (lebih sederhana)
	if err != nil {
		switch finalError := err.(type) {
		case *validationError:
			fmt.Println("Error:", finalError.Error())
		case *notFoundError:
			fmt.Println("Error:", finalError.Error())
		default:
			fmt.Println("What the error?", finalError.Error())
		}
	} else {
		fmt.Println("Success Added")
	}
}
