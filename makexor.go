package main

import(
	"fmt"
)

var(
	xor =[]byte{} //shellcode

)	 


func makexor(xor []byte){
	for i := 0; i<len(xor); i++{
		xor[i] = xor[i] ^ 0xff
	}
}

func main(){
	for i:= 0 ; i<len(xor); i++{
		if xor[i] > 0x0f{
			fmt.Printf("0x%x,",xor[i])
		}else{
			fmt.Printf("0x0%x,",xor[i])
		}	
	}
}


