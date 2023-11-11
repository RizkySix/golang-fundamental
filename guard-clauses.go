package main

import ("fmt")


func main(){
	fmt.Println(getInsuranceAmountNastyIf(530))
	fmt.Println(getInsuranceAmount(540))
}

//nested if susah dibaca jika tidak menggunakan guard clauses
func getInsuranceAmountNastyIf(amount int) string {
	msg := "You have no Insurance"
	if amount < 100 {
		msg = "Your insurance in Standar class"
	} else {
	  if amount >= 101 && amount < 500 {
		msg = "Your insurance in Medium class"
	  } else {
		if amount >= 500 {
			msg = "Your insurance in High class"
		  if amount >= 500 && amount < 550{
			msg = "Your insurance in High-midle class"
		  }
		} else {
			msg = "You have no Insurance"
		}
	  }
	}
	return msg
}


//menggunakan guard clauses code menjadi clean dan mudah dibaca
func getInsuranceAmount(amount int) string {
	if amount < 100 {
		 return "Your insurance in Standar class"
	}
	if amount >= 101 && amount < 500 {
		 return "Your insurance in Medium class"
	}
	if amount >= 550{
		 return "Your insurance in High class"
	}
	if amount >= 500 && amount < 550{
		 return "Your insurance in High-midle class"
	}
	 return "You have no Insurance"
  }

  
