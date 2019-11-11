package main

import (
  "fmt"
  )

func main() {
	fmt.Println("Hello World")
	fakeFunc(0,0,0,0,"","","")


}



func fakeFunc(a,b,c,d int, e,f,g string) {
	x := Agreement{
		Validations:[]Validation{

		},
		Truth:&Truths{
			Signifier:"and",
		},
	}

	if !x.Argue(a,b,c,d,e,f,g) {
			fmt.Println("failed")
		return
	}
		fmt.Println("pass")

  
}
/*

{
  "validations": [
    {
      "variable": "",
      "min": 5,
      "max": 10
    },
    {
      "variable": "",
      "regex": "[a-zA-Z]{2,}"
    }

  ],
  "truth": [
    {
      "signifier": "and",
      "inputs": []
    }
  ]
}

*/
