package main

import "fmt"

type payment interface {
	paid(money int , price int) (sisa float64 , status bool)
}

type bca struct {
	payloadID string
	money int
	price int
}

type bri struct {
	payloadID string
	money int
	price int
}

type mandiri struct {
	payloadID string
	money int
	price int
}

type invalid struct{}

func (bca bca) paid(money int , price int) (sisa float64 , status bool) {
	count := money - price
	if count >= 0 {
		return float64(count) , true
	}
	
	return float64(count) , false
}

func (bri bri) paid(money int , price int) (sisa float64 , status bool) {
	count := money - price
	if count >= 0 {
		return float64(count) , true
	}
	
	return float64(count) , false
}


func (mandiri mandiri) paid(money int , price int) (sisa float64 , status bool) {
	count := money - price
	if count >= 0 {
		return float64(count) , true
	}
	
	return float64(count) , false
}

func (i invalid) paid(money int , price int) (sisa float64 , status bool) {
	return 0.0 , false
}




func getMessage(p payment) (string , float64 , bool){
	switch v := p.(type) {
	case bca:
		sisa , status := p.paid(v.money , v.price)
		return fmt.Sprintf("%s-bca" , v.payloadID) , sisa , status
	case bri:
		sisa , status := p.paid(v.money , v.price)
		return fmt.Sprintf("%s-bri" , v.payloadID) , sisa , status
	case mandiri:
		sisa , status := p.paid(v.money , v.price)
		return fmt.Sprintf("%s-mandiri" , v.payloadID) , sisa , status
	default:
		return "" , 0.0, false

	}
} 

func test(p payment) {
	payloadID , sisa , status := getMessage(p)

	msg := fmt.Sprintf("Invalid payment method")
	if status == true && payloadID != "" {
		msg = fmt.Sprintf("Pembayaran berhasil, sisa saldo Rp %.2f payloadID=%s" , sisa , payloadID)
	}
	if status == false && payloadID != "" {
		msg = fmt.Sprintf("Pembayaran gagal, total saldo kurang Rp %.2f payloadID=%s" , sisa , payloadID)
	}
	fmt.Println(msg)
	fmt.Println("=====================================================")
}

func main(){

	test(bri{
		payloadID : "f821jk82",
		money : 10000,
		price : 5000,
	})

	test(bca{
		payloadID : "ujk6621stw",
		money : 80000,
		price : 250000,
	})

	test(mandiri{
		payloadID : "lkjasf8761g",
		money : 253200,
		price : 250000,
	})
	

	test(invalid{})
}