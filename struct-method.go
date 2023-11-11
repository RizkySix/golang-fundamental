package main

import "fmt"

type authenticationInfo struct {
	username string
	password string
}


//method struct(behavior) merupakan sebuah func yang menerima(receiver) parameter dari struct sebelum nama func
func (auth authenticationInfo) getBasicAuth() string{
	return fmt.Sprintf("Authorization: Basic %s:%s" , auth.username , auth.password)
} 

// don't touch below this line

func test(authInfo authenticationInfo) {
	fmt.Println(authInfo.getBasicAuth()) //memanggil behavior/method getBasicAuth dari struct authenticationInfo
	fmt.Println("====================================")
}

func main() {
	test(authenticationInfo{
		username: "Google",
		password: "12345",
	})
	test(authenticationInfo{
		username: "Bing",
		password: "98765",
	})
	test(authenticationInfo{
		username: "DDG",
		password: "76921",
	})
}

//full course tentanng method struct pada link berikut : https://boot.dev/course/3b39d0f6-f944-4f1b-832d-a1daba32eda4/59d90390-f479-4472-bb16-9af599285229/57ab6160-2769-4e7a-859c-6c4222768fdd